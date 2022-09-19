package hirezapi

import (
	"github.com/tusharlock10/pego/apiResponse"
)

// GetItems returns all items and their various attributes.
func (a *APIClient) GetItems() ([]apiResponse.Item, error) {
	var output []apiResponse.Item
	err := a.MakeRequest("getitems", "1", &output)
	return output, err
}
