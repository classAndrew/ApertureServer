package data

import (
	"encoding/json"
	"fmt"

	"github.com/classAndrew/ApertureServer/pkg/server"
)

// Perhaps move all of this into an interface method?

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

// PlanetToJSON return json string of a planet
func PlanetToJSON(data *server.Planet) string {
	buff, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return string(buff)
}
