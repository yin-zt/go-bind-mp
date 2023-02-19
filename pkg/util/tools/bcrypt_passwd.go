package tools

import (
	"fmt"
	"github.com/yin-zt/go-bind-mp/pkg/config"
)

// 密码加密 使用自适应hash算法, 不可逆
// func GenPasswd(passwd string) string {
// 	hashPasswd, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
// 	return string(hashPasswd)
// }

// 通过比较两个字符串hash判断是否出自同一个明文
// hashPasswd 需要对比的密文
// passwd 明文
// func ComparePasswd(hashPasswd string, passwd string) error {
// 	// if err := bcrypt.CompareHashAndPassword([]byte(hashPasswd), []byte(passwd)); err != nil {
// 	// 	return err
// 	// }

// 	return nil
// }

// 密码加密
func NewGenPasswd(passwd string) string {
	fmt.Println([]byte(passwd))
	fmt.Println(config.Conf.System.RSAPublicBytes)
	fmt.Println("oooooooooooooooooooo")
	pass, _ := RSAEncrypt([]byte(passwd), config.Conf.System.RSAPublicBytes)
	return string(pass)
}

// 密码解密
func NewParPasswd(passwd string) string {
	pass, _ := RSADecrypt([]byte(passwd), config.Conf.System.RSAPrivateBytes)
	return string(pass)
}
