package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试userdao中的函数")
	t.Run ("验证用户名和密码：", testLogin)
	t.Run ("验证用户名：", testRegist)
	t.Run ("保存用户：", testSave)
}

func testLogin(t *testing.T) {
	USER, _ := CheckIDAndPassword("","123456")
	fmt.Println("获取客户信息为：", user)
}

func testRegist(t *testing.T) {
	USER, _ := CheckID("")
	fmt.Println("获取客户信息为：", user)
}

func testSave(t *testing.T) {
	Saveuser("", "", "", "")
}