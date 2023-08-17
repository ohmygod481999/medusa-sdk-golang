package articles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/ohmygod481999/medusa-sdk-golang"
	"github.com/ohmygod481999/medusa-sdk-golang/common"
	"github.com/ohmygod481999/medusa-sdk-golang/request"
	"github.com/ohmygod481999/medusa-sdk-golang/response"
	"github.com/ohmygod481999/medusa-sdk-golang/schema"
	"github.com/ohmygod481999/medusa-sdk-golang/utils"
)

type ListArticleData struct {
	// Array of articles
	Articles []*schema.Article `json:"articles"`

	// The total number of items available
	Count uint `json:"count"`

	// The number of items skipped before these items
	Offset uint `json:"offset"`

	// The number of items per page
	Limit uint `json:"limit"`
}

type ListArticleResponse struct {
	// Success response
	Data *ListArticleData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type ListArticle struct {
	// Query used for searching products by title, description, variant's title, variant's sku, and collection's title
	Q string `json:"q,omitempty" url:"q,omitempty"`

	// product IDs to search for.
	Ids []string `json:"id,omitempty" url:"id,omitempty"`

	// Tag IDs to search for
	Tags []string `json:"tags,omitempty" url:"tags,omitempty"`

	// title to search for.
	Title string `json:"title,omitempty" url:"title,omitempty"`

	// content to search for
	Content string `json:"content,omitempty" url:"content,omitempty"`

	// handle to search for.
	Handle string `json:"handle,omitempty" url:"handle,omitempty"`

	// Date comparison for when resulting products were created.
	CreatedAt *common.DateComparison `json:"created_at,omitempty" url:"created_at,omitempty"`

	// Date comparison for when resulting products were updated.
	UpdatedAt *common.DateComparison `json:"updated_at,omitempty" url:"updated_at,omitempty"`

	// How many products to skip in the result.
	Offset int `json:"offset" url:"offset"`

	// Limit the number of products returned.
	Limit int `json:"limit" url:"limit"`

	// (Comma separated) Which fields should be expanded in each order of the result.)
	Expand string `json:"expand,omitempty" url:"expand,omitempty"`

	// (Comma separated) Which fields should be included in each order of the result.
	Fields string `json:"fields,omitempty" url:"fields,omitempty"`
}

func NewListArticle() *ListArticle {
	p := new(ListArticle)
	p.Offset = 0
	p.Limit = 100
	return p
}

func (p *ListArticle) SetQ(q string) *ListArticle {
	p.Q = q
	return p
}

func (p *ListArticle) SetIds(ids []string) *ListArticle {
	p.Ids = ids
	return p
}

func (p *ListArticle) SetTags(tags []string) *ListArticle {
	p.Tags = tags
	return p
}

func (p *ListArticle) SetTitle(title string) *ListArticle {
	p.Title = title
	return p
}

func (p *ListArticle) SetContent(content string) *ListArticle {
	p.Content = content
	return p
}

func (p *ListArticle) SetHandle(handle string) *ListArticle {
	p.Handle = handle
	return p
}

func (p *ListArticle) SetCreatedAt(creatdAt *common.DateComparison) *ListArticle {
	p.CreatedAt = creatdAt
	return p
}

func (p *ListArticle) SetUpdatedAt(updatedAt *common.DateComparison) *ListArticle {
	p.UpdatedAt = updatedAt
	return p
}

func (p *ListArticle) SetOffset(offset int) *ListArticle {
	p.Offset = offset
	return p
}

func (p *ListArticle) SetLimit(limit int) *ListArticle {
	p.Limit = limit
	return p
}

func (p *ListArticle) SetExpand(expand string) *ListArticle {
	p.Expand = expand
	return p
}

func (p *ListArticle) SetFields(fields string) *ListArticle {
	p.Fields = fields
	return p
}

// Retrieve a list of Articles.
func (c *ListArticle) List(config *medusa.Config) (*ListArticleResponse, error) {
	path := "/store/article"

	qs, err := query.Values(c)
	if err != nil {
		return nil, err
	}

	parseStr := qs.Encode()

	path = fmt.Sprintf("%v?%v", path, parseStr)

	resp, err := request.
		NewRequest().
		SetMethod(http.MethodGet).
		SetPath(path).
		Send(config)

	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(ListArticleResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ListArticleData)
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
