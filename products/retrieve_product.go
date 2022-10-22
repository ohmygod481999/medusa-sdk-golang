package products

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type RetrieveProductData struct {
	Product []*schema.Product `json:"product"`
}

type ListCollectionResponse struct {
	// Success response
	Data *RetrieveProductData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

func Retrieve(id string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/products/%v", id)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}