package main

import (
	"encoding/json"
	"fmt"
)

// userDataToJSON returns json string from name
func userDataToJSON(data UserData) string {
	buff, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return string(buff)
}
