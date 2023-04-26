package main

//
// //TODO: use /pkg/errors and errors.Wrap() to capture stack trace!
// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// )

// // Print the errors with their stack trace!
// func setupLoggin() {
// 	out, err := os.OpenFile("github_api.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return
// 	}

// 	log.SetOutput(out)
// }

// type User struct {
// 	Login    string `json:"login"`
// 	Name     string `json:"name"`
// 	NumRepos int    `json:"public_repos"`
// }

// // const (
// // 	ErrUserNotFound   error.New("User not found!")
// // 	ErrInvalidRequest error.New("Invalid Request")
// // )

// //UserInfo: return information of the user!
// func UserInfo(login string) (*User, error) {
// 	// call the github API for a given login
// 	// and return the user struct
// 	if login == "" {
// 		return nil, fmt.Errorf("error - please pass a username")
// 	}
// 	endpoint := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))
// 	resp, err := http.Get(endpoint)

// 	if err != nil {
// 		return nil, fmt.Errorf("error making request - %s", err.Error())
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("error making request - %s", err.Error())
// 	}

// 	// Decode JSON: Should be placed on it's own function
// 	user := User{Login: login}
// 	dec := json.NewDecoder(resp.Body)
// 	if err := dec.Decode(&user); err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// //return slice of informations
// // func GetUserFollowersInformation(login string) (*User, error) {
// // 	return nil, nil
// // }

// func main() {

// 	// if len(os.Args) < 2 {
// 	// 	log.Fatal("Please provide argumentas!")
// 	// }
// 	setupLoggin()
// 	user, err := UserInfo("mrbomber0x001")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s\n", err)
// 		log.Printf("error: %+v", err) // will use setupLogging();
// 		os.Exit(1)
// 	}
// 	fmt.Println(user, err)
// }
