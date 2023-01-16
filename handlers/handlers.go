package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MelihEmreGuler/user-information-system/dataloaders"
	"github.com/MelihEmreGuler/user-information-system/models"
)

func Run() {
	http.HandleFunc("/users", UsersHandler)
	http.ListenAndServe(":8080", nil)

}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	//pages
	page := models.Page{
		ID:          7,
		Name:        "Users",
		Description: "User List",
		URI:         "/users",
	}

	//data loaders
	users := dataloaders.LoadUsers()
	interests := dataloaders.LoadInterests()
	interestMappings := dataloaders.LoadinterestsMappings()

	//process data
	var newUsers []models.User
	for _, user := range users {
		var newInterests []models.Interest
		for _, interestMapping := range interestMappings {
			if interestMapping.UserID == user.ID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						newInterests = append(newInterests, interest)
					}
				}
			}
		}
		user.Interests = newInterests
		newUsers = append(newUsers, user)
	}

	viewModel := models.UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel) //convert to json
	w.Write([]byte(data))              //write to response

}
