// decode json using sructs
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Marshal encode to json
// UnMarshal decode from json to value
////read users. json using scan ->[]byte -> slice of struct using Unmarshal
/* cat users.json
[
 {
  "Firstname": "usr1"
 },
 {
  "Firstname": "usr2",
  "Permission": {
   "write": true
  }
 },
 {
  "Firstname": "usr3",
  "Permission": {
   "write": true
  }
 },
 {
  "Firstname": "usr4"
 }
]
*/ //go run unmarshal.go<users.json
type user struct {
	Fname string `json:"Firstname"` //json.field name should match with users.json
	//pwd        string
	Permmision map[string]bool `json:"perms"` //json.field name should match with users.json
}

func main() {

	var users []user
	var slc []byte
	for rd := bufio.NewScanner(os.Stdin); rd.Scan(); {
		//fmt.Println(string(rd.Bytes()))
		//fmt.Println(string(rd.Text()))
		slc = append(slc, rd.Bytes()...) //
		//slc = append(slc, "user"...)
		//slc = append(slc, rd.Text()...)
	}
	//err := json.Unmarshal(slc, &users) // decode json from slc into users slice of struct
	if err := json.Unmarshal(slc, &users); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(users)
	for _, i := range users {
		//fmt.Println(i.Fname, i.Permmision)
		switch p := i.Permmision; {
		case p["admin"]:
			fmt.Println(i.Fname, " is admin")
		case p == nil:
			fmt.Println(i.Fname, " has no permission")

		}
	}
}
