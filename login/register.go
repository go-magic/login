package login

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RegisterByAccount(c *gin.Context) {
	c.PostForm("username")
	c.PostForm("password")
}

func RegisterByPhone(c *gin.Context) {

}

func RegisterByEmail(c *gin.Context) {

}

func getPhoneCode() int {
	h := rand.Intn(8) + 1
	l := rand.Intn(99999)
	return h*100000 + l
}
