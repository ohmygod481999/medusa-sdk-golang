package article_categories

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

type RetrieveArticleCategoryData struct {
	ArticleCategory *schema.ArticleCategory `json:"article_category"`
}

type RetrieveArticleCategoryResponse struct {
	// Success response
	Data *RetrieveArticleCategoryData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

// Retrieves a Product.
func Retrieve(id string, config *medusa.Config) (*RetrieveArticleCategoryResponse, error) {
	path := fmt.Sprintf("/store/article-category/%v", id)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(RetrieveArticleCategoryResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(RetrieveArticleCategoryData)
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
