package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	NumRepos string `json:""`
}

// const (
// 	ErrUserNotFound   error.New("User not found!")
// 	ErrInvalidRequest error.New("Invalid Request")
// )

//UserInfo: return information of the user!

func UserInfo(login string) (*User, error) {
	// call the github API for a given login
	// and return the user struct
	if login == "" {
		return nil, fmt.Errorf("error - please pass a username")
	}
	endpoint := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("error making request - %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making request - %s", err.Error())
	}
	return nil, nil
}

//return slice of informations
func GetUserFollowersInformation(login string) (*User, error) {
	return nil, nil
}

func do() {

	// if len(os.Args) < 2 {
	// 	log.Fatal("Please provide argumentas!")
	// }
	user, err := UserInfo("mrbomber0x001")
	if err != nil {
		log.Fatal("error")
	}
	fmt.Println(user, err)
}
