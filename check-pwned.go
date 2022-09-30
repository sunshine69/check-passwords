package main

import (
	"fmt"
	"strings"

	u "github.com/sunshine69/golang-tools/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	haveibeenpwned_api_url = "https://api.pwnedpasswords.com"
)

func (a *App) CheckPasswords(passText string) string {
	passList := strings.Split(passText, "\n")
	//fmt.Println("passList", passList)
	output := []string{}
	for _, pass := range passList {
		sha1text := strings.ToUpper(u.Sha1Sum(pass))
		// fmt.Println( sha1text )
		o := CheckPasswordByHash(sha1text)
		if o != "" {
			output = append(output, "'"+pass+"'"+" | appear count: "+o)
		}
	}
	if len(output) == 0 {
		_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:   "It's OK",
			Message: "OK no compromised found.",
			// Buttons: []string{"one", "two", "three", "four"},
		})
		u.CheckErr(err, "MessageDialog")
		return "OK no compromised found."
	}
	msg := "ERROR Your passwords have been found in haveibeenpwned database. Change the password ASSAP. See below for each compromized pass and how many times it appear to be compormised\r\n\r\n" + strings.Join(output, "\r\n")
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Title:   "It's NOT OK",
		Message: msg,
		// Buttons: []string{"one", "two", "three", "four"},
	})

	u.CheckErr(err, "MessageDialog")
	return "ERROR"
}

func CheckPasswordByHash(hash string) string {
	hash_prefix := hash[0:5]
	o, err := u.Curl("GET", fmt.Sprintf("%s/range/%s", haveibeenpwned_api_url, hash_prefix), "", "", []string{})
	u.CheckErr(err, "Can not fetch data from api")
	res_list := strings.Split(o, "\n")
	for _, r := range res_list {
		if r == "" {
			continue
		}
		_tlist := strings.Split(r, ":")
		hash_suffix, appear_count := _tlist[0], _tlist[1]
		if hash == hash_prefix+hash_suffix {
			return appear_count
		}
	}
	return ""
}
