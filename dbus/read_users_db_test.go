package dbus

import (
	"fmt"
	"testing"
)

func TestReadUsersDB(t *testing.T) {
	users := ReadUsersDB()
	isBool := users["bvtujotnf4q4d12u9700"][0]
	fmt.Println("isBool: ", isBool)
	if isBool != false {
		t.Fatal("failed test")
	}
}
