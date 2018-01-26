package ionic

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ion-channel/ionic/users"
)

const (
	sessionsLoginEndpoint = "v1/sessions/login"
)

// Session represents the BearerToken and User for the current session
type Session struct {
	BearerToken string     `json:"jwt"`
	User        users.User `json:"user"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login performs basic auth requests with a username and returns a Login
// response with bearer token and user for the session.  Returns an error for
// HTTP and JSON errors.
func (ic *IonClient) Login(username, password string) (*Session, error) {
	auth := fmt.Sprintf("%v:%v", username, password)
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(auth))))
	headers.Add("Content-Type", "application/json; charset=UTF-8")

	login := loginRequest{username, password}
	b, err := json.Marshal(login)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal login body to JSON: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(sessionsLoginEndpoint, "", nil, *buff, headers)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err.Error())
	}

	var resp Session
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal login response: %v", err.Error())
	}

	return &resp, nil
}
