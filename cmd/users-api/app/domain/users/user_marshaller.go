package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Status      string `json:"status"`
}

func (u *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          u.Id,
			DateCreated: u.DateCreated,
			DateUpdated: u.DateUpdated,
			Status:      u.Status,
		}
	}

	userJson, _ := json.Marshal(u)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}

	return privateUser
}

func (u Users) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, len(u))
	for i, user := range u {
		result[i] = user.Marshall(isPublic)
	}
	return result
}
