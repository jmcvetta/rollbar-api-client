package client

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

// TestListProjectAccessTokens tests listing Rollbar project access tokens.
func (s *Suite) TestListProjectAccessTokens() {
	projectID := 12116
	u := apiUrl + pathPatList
	u = strings.ReplaceAll(u, "{projectId}", strconv.Itoa(projectID))

	r := responderFromFixture("project_access_token/list.json", http.StatusOK)
	httpmock.RegisterResponder("GET", u, r)

	// Valid project ID
	expected := []ProjectAccessToken{
		{
			AccessToken: "80f235b890c34ca49bcea692c2b90421",
			//"cur_rate_limit_window_count": null,
			//"cur_rate_limit_window_start": null,
			DateCreated:          1601982124,
			DateModified:         1601982124,
			Name:                 "post_client_item",
			ProjectID:            411334,
			RateLimitWindowCount: nil,
			RateLimitWindowSize:  nil,
			Scopes: []Scope{
				ScopePostClientItem,
			},
			Status: "enabled",
		},
		{
			AccessToken: "8d4b7e0e6a1a498db82cffd1eda93376",
			//"cur_rate_limit_window_count": null,
			//"cur_rate_limit_window_start": null,
			DateCreated:          1601982124,
			DateModified:         1601982124,
			Name:                 "post_server_item",
			ProjectID:            411334,
			RateLimitWindowCount: nil,
			RateLimitWindowSize:  nil,
			Scopes: []Scope{
				ScopePostServerItem,
			},
			Status: "enabled",
		},
		{
			AccessToken: "90b2521327a647f9aa80ef6d84427485",
			//"cur_rate_limit_window_count": null,
			//"cur_rate_limit_window_start": null,
			DateCreated:          1601982124,
			DateModified:         1601982124,
			Name:                 "read",
			ProjectID:            411334,
			RateLimitWindowCount: nil,
			RateLimitWindowSize:  nil,
			Scopes: []Scope{
				ScopeRead,
			},
			Status: "enabled",
		},
		{
			AccessToken: "d6d4b456f72048dfb8a933afe3ac66f6",
			//"cur_rate_limit_window_count": null,
			//"cur_rate_limit_window_start": null,
			DateCreated:          1601982124,
			DateModified:         1601982124,
			Name:                 "write",
			ProjectID:            411334,
			RateLimitWindowCount: nil,
			RateLimitWindowSize:  nil,
			Scopes: []Scope{
				ScopeWrite,
			},
			Status: "enabled",
		},
	}
	actual, err := s.client.ListProjectAccessTokens(projectID)
	s.Nil(err)
	s.Equal(expected, actual)

	// Unreachable server
	httpmock.Reset()
	_, err = s.client.ListProjectAccessTokens(projectID)
	s.NotNil(err)
	s.NotEqual(ErrNotFound, err)

	// Unauthorized
	r = httpmock.NewJsonResponderOrPanic(http.StatusUnauthorized,
		ErrorResult{Err: 401, Message: "Unauthorized"})
	httpmock.RegisterResponder("GET", u, r)
	_, err = s.client.ListProjectAccessTokens(projectID)
	s.Equal(ErrUnauthorized, err)
}

// TestReadProjectAccessToken tests reading a Rollbar project access token from
// the API.
func (s *Suite) TestReadProjectAccessToken() {
	projectID := 411334
	u := apiUrl + pathPatList
	u = strings.ReplaceAll(u, "{projectId}", strconv.Itoa(projectID))

	r := responderFromFixture("project_access_token/list.json", http.StatusOK)
	httpmock.RegisterResponder("GET", u, r)

	accessToken := "80f235b890c34ca49bcea692c2b90421"
	// PAT exists
	expected := ProjectAccessToken{
		AccessToken: accessToken,
		//"cur_rate_limit_window_count": null,
		//"cur_rate_limit_window_start": null,
		DateCreated:          1601982124,
		DateModified:         1601982124,
		Name:                 "post_client_item",
		ProjectID:            projectID,
		RateLimitWindowCount: nil,
		RateLimitWindowSize:  nil,
		Scopes: []Scope{
			ScopePostClientItem,
		},
		Status: "enabled",
	}
	actual, err := s.client.ReadProjectAccessToken(projectID, expected.AccessToken)
	s.Nil(err)
	s.Equal(expected, actual)

	// PAT does not exist
	_, err = s.client.ReadProjectAccessToken(projectID, "does-not-exist")
	s.Equal(ErrNotFound, err)

	// Not found
	r = httpmock.NewJsonResponderOrPanic(http.StatusNotFound, ErrorResult{Err: 404, Message: "Not Found"})
	httpmock.RegisterResponder("GET", u, r)
	_, err = s.client.ReadProjectAccessToken(projectID, expected.AccessToken)
	s.Equal(ErrNotFound, err)

	// Internal server error
	r = httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError,
		ErrorResult{Err: 500, Message: "Internal Server Error"})
	httpmock.RegisterResponder("GET", u, r)
	_, err = s.client.ReadProjectAccessToken(projectID, expected.AccessToken)
	s.NotNil(err)
	s.NotEqual(ErrNotFound, err)

	// Unauthorized
	r = httpmock.NewJsonResponderOrPanic(http.StatusUnauthorized,
		ErrorResult{Err: 401, Message: "Unauthorized"})
	httpmock.RegisterResponder("GET", u, r)
	_, err = s.client.ReadProjectAccessToken(projectID, expected.AccessToken)
	s.Equal(ErrUnauthorized, err)
}

// TestReadProjectAccessTokenByName tests reading a Rollbar project access token
// from the API.
func (s *Suite) TestReadProjectAccessTokenByName() {
	projectID := 411334
	u := apiUrl + pathPatList
	u = strings.ReplaceAll(u, "{projectId}", strconv.Itoa(projectID))

	r := responderFromFixture("project_access_token/list.json", http.StatusOK)
	httpmock.RegisterResponder("GET", u, r)

	// PAT with name exists
	expected := ProjectAccessToken{
		AccessToken: "80f235b890c34ca49bcea692c2b90421",
		//"cur_rate_limit_window_count": null,
		//"cur_rate_limit_window_start": null,
		DateCreated:          1601982124,
		DateModified:         1601982124,
		Name:                 "post_client_item",
		ProjectID:            projectID,
		RateLimitWindowCount: nil,
		RateLimitWindowSize:  nil,
		Scopes: []Scope{
			ScopePostClientItem,
		},
		Status: "enabled",
	}
	actual, err := s.client.ReadProjectAccessTokenByName(projectID, expected.Name)
	s.Nil(err)
	s.Equal(expected, actual)

	// PAT with name does not exist
	_, err = s.client.ReadProjectAccessTokenByName(projectID, "this-name-does-not-exist")
	s.Equal(ErrNotFound, err)

	// Project ID not found
	r = httpmock.NewJsonResponderOrPanic(http.StatusNotFound, ErrorResult{Err: 404, Message: "Not Found"})
	httpmock.RegisterResponder("GET", u, r)
	_, err = s.client.ReadProjectAccessTokenByName(projectID, "this-name-does-not-exist")
	s.Equal(ErrNotFound, err)

	// Internal server error
	r = httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError,
		ErrorResult{Err: 500, Message: "Internal Server Error"})
	httpmock.RegisterResponder("GET", u, r)
	_, err = s.client.ReadProjectAccessTokenByName(projectID, "this-name-does-not-exist")
	s.NotNil(err)
	s.NotEqual(ErrNotFound, err)

	// Unauthorized
	r = httpmock.NewJsonResponderOrPanic(http.StatusUnauthorized,
		ErrorResult{Err: 401, Message: "Unauthorized"})
	httpmock.RegisterResponder("GET", u, r)
	_, err = s.client.ReadProjectAccessTokenByName(projectID, "this-name-does-not-exist")
	s.Equal(ErrUnauthorized, err)
}

// TestDeleteProjectAccessToken tests deleting a Rollbar project access token.
func (s *Suite) TestDeleteProjectAccessToken() {
	// FIXME: actually test this
	//  https://github.com/rollbar/terraform-provider-rollbar/issues/12
	err := s.client.DeleteProjectAccessToken(1234, "does_not_matter")
	s.Nil(err)
	log.Warn().Msg("Delete project access token is not yet implemented in Rollbar API")
}

func (s *Suite) TestCreateProjectAccessToken() {
	projID := 411334

	args := ProjectAccessTokenCreateArgs{
		ProjectID: projID,
		Name:      "foobar",
		Scopes:    []Scope{ScopeRead, ScopeWrite},
		Status:    StatusEnabled,
	}
	u := apiUrl + pathPatCreate
	u = strings.ReplaceAll(u, "{projectId}", strconv.Itoa(projID))
	rs := responseFromFixture("project_access_token/create.json", http.StatusOK)
	var r httpmock.Responder
	r = func(req *http.Request) (*http.Response, error) {
		a := ProjectAccessTokenCreateArgs{}
		err := json.NewDecoder(req.Body).Decode(&a)
		log.Debug().
			Interface("args", a).
			Msg("arguments sent to API")
		s.Nil(err)
		s.Equal(args.Name, a.Name)
		s.Equal(args.Scopes, a.Scopes)
		return rs, nil
	}
	httpmock.RegisterResponder("POST", u, r)

	//
	// Sanity Checks
	//
	// Invalid project ID
	badArgs := args
	badArgs.ProjectID = 0
	_, err := s.client.CreateProjectAccessToken(badArgs)
	s.NotNil(err)
	badArgs = args
	badArgs.ProjectID = -234
	_, err = s.client.CreateProjectAccessToken(badArgs)
	s.NotNil(err)
	// Invalid project name
	badArgs = args
	badArgs.Name = ""
	_, err = s.client.CreateProjectAccessToken(badArgs)
	s.NotNil(err)
	// No scopes specified
	badArgs = args
	badArgs.Scopes = []Scope{}
	_, err = s.client.CreateProjectAccessToken(badArgs)
	s.NotNil(err)
	// Invalid scope
	badArgs = args
	derpScope := Scope("derp!")
	badArgs.Scopes = []Scope{derpScope}
	_, err = s.client.CreateProjectAccessToken(badArgs)
	s.NotNil(err)
	// Invalid status
	badArgs = args
	derpStatus := Status("derp!")
	badArgs.Status = derpStatus
	_, err = s.client.CreateProjectAccessToken(badArgs)
	s.NotNil(err)

	// Success
	t, err := s.client.CreateProjectAccessToken(args)
	s.Nil(err)
	s.NotEmpty(t.AccessToken)
	s.Equal(args.Name, t.Name)
	s.Equal(args.Scopes, t.Scopes)
	s.Equal(args.ProjectID, t.ProjectID)

	// Unauthorized
	r = httpmock.NewJsonResponderOrPanic(http.StatusUnauthorized,
		ErrorResult{Err: 401, Message: "Unauthorized"})
	httpmock.RegisterResponder("POST", u, r)
	_, err = s.client.CreateProjectAccessToken(args)
	s.Equal(ErrUnauthorized, err)

	// Unreachable server
	httpmock.Reset()
	_, err = s.client.CreateProjectAccessToken(args)
	s.NotNil(err)

	// Internal server error
	r = httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError,
		ErrorResult{Err: 500, Message: "Internal Server Error"})
	httpmock.RegisterResponder("POST", u, r)
	_, err = s.client.CreateProjectAccessToken(args)
	s.NotNil(err)
	s.NotEqual(ErrNotFound, err)
}

func (s *Suite) TestUpdateProjectAccessToken() {
	projID := 411334
	accessToken := "055ab702454e40798fd22bdac249eb2e" // Doesn't actually matter for this test

	args := ProjectAccessTokenUpdateArgs{
		ProjectID:            projID,
		AccessToken:          accessToken,
		RateLimitWindowSize:  1000,
		RateLimitWindowCount: 2500,
	}
	u := apiUrl + pathPatUpdate
	u = strings.ReplaceAll(u, "{projectId}", strconv.Itoa(projID))
	u = strings.ReplaceAll(u, "{accessToken}", accessToken)
	rs := responseFromFixture("project_access_token/update.json", http.StatusOK)
	var r httpmock.Responder
	r = func(req *http.Request) (*http.Response, error) {
		a := ProjectAccessTokenUpdateArgs{}
		err := json.NewDecoder(req.Body).Decode(&a)
		log.Debug().
			Interface("args", args).
			Msg("arguments sent to API")
		s.Nil(err)
		s.Equal(args.RateLimitWindowCount, a.RateLimitWindowCount)
		s.Equal(args.RateLimitWindowSize, a.RateLimitWindowSize)
		return rs, nil
	}
	httpmock.RegisterResponder("PATCH", u, r)

	//
	// Sanity Checks
	//
	// Invalid project ID
	badArgs := args
	badArgs.ProjectID = 0
	err := s.client.UpdateProjectAccessToken(badArgs)
	s.NotNil(err)
	badArgs = args
	badArgs.ProjectID = -234
	err = s.client.UpdateProjectAccessToken(badArgs)
	s.NotNil(err)
	// Invalid access token
	badArgs = args
	badArgs.AccessToken = ""
	err = s.client.UpdateProjectAccessToken(badArgs)
	s.NotNil(err)

	// Success
	err = s.client.UpdateProjectAccessToken(args)
	s.Nil(err)

	// Unauthorized
	r = httpmock.NewJsonResponderOrPanic(http.StatusUnauthorized,
		ErrorResult{Err: 401, Message: "Unauthorized"})
	httpmock.RegisterResponder("PATCH", u, r)
	err = s.client.UpdateProjectAccessToken(args)
	s.Equal(ErrUnauthorized, err)

	// Unreachable server
	httpmock.Reset()
	err = s.client.UpdateProjectAccessToken(args)
	s.NotNil(err)

	// Internal server error
	r = httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError,
		ErrorResult{Err: 500, Message: "Internal Server Error"})
	httpmock.RegisterResponder("PATCH", u, r)
	err = s.client.UpdateProjectAccessToken(args)
	s.NotNil(err)
	s.NotEqual(ErrNotFound, err)
}
