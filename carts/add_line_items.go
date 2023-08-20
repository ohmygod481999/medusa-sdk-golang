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

type AddLineItemsData struct {
	Cart *schema.Cart `json:"cart"`
}

type AddLineItemsResponse struct {
	// Success response
	Data *AddLineItemsData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type LineItem struct {
	VariantId string `json:"variant_id"`
	Quantity  int    `json:"quantity"`
}

func NewLineItem() *LineItem {
	return new(LineItem)
}

func (s *LineItem) SetVariantId(variantId string) *LineItem {
	s.VariantId = variantId
	return s
}

func (s *LineItem) SetQuantity(quantity int) *LineItem {
	s.Quantity = quantity
	return s
}

// Adds a Shipping Method to the Cart.
func (s *LineItem) Add(cartId string, config *medusa.Config) (*AddLineItemsResponse, error) {
	path := fmt.Sprintf("/store/carts/%v/line-items", cartId)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(s).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(AddLineItemsResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(AddLineItemsData)
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
