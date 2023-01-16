package dataloaders

import (
	"encoding/json"

	"github.com/MelihEmreGuler/user-information-system/models"
	"github.com/MelihEmreGuler/user-information-system/utils"
)

func LoadUsers() []models.User {
	bytes, err := utils.ReadFile("json/user.json")
	if err != nil {
		panic(err)
	}
	//fmt.Println("LoadUsers:\n" + bytes)
	var data []models.User
	json.Unmarshal([]byte(bytes), &data)

	return data
}
func LoadInterests() []models.Interest {
	bytes, err := utils.ReadFile("json/interests.json")
	if err != nil {
		panic(err)
	}
	//fmt.Println("LoadInterests:\n" + bytes)
	var data []models.Interest
	json.Unmarshal([]byte(bytes), &data)

	return data
}
func LoadinterestsMappings() []models.InterestMapping {
	bytes, err := utils.ReadFile("json/userInterestsMappings.json")
	if err != nil {
		panic(err)
	}
	//fmt.Println("LoadinterestsMappings:\n" + bytes)
	var data []models.InterestMapping
	json.Unmarshal([]byte(bytes), &data)

	return data
}
