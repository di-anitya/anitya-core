package apiserver

import (
	"fmt"

	"../models"
)

var currentID int
var users models.Users

// Give us some seed data
func init() {
	RepoCreateUser(models.User{Name: "Write presentation"})
	RepoCreateUser(models.User{Name: "Host meetup"})
}

// RepoFindUser is ..
func RepoFindUser(id int) models.User {
	for _, t := range users {
		if t.ID == id {
			return t
		}
	}
	// return empty User if not found
	return models.User{}
}

// RepoCreateUser is ..
func RepoCreateUser(t models.User) models.User {
	currentID++
	t.ID = currentID
	users = append(users, t)
	return t
}

// RepoDestroyUser is ..
func RepoDestroyUser(id int) error {
	for i, t := range users {
		if t.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}
