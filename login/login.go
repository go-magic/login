package login

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	//cookie := c.GetHeader("set_cookies")
	//if cookie != "" {
	//
	//}
}

func verifyCookies(cookie string) bool {
	userinfo, _ := base64.StdEncoding.DecodeString(cookie)
	arr := strings.Split(string(userinfo), ":")
	if len(arr) != 2 {
		return false
	}
	return true
}
