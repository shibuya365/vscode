package dbscs // 独自の設定ファイルパッケージ

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Sc はDB設定の構造体
type Shortcut struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Shortcut    string `json:"shortcut"`
	Description string `json:"description"`
	Visiable    bool
	// `json:"visiable"`
}

// Shortcuts はScの配列
type Shortcuts []Shortcut

// ReadShortcutsDB はDB設定読み込み関数
func ReadShortcutsDB() (scs Shortcuts) {
	jsonFromFile, err := ioutil.ReadFile("dbscs/shortcuts.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(jsonFromFile, &scs)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("scs[0].Visiable: ", scs[0].Visiable)
	// for i := 0; i < len(scs); i++ {
	// 	scs[i].Visiable = true
	// }
	return scs
}
