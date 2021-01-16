package dbus // 独自の設定ファイルパッケージ

import (
	"encoding/gob"
	"fmt"
	"os"
)

// ReadUsersDB はDB読み込み関数
func ReadUsersDB() (users map[string][]string) {
	// ファイルを開く
	f, err := os.Open("dbus/save.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		// なかった場合の仮のデータ作成
		var strs []string
		users["bvtujotnf4q4d12u9700"] = strs
		fmt.Println("Make initial users: ", users)
	}

	dec := gob.NewDecoder(f)
	if err := dec.Decode(&users); err != nil {
		fmt.Println(err)

		// なかった場合の仮のデータ作成
		var strs []string
		users["bvtujotnf4q4d12u9700"] = strs
		fmt.Println("Make initial users: ", users)
	}

	return users
}
