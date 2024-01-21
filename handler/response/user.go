package response

import "github.com/hnpatil/messages/entity"

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

func ToUserResponse(input *entity.User) *User {
	u := &User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	return u
}
