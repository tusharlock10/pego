package hirezapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tusharlock10/pego/apiResponse"
	"github.com/tusharlock10/pego/constants"
)

// CreateSession is a required step to Authenticate the developerId/signature for further API use.
func (a *APIClient) CreateSession() error {
	sig, stamp := a.GenerateSignature("createsession")
	path := fmt.Sprintf(
		"%s/%s%s/%s/%s/%s",
		constants.PaladinsURL,
		"createsession",
		"json",
		a.DeveloperID,
		sig,
		stamp,
	)
	resp, err := http.Get(path)
	if err != nil {
		return fmt.Errorf("error creating session: %v", err)
	}
	defer resp.Body.Close()
	session := &apiResponse.Session{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}
	err = json.Unmarshal(body, session)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}
	if session.RetMsg != "Approved" {
		return fmt.Errorf("error creating session: %v", session.RetMsg)
	}
	a.SessionID = session.SessionID
	a.SessionStamp = session.Timestamp
	return nil
}

// GetHiRezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
func (a *APIClient) GetHiRezServerStatus() ([]apiResponse.HiRezServerStatus, error) {
	var output []apiResponse.HiRezServerStatus
	err := a.MakeRequest("gethirezserverstatus", "", &output)
	return output, err
}

// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
func (a *APIClient) GetDataUsed() ([]apiResponse.DataUsed, error) {
	var output []apiResponse.DataUsed
	err := a.MakeRequest("getdataused", "", &output)
	return output, err
}
