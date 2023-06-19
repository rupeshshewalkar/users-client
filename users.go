package usersclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAllUsers GetAllAvengers returns list of Avengers
func (c Client) GetAllUsers() ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/getAllUsers", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// CreateAvenger will create an Avenger
func (c *Client) CreateUser(user User) (*User, error) {
	avg, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/createNewUser", c.HostURL), strings.NewReader(string(avg)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}
	var insertedID InsertedResult
	err = json.Unmarshal(body, &insertedID)
	if err != nil {
		return nil, err
	}
	user.ID = insertedID.InsertedID
	return &user, nil
}

// UpdateAvengerByName will update an Avenger
func (c *Client) UpdateUserByUserName(user User) (*UpdateResult, error) {
	avg, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/updateUserByUserName", c.HostURL), strings.NewReader(string(avg)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var updateResult UpdateResult
	err = json.Unmarshal(body, &updateResult)
	if err != nil {
		return nil, err
	}

	return &updateResult, nil
}

// DeleteAvengerByName will delete an Avenger
func (c *Client) DeleteUserByUserName(userName string) (*DeleteResult, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/users/deleteUserByUserName", c.HostURL), http.NoBody)
	req.URL.Query().Add("name", userName)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var deleteResult DeleteResult
	err = json.Unmarshal(body, &deleteResult)
	if err != nil {
		return nil, err
	}

	return &deleteResult, nil
}
