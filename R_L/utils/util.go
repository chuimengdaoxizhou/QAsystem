package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"qasystem/R_L/models"
	"qasystem/config"
	"time"
)

func GenerateToken(userID int64, username string) (string, error) {
	claims := &models.Claims{
		UserID:   userID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)), // 设置过期时间为 12 小时后
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 设置签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 设置生效时间
			Issuer:    "QASystem",                                         // 设置签发者
		},
	}

	// 使用指定的签名方式获取Token
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 从配置文件中获取jwtSecret
	con := config.GetConfig()
	jwtSecret := con.GetsecretKey()

	// 使用密钥签名 Token 并获取完整编码后的字符串 token
	signedToken, err := Token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func ParseJWT(tokenString string) (string, error) {
	con := config.GetConfig()
	secertKey := con.GetsecretKey()

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected Signing Method")
		}
		return []byte(secertKey), nil
	})
	if err != nil { // 错误
		fmt.Println("err", err)
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("Username claim is not a string")
		}
		return username, nil
	}
	return "", err
}
func GetHashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashpassword), err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
