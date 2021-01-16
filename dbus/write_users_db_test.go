package dbus

import (
	"fmt"
	"testing"
)

func TestWriteUsersDB(t *testing.T) {
	var users map[string][]string
	// なかった場合の仮のデータ作成
	var strs []string
	fmt.Println("read_users_db strs: ", strs)
	users["bvtujotnf4q4d12u9700"] = strs
	fmt.Println("read_users_db users: ", users)
	WriteUsersDB(users)
}
