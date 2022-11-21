package model

type User struct {
	UserID        int    `json:"userid,omitempty`
	FirstName     string `json:"firstname,omitempty`
	LastName      string `json:"lastname,omitempty`
	Email         string `json:"email,omitempty`
	Phone         int    `json:"phone,omitempty`
	Status        bool   `json:"status,omitempty`
	Password      string `json:"password,omitempty`
	LastUpdatedAt string `json:"lastupdatedat,omitempty`
}
