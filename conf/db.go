package conf // 独自の設定ファイルパッケージ

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// type ConfDB struct {
// 	Host   string `json:"host"`    // ホスト名
// 	Port   int    `json:"port"`    // ポート番号
// 	DbName string `json:"db-name"` // 接続先DB名
// 	User   string `json:"user"`    // 接続ユーザ名
// 	Pass   string `json:"pass"`    // 接続パスワード
// }

// DB設定の構造体
type Line struct {
	Title    string `json:"title"`
	Shortcut string `json:"shortcut"`
}

type Lines []Line

// DB設定読み込み関数
func ReadConfDB() (Lines, error) {
	jsonFromFile, err := ioutil.ReadFile("conf/shortcut.json")
	if err != nil {
		fmt.Println(err)
	}

	var jsonData Lines
	err = json.Unmarshal(jsonFromFile, &jsonData)
	if err != nil {
		return jsonData, err
	}

	// fmt.Println(jsonData)

	// JSON => 構造体

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
	return jsonData, nil
}
