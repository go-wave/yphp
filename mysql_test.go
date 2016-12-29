package yphp

import "testing"

func Test_Mysql_escape_string(t *testing.T) {
	str := "select * from table where content = '\"s'"
	real := Mysql_escape_string(str)
	if real != "select * from table where content = \\'\\\"s\\'" {
		t.Error("escape string error:", real)
	}
}
