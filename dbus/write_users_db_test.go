package dbus

import "testing"

func TestWriteUsersDB(t *testing.T) {
	var users map[string][]bool
	n := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	users["bvtujotnf4q4d12u9700"] = n
	WriteUsersDB(users)
}
