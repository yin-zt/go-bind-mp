package isql

import (
	"errors"
	"fmt"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/model/user"
	"github.com/yin-zt/go-bind-mp/pkg/util/common"
	"github.com/yin-zt/go-bind-mp/pkg/util/tools"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type UserService struct{}

// 当前用户信息缓存，避免频繁获取数据库
var userInfoCache = cache.New(24*time.Hour, 48*time.Hour)

// Add 添加资源
func (s UserService) Add(user *user.User) error {
	user.Password = tools.NewGenPasswd(user.Password)
	//result := common.DB.Create(user)
	//return user.ID, result.Error
	return common.DB.Create(user).Error
}

// List 获取数据列表
func (s UserService) List(req *request.UserListReq) ([]*user.User, error) {
	var list []*user.User
	db := common.DB.Model(&user.User{}).Order("id DESC")

	username := strings.TrimSpace(req.Username)
	if username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", username))
	}
	nickname := strings.TrimSpace(req.Nickname)
	if nickname != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", nickname))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Roles").Find(&list).Error
	return list, err
}

// ListCout 获取符合条件的数据列表条数
func (s UserService) ListCount(req *request.UserListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&user.User{}).Order("id DESC")

	username := strings.TrimSpace(req.Username)
	if username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", username))
	}
	nickname := strings.TrimSpace(req.Nickname)
	if nickname != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", nickname))
	}

	err := db.Count(&count).Error
	return count, err
}

// List 获取数据列表
func (s UserService) ListAll() (list []*user.User, err error) {
	err = common.DB.Model(&user.User{}).Order("created_at DESC").Find(&list).Error

	return list, err
}

// Count 获取数据总数
func (s UserService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&user.User{}).Count(&count).Error
	return count, err
}

// Exist 判断资源是否存在
func (s UserService) Exist(filter map[string]interface{}) bool {
	var dataObj user.User
	err := common.DB.Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Find 获取同名用户已入库的序号最大的用户信息
func (s UserService) FindTheSameUserName(username string, data *user.User) error {
	return common.DB.Where("username REGEXP ? ", fmt.Sprintf("^%s[0-9]{0,3}$", username)).Order("username desc").First(&data).Error
}

// Find 获取单个资源
func (s UserService) Find(filter map[string]interface{}, data *user.User) error {
	return common.DB.Where(filter).First(&data).Error
}

// Delete 批量删除
func (s UserService) Delete(ids []uint) error {
	// 用户和角色存在多对多关联关系
	var users []user.User
	for _, id := range ids {
		// 根据ID获取用户
		filter := tools.H{"id": id}

		user := new(user.User)
		err := s.Find(filter, user)
		if err != nil {
			return fmt.Errorf("获取用户信息失败，err: %v", err)
		}
		users = append(users, *user)
	}

	err := common.DB.Debug().Select("Roles").Unscoped().Delete(&users).Error
	if err != nil {
		return err
	}

	// 删除用户在group的关联
	err = common.DB.Debug().Exec("DELETE FROM group_users WHERE user_id IN (?)", ids).Error
	if err != nil {
		return err
	}

	return err
}

// GetUserByIds 根据用户ID获取用户角色排序最小值
func (s UserService) GetUserByIds(ids []uint) ([]user.User, error) {
	// 根据用户ID获取用户信息
	var userList []user.User
	err := common.DB.Where("id IN (?)", ids).Preload("Roles").Find(&userList).Error
	return userList, err
}

// ChangePwd 更新密码
func (s UserService) ChangePwd(username string, hashNewPasswd string) error {
	err := common.DB.Model(&user.User{}).Where("username = ?", username).Update("password", hashNewPasswd).Error
	// 如果更新密码成功，则更新当前用户信息缓存
	// 先获取缓存
	cacheUser, found := userInfoCache.Get(username)
	if err == nil {
		if found {
			user := cacheUser.(user.User)
			user.Password = hashNewPasswd
			userInfoCache.Set(username, user, cache.DefaultExpiration)
		} else {
			// 没有缓存就获取用户信息缓存
			var user user.User
			common.DB.Where("username = ?", username).Preload("Roles").First(&user)
			userInfoCache.Set(username, user, cache.DefaultExpiration)
		}
	}

	return err
}

// GetCurrentLoginUser 获取当前登录用户信息
// 需要缓存，减少数据库访问
func (s UserService) GetCurrentLoginUser(c *gin.Context) (user.User, error) {
	var newUser user.User
	ctxUser, exist := c.Get("user")
	if !exist {
		return newUser, errors.New("用户未登录")
	}
	u, _ := ctxUser.(user.User)

	// 先获取缓存
	cacheUser, found := userInfoCache.Get(u.Username)
	var user1 user.User
	var err error
	if found {
		user1 = cacheUser.(user.User)
		err = nil
	} else {
		// 缓存中没有就获取数据库
		user1, err = s.GetUserById(u.ID)
		// 获取成功就缓存
		if err != nil {
			userInfoCache.Delete(u.Username)
		} else {
			userInfoCache.Set(u.Username, user1, cache.DefaultExpiration)
		}
	}
	return user1, err
}

// Login 登录
func (s UserService) Login(user2 *user.User) (*user.User, error) {
	// 根据用户名获取用户(正常状态:用户状态正常)
	var firstUser user.User
	// err := common.DB.
	// 	Where("username = ?", user.Username).
	// 	Preload("Roles").
	// 	First(&firstUser).Error
	// if err != nil {
	// 	return nil, errors.New("用户不存在")
	// }
	err := s.Find(tools.H{"username": user2.Username}, &firstUser)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 判断用户拥有的所有角色的状态,全部角色都被禁用则不能登录
	// roles := firstUser.Roles
	// isValidate := false
	// for _, role := range roles {
	// 	// 有一个正常状态的角色就可以登录
	// 	if role.Status == 1 {
	// 		isValidate = true
	// 		break
	// 	}
	// }

	// if !isValidate {
	// 	return nil, errors.New("用户角色被禁用")
	// }

	if tools.NewParPasswd(firstUser.Password) != user2.Password {
		return nil, errors.New("密码错误")
	}

	// 校验密码
	// err = tools.ComparePasswd(firstUser.Password, user.Password)
	// if err != nil {
	// 	return &firstUser, errors.New("密码错误")
	// }
	return &firstUser, nil
}

// ClearUserInfoCache 清理所有用户信息缓存
func (s UserService) ClearUserInfoCache() {
	userInfoCache.Flush()
}

// GetUserById 获取单个用户
func (us UserService) GetUserById(id uint) (user.User, error) {
	var user user.User
	err := common.DB.Where("id = ?", id).First(&user).Error
	return user, err
}
