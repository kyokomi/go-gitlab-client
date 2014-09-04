package gogitlab

import (
	"encoding/json"
)

const (
	usersURL       = "/users"     // Get users list
	userURL        = "/users/:id" // Get a single user.
	currentUserURL = "/user"      // Get current user
)

// User ...
type User struct {
	ID            int    `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	Name          string `json:"name,omitempty"`
	State         string `json:"state,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	Bio           string `json:"bio,omitempty"`
	Skype         string `json:"skype,omitempty"`
	LinkedIn      string `json:"linkedin,omitempty"`
	Twitter       string `json:"twitter,omitempty"`
	ExternUID     string `json:"extern_uid,omitempty"`
	Provider      string `json:"provider,omitempty"`
	ThemeID       int    `json:"theme_id,omitempty"`
	ColorSchemeID int    `json:"color_scheme_id,color_scheme_id"`
}

// Users a All User
func (g *Gitlab) Users() ([]*User, error) {

	url := g.ResourceUrl(usersURL, nil)

	var users []*User

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &users)
	}

	return users, err
}

/*
User Get a single user.

    GET /users/:id

Parameters:

    id The ID of a user

Usage:

	user, err := gitlab.User("your_user_id")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", user)
*/
func (g *Gitlab) User(id string) (*User, error) {

	url := g.ResourceUrl(userURL, map[string]string{":id": id})

	user := new(User)

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, err
}

// DeleteUser DELETE a user by the authenticated user.
func (g *Gitlab) DeleteUser(id string) error {
	url := g.ResourceUrl(userURL, map[string]string{":id": id})
	var err error
	_, err = g.buildAndExecRequest("DELETE", url, nil)
	return err
}

// CurrentUser GET a current user by the authenticated user.
func (g *Gitlab) CurrentUser() (User, error) {
	url := g.ResourceUrl(currentUserURL, nil)
	var user User

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, err
}
