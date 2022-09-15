package hirezapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tusharlock10/pego/constants"
	"github.com/tusharlock10/pego/models"
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
	sess := &models.Session{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}
	err = json.Unmarshal(body, sess)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}
	if sess.RetMsg != "Approved" {
		return fmt.Errorf("error creating session: %v", sess.RetMsg)
	}
	a.SessionID = sess.SessionID
	a.SessionStamp = sess.Timestamp
	return nil
}

// GetHiRezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
func (a *APIClient) GetHiRezServerStatus() ([]models.HiRezServerStatus, error) {
	resp, err := a.MakeRequest("gethirezserverstatus", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.HiRezServerStatus
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
func (a *APIClient) GetDataUsed() ([]models.DataUsed, error) {
	resp, err := a.MakeRequest("getdataused", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.DataUsed
	err = json.Unmarshal(body, &output)
	return output, err
}
