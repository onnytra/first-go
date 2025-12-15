package external

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/onnytra/first-go/internal/utils"
	"github.com/pkg/errors"
)

type UserData struct {
	ID    int
	Email string
	Role  string
}

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}

func ValidateToken(ctx context.Context, accessToken string) (*UserData, error) {
	request := ValidateTokenRequest{AccessToken: accessToken}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request")
	}

	host := utils.GetEnv("AUTH_HOST", "")
	baseURL := fmt.Sprintf("%s/api/v2/auth/validate-token", host)
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}

	httpReq, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new request")
	}

	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed connect to validate token, status code: %d", resp.StatusCode)
	}

	resBody := UserData{}
	err = json.NewDecoder(resp.Body).Decode(&resBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return &resBody, nil
}
