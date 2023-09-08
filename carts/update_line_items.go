package carts

import (
	"encoding/json"
	"fmt"
	"net/http"

	medusa "github.com/ohmygod481999/medusa-sdk-golang"
	"github.com/ohmygod481999/medusa-sdk-golang/request"
	"github.com/ohmygod481999/medusa-sdk-golang/response"
	"github.com/ohmygod481999/medusa-sdk-golang/schema"
	"github.com/ohmygod481999/medusa-sdk-golang/utils"
)

type UpdateLineItemsBody struct {
	Quantity int `json:"quantity"`
}

type UpdateLineItemsData struct {
	Cart *schema.Cart `json:"cart"`
}

type UpdateLineItemsResponse struct {
	// Success response
	Data *UpdateLineItemsData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

// Adds a Shipping Method to the Cart.
func UpdateLineItem(cartId string, lineId string, quantity int, config *medusa.Config) (*UpdateLineItemsResponse, error) {
	path := fmt.Sprintf("/store/carts/%v/line-items/%s", cartId, lineId)
	reqBody := UpdateLineItemsBody{
		Quantity: quantity,
	}
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(reqBody).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(UpdateLineItemsResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(UpdateLineItemsData)
		if err := json.Unmarshal(body, respData); err != nil {
			return nil, err
		}
		respBody.Data = respData

	case http.StatusUnauthorized:
		respErr := utils.UnauthorizeError()
		respBody.Error = respErr

	case http.StatusBadRequest:
		respErrors, err := utils.ParseErrors(body)
		if err != nil {
			return nil, err
		}
		if len(respErrors.Errors) == 0 {
			respError, err := utils.ParseError(body)
			if err != nil {
				return nil, err
			}
			respBody.Error = respError
		} else {
			respBody.Errors = respErrors
		}

	default:
		respErr, err := utils.ParseError(body)
		if err != nil {
			return nil, err
		}
		respBody.Error = respErr
	}

	return respBody, nil
}
