package login

import (
	"dat2/db"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID      int      `json:"user_id"`
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
}

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

func AuthToken(c *gin.Context) {
	strToken := c.GetHeader("token")
	claim, err := verifyAction(strToken)
	if err == nil {
		c.Next()
		return
	}
	fmt.Println(claim)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !db.AuthUserInfo(username, password) {
		c.Abort()
		c.JSON(http.StatusUnauthorized, &struct{}{})
		return
	}
	claims := &JWTClaims{
		Username:    username,
		Password:    password,
		FullName:    username,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err := getToken(claims)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusUnauthorized, &struct{}{})
		return
	}
	c.JSON(http.StatusOK, &gin.H{
		"token": signedToken,
	})
}

func refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := getToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

func verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New("error")
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("error")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("error")
	}
	return claims, nil
}

func getToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New("error")
	}
	return signedToken, nil
}
