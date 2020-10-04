package client

import (
//"encoding/json"
//"strconv"
)

// ProjectAccessToken represents a project access token
type ProjectAccessToken struct {
	ProjectID    int    `json:"project_id"`
	AccessToken  string `json:"access_token"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	DateCreated  int    `json:"date_created"`
	DateModified int    `json:"date_modified"`
}

// listProjectAccessTokensResponse represents the list-project-access-tokens response
type listProjectAccessTokensResponse struct {
	Error  int `json:"err"`
	Result []ProjectAccessToken
}

// ListProjectAccessTokens lists the projects access tokens for a given project id
// https://docs.rollbar.com/reference#list-all-project-access-tokens
func (c *RollbarApiClient) ListProjectAccessTokens(projectID int) ([]ProjectAccessToken, error) {
	return []ProjectAccessToken{}, nil
	/*
		var data listProjectAccessTokensResponse

		bytes, err := c.get("project", strconv.Itoa(projectID), "access_tokens")
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return nil, err
		}

		tokens := []*ProjectAccessToken{}
		tokens = append(tokens, data.Result...)

		return tokens, nil

	*/
}

// GetProjectAccessTokenByProjectIDAndName returns the first project access token from the
// list-projects-tokens call that matches a given name nad project id. If there is no
// matching project, return nil.
func (c *RollbarApiClient) GetProjectAccessTokenByProjectIDAndName(projectID int, name string) (ProjectAccessToken, error) {
	return ProjectAccessToken{}, nil
	/*
		tokens, err := c.ListProjectAccessTokens(projectID)
		if err != nil {
			return nil, err
		}

		for _, t := range tokens {
			if t.Name == name {
				return t, nil
			}
		}

		return nil, nil

	*/
}
