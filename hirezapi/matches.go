package hirezapi

import (
	"fmt"
	"io/ioutil"
)

// GetMatchDetails returns the statistics for a particular completed match.
func (a *APIClient) GetMatchDetails() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}

// GetMatchDetailsBatch returns the statistics for a particular set of completed matches. [limit batch query to 5-10 matchIDs]
func (a *APIClient) GetMatchDetailsBatch() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}

// GetMatchPlayerDetails returns player information for a live match.
func (a *APIClient) GetMatchPlayerDetails() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}

// GetMatchIDsByQueue list all MatchIDs for a particular match queue
func (a *APIClient) GetMatchIDsByQueue() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}

// GetQueueStats returns match summary stats for a player + queue, grouped by Gods played.
func (a *APIClient) GetQueueStats() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}
