package data

import (
	"encoding/json"
	"fmt"

	"github.com/classAndrew/ApertureServer/pkg/server"
)

// UserDataToJSON returns json string from name
func UserDataToJSON(data server.UserData) string {
	buff, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return string(buff)
}

// StarSystemToJSON retusn json string from starsystem
func StarSystemToJSON(data *server.StarSystem) string {
	buff, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return string(buff)
}
