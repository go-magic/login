package login

import (
	"fmt"
	"testing"
)

func TestGetPhoneCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(getPhoneCode())
	}
}
