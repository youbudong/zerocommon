package encodex

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(password string) (string, error) {
	// 生成 salt 值，使用默认的 10 为 cost 参数
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	// 返回生成的哈希值
	return hex.EncodeToString(salt), nil
}

func CheckPassword(password string, hashString string) bool {
	// 使用哈希值校验密码
	hash, err := hex.DecodeString(hashString)
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
