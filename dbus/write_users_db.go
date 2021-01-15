package dbus // 独自の設定ファイルパッケージ

import (
	"encoding/gob"
	"fmt"
	"os"
)

// WriteUsersDB 関数
func WriteUsersDB(users map[string][]bool) {
	fmt.Println("write users: ", users)

	f, err := os.Create("dbus/save.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	enc := gob.NewEncoder(f)

	if err := enc.Encode(users); err != nil {
		fmt.Println(err)
	}
}
