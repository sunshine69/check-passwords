package main

import (
	"log"
	"testing"

	u "github.com/sunshine69/golang-tools/utils"
)

func TestCheckPasswords(t *testing.T) {
	o := CheckPasswords([]string{""})
	u.Assert(len(o) == 0, "OK, compromized", true)
	log.Println(o)
}
