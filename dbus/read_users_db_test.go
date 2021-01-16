package dbus

import (
	"fmt"
	"testing"
)

func TestReadUsersDB(t *testing.T) {
	users := ReadUsersDB()
	strs := users["bvtujotnf4q4d12u9700"]
	fmt.Println("isBool: ", strs)
	if strs[0] != "" {
		t.Fatal("failed test")
	}
}
