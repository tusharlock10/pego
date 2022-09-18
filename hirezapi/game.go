package hirezapi

import (
	"encoding/json"
	"io"

	"github.com/tusharlock10/pego/apiResponse"
)

// GetItems returns all items and their various attributes.
func (a *APIClient) GetItems() ([]apiResponse.Item, error) {
	resp, err := a.MakeRequest("getitems", "1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []apiResponse.Item
	err = json.Unmarshal(body, &output)
	return output, err
}
