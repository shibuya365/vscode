package dbus // 独自の設定ファイルパッケージ

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Users ユーザー
// type Users map[string][]int

// ReadUsersDB はDB設定読み込み関数
func ReadUsersDB() (users map[string][]bool) {
	// var users Users
	f, err := os.Open("dbus/save.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		users := make(map[string][]bool)
		var n []bool
		for i := 0; i < 62; i++ {
			n = append(n, false)
		}
		users["bvtujotnf4q4d12u9700"] = n
		return users
	}

	dec := gob.NewDecoder(f)
	if err := dec.Decode(&users); err != nil {
		fmt.Println(err)
		users := make(map[string][]bool)
		var n []bool
		for i := 0; i < 62; i++ {
			n = append(n, false)
		}
		users["bvtujotnf4q4d12u9700"] = n
		return users
	}

	return users
}
