package middleware

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/config"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/model/response"
	"github.com/yin-zt/go-bind-mp/pkg/model/user"
	"github.com/yin-zt/go-bind-mp/pkg/service/isql"
	"github.com/yin-zt/go-bind-mp/pkg/util/tools"
	"time"
)

// 初始化jwt中间件
func InitAuth() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           config.Conf.Jwt.Realm,                                 // jwt标识
		Key:             []byte(config.Conf.Jwt.Key),                           // 服务端密钥
		Timeout:         time.Hour * time.Duration(config.Conf.Jwt.Timeout),    // token过期时间
		MaxRefresh:      time.Hour * time.Duration(config.Conf.Jwt.MaxRefresh), // token最大刷新时间(RefreshToken过期时间=Timeout+MaxRefresh)
		PayloadFunc:     payloadFunc,                                           // 有效载荷处理
		IdentityHandler: identityHandler,                                       // 解析Claims
		Authenticator:   login,                                                 // 校验token的正确性, 处理登录逻辑
		Authorizator:    authorizator,                                          // 用户登录校验成功处理,可以判断判断当前JWT是否有权限继续访问
		Unauthorized:    unauthorized,                                          // 用户登录校验失败处理
		LoginResponse:   loginResponse,                                         // 登录成功后的响应
		LogoutResponse:  logoutResponse,                                        // 登出后的响应
		RefreshResponse: refreshResponse,                                       // 刷新token后的响应
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",    // 自动在这几个地方寻找请求中的token
		TokenHeadName:   "Bearer",                                              // header名称
		TimeFunc:        time.Now,
	})
	return authMiddleware, err
}

// 有效载荷处理
func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(tools.H); ok {
		var user user.User
		// 将用户json转为结构体
		tools.JsonI2Struct(v["user"], &user)
		return jwt.MapClaims{
			jwt.IdentityKey: user.ID,
			"user":          v["user"],
		}
	}
	return jwt.MapClaims{}
}

// 解析Claims
func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	// 此处返回值类型map[string]interface{}与payloadFunc和authorizator的data类型必须一致, 否则会导致授权失败还不容易找到原因
	return tools.H{
		"IdentityKey": claims[jwt.IdentityKey],
		"user":        claims["user"],
	}
}

// 校验token的正确性, 处理登录逻辑
func login(c *gin.Context) (interface{}, error) {
	var req request.RegisterAndLoginReq
	// 请求json绑定
	if err := c.ShouldBind(&req); err != nil {
		return "", err
	}

	// 密码通过RSA解密
	decodeData, err := tools.RSADecrypt([]byte(req.Password), config.Conf.System.RSAPrivateBytes)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Username: req.Username,
		Password: string(decodeData),
	}

	// 密码校验
	user, err := isql.User.Login(u)
	if err != nil {
		return nil, err
	}
	// 将用户以json格式写入, payloadFunc/authorizator会使用到
	return tools.H{
		"user": tools.Struct2Json(user),
	}, nil
}

// 用户登录校验成功处理
func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(tools.H); ok {
		userStr := v["user"].(string)
		var user user.User
		// 将用户json转为结构体
		tools.Json2Struct(userStr, &user)
		// 将用户保存到context, api调用时取数据方便
		c.Set("user", user)
		return true
	}
	return false
}

// 用户登录校验失败处理
func unauthorized(c *gin.Context, code int, message string) {
	log.Error("JWT认证失败, 错误码: %d, 错误信息: %s", code, message)
	response.Response(c, code, code, nil, fmt.Sprintf("JWT认证失败, 错误码: %d, 错误信息: %s", code, message))
}

// 登录成功后的响应
func loginResponse(c *gin.Context, code int, token string, expires time.Time) {
	response.Response(c, code, code,
		gin.H{
			"token":   token,
			"expires": expires.Format("2006-01-02 15:04:05"),
		},
		"登录成功")
}

// 登出后的响应
func logoutResponse(c *gin.Context, code int) {
	response.Success(c, nil, "退出成功")
}

// 刷新token后的响应
func refreshResponse(c *gin.Context, code int, token string, expires time.Time) {
	response.Response(c, code, code,
		gin.H{
			"token":   token,
			"expires": expires,
		},
		"刷新token成功")
}
