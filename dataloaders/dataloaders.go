package dataloaders

import (
	"encoding/json"

	. "../models"
	util "../utils"
)

func LoadUsers() []User {
	bytes, _ := util.ReadFile("../json/user.json")

	var data []User
	json.Unmarshal([]byte(bytes), &data)

	return data
}
func LoadInterests() []Interest {
	bytes, _ := util.ReadFile("../json/interests.json")

	var data []Interest
	json.Unmarshal([]byte(bytes), &data)

	return data
}
func LoadinterestsMappings() []InterestMapping {
	bytes, _ := util.ReadFile("../json/userInterestsMappings.json")

	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)

	return data
}
