// o/p go run test.go>users.json Marshal/encode  struct object to json
package main

import (
	"encoding/json"
	"fmt"
)

type permission map[string]bool
type name struct {
	Fname      string     `json:"Firstname"`       // encode field with different name
	Pswd       string     `json:"-"`               /// hide paswd using json field tag
	Permission permission `json:"perms,omitempty"` //// no encoding for permission if they are empty/null
}

func main() {

	//var mp map[string]bool  // error panic: assignment to entry in nil map
	perm := make(map[string]bool)
	//mp["read"] = true

	names := []name{
		{"usr1", "g", nil},
		{"usr2", "g", map[string]bool{"admin": true}},
		{"usr3", "g", permission{"write": true}},
		{"usr4", "g", perm},
	}
	slc, err := json.MarshalIndent(names, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(slc))

	// fmt.Println(mp)
}
