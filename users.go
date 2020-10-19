/*
 * Copyright (c) 2020 Rollbar, Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package client

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"strconv"
)

// User represents a Rollbar user.
type User struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	Email    string `json:"email"`
}

// ListUsers : A function for listing the users.
func (c *Client) ListUsers() (*ListUsersResponse, error) {
	var data ListUsersResponse

	bytes, err := c.get("users")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// This response doesn't have pagination so it might break
// in the future.
func (c *Client) getID(email string) (int, error) {
	var userID int

	l, err := c.ListUsers()
	if err != nil {
		return 0, err
	}

	for _, user := range l.Result.Users {
		if user.Email == email {
			userID = user.ID
		}
	}

	return userID, nil
}

// GetUser fetches one user.
func (c *Client) GetUser(email string) (int, error) {
	userID, err := c.getID(email)
	if err != nil {
		return 0, fmt.Errorf("There was a problem with getting the user id %s", err)
	}
	return userID, nil

}

// RemoveUserTeam removes a user from a team.
func (c *Client) RemoveUserTeam(email string, teamID int) error {
	userID, err := c.GetUser(email)
	if err != nil {
		return err
	}

	err = c.delete("team", strconv.Itoa(teamID), "user", strconv.Itoa(userID))
	if err != nil {
		return err
	}

	return nil
}

/*
 * Containers for unmarshalling Rollbar API responses
 */

type listUsersResponse struct {
	Error  int    `json:"err"`
	Result []User `json:"result"`
}
