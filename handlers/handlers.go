package handlers

import (
	"encoding/json"
	"net/http"

	. "../dataloaders"
	. "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	//pages
	page := Page{
		ID:          7,
		Name:        "Users",
		Description: "User List",
		URI:         "/users",
	}

	//data loaders
	users := LoadUsers()
	interests := LoadInterests()
	interestMappings := LoadinterestsMappings()

	//process data
	var newUsers []User
	for _, user := range users {
		var newInterests []Interest
		for _, interest := range interests {
			for _, interestMapping := range interestMappings {
				if interestMapping.UserID == user.ID && interestMapping.InterestID == interest.ID {
					newInterests = append(newInterests, interest)
				}
			}
		}
		user.Interests = newInterests
		newUsers = append(newUsers, user)
	}

	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel) //convert to json
	w.Write(data)					 //write to response

}
