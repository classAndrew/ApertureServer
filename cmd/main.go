package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/classAndrew/ApertureServer/pkg/server"

	"github.com/classAndrew/ApertureServer/pkg/data"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	user := func(w *http.ResponseWriter, req *http.Request, user string) {
		io.WriteString(*w, data.UserDataToJSON(data.DataMngr.GetUser(user))+"\n")
	}

	register := func(w *http.ResponseWriter, req *http.Request, name string) {
		// randPlanet := data.DataMngr.RandomStarSystem() // Get random star system instead
		sys := server.GenerateStarSystem()
		data.DataMngr.RegisterStarSystem(&sys)
		user := server.PM.CreateUser(name)
		homeName := data.DataMngr.SetRandomPlanetAttribute("Owner", user.Name)
		user.HomePlanet = homeName
		user.CelestialBodies = append(user.CelestialBodies, homeName)
		res, status := data.DataMngr.RegisterUser(user)
		if status {
			io.WriteString(*w, res+"\n")
			return
		}
		io.WriteString(*w, "Username Already Registered / Something went wrong\n")
	}

	system := func(w *http.ResponseWriter, req *http.Request, name string) {
		io.WriteString(*w, data.StarSystemToJSON(data.DataMngr.GetStarSystem(name))+"\n")
		sys := data.DataMngr.GetStarSystem(name)
		io.WriteString(*w, data.StarSystemToJSON(sys)+"\n")
	}

	root := func(w http.ResponseWriter, req *http.Request) {
		split := strings.Split(req.URL.String(), "/")
		if len(split) > 1 {
			switch strings.ToLower(split[1]) {
			case "user":
				if len(split) == 3 {
					user(&w, req, split[2])
				}
				break
			case "register":
				if len(split) == 3 {
					register(&w, req, split[2])
				}
				break
			case "system":
				if len(split) == 3 {
					system(&w, req, split[2])
				}
				break
			default:
				io.WriteString(w, "Aperture API\n")
			}
		} else {
			io.WriteString(w, "Aperture API\n")
		}
	}
	http.HandleFunc("/", root)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
