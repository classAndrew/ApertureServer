package main

import (
	"io"
	"math/big"
	"net/http"
	"strings"
)

func main() {
	monHand := NewMongoHandler("localhost", "27017")
	dataMngr := DataManager{128, 10, big.ToZero, monHand}
	user := func(w *http.ResponseWriter, req *http.Request, user string) {
		io.WriteString(*w, userDataToJSON(dataMngr.GetUser(user))+"\n")
	}

	root := func(w http.ResponseWriter, req *http.Request) {
		split := strings.Split(req.URL.String(), "/")
		if len(split) > 1 {
			if strings.ToLower(split[1]) == "user" && len(split) == 3 {
				user(&w, req, split[2])
			} else {
				io.WriteString(w, "GoCurrency API\n")
			}
		} else {
			io.WriteString(w, "GoCurrency API\n")
		}
	}

	http.HandleFunc("/", root)
	http.ListenAndServe(":8000", nil)
}
