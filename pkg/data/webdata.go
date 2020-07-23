package data

import (
	"encoding/json"
	"fmt"
)

// UserDataToJSON returns json string from name
func UserDataToJSON(data UserData) string {
	buff, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return string(buff)
}
