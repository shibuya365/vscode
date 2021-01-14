package dbus // 独自の設定ファイルパッケージ

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Users ユーザー
// type Users map[string][]int

// ReadUsersDB はDB設定読み込み関数
func ReadUsersDB() (users map[string][]int) {
	// var users Users
	f, err := os.Open("dbus/save.txt")
	if err != nil {
		fmt.Println(err)
		n := []int{1, 2}
		users["bvtujotnf4q4d12u9700"] = n
		return users
	}
	defer f.Close()

	dec := gob.NewDecoder(f)
	if err := dec.Decode(&users); err != nil {
		fmt.Println(err)
	}

	return users
}

// jsonFromFile, err := ioutil.ReadFile("dbus/users.json")
// if err != nil {
// 	fmt.Println("users: ", users)
// 	return users
// }

// err = json.Unmarshal(jsonFromFile, &users)
// if err != nil {
// 	fmt.Println(err)
// }
// return users
