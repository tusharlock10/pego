package hirezapi

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tusharlock10/pego/constants"
)

// APIClient stores data needed to interface with the HiRez API and implements the HiRezAPI interface methods.
type APIClient struct {
	DeveloperID  string
	AuthKey      string
	SessionID    string
	SessionStamp string
}

// Init creates teh api instance and initializes a session.
func Init(devID, key string) (HiRezAPI, error) {
	if devID == "" {
		return nil, errors.New(`must provide a developerID (eg, 1004)`)
	}
	if key == "" {
		return nil, errors.New(`must provide an auth key (eg, 23DF3C7E9BD14D84BF892AD206B6755C)`)
	}
	api := &APIClient{
		DeveloperID: devID,
		AuthKey:     key,
	}
	err := api.CreateSession()
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (a *APIClient) MakeRequest(methodName, path string, output interface{}) error {
	signature, timestamp := a.GenerateSignature(methodName)
	apiURL := fmt.Sprintf("%s/%s%s/%s/%s/%s/%s", constants.PaladinsURL, methodName, "json", a.DeveloperID, signature, a.SessionID, timestamp)
	if path != "" {
		apiURL = fmt.Sprintf("%s/%s", apiURL, path)
	}
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}

// GenerateSignature takes in the requested methodName and generates an md5 hashed signature for sending a request
func (a *APIClient) GenerateSignature(methodName string) (string, string) {
	utcNow := time.Now().UTC().Format(constants.TimeFormat)
	sigStr := fmt.Sprintf("%s%s%s%s", a.DeveloperID, methodName, a.AuthKey, utcNow)
	bs := []byte(sigStr)
	return fmt.Sprintf("%x", md5.Sum(bs)), utcNow
}
