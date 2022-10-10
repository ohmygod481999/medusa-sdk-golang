package carts

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type CartContext map[string]any

type CartItem struct {
	VariantId string `json:"variant_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}
type CreateCart struct {
	RegionId       string      `json:"region_id,omitempty"`
	SalesChannelId string      `json:"sales_channel_id,omitempty"`
	CountryCode    string      `json:"country_code,omitempty"`
	Items          []*CartItem `json:"items,omitempty"`
	Context        CartContext `json:"context,omitempty"`
}

func NewCreateCart() *CreateCart {
	return new(CreateCart)
}

func (c *CreateCart) SetRegionId(regionId string) *CreateCart {
	c.RegionId = regionId
	return c
}
func (c *CreateCart) SetSalesChannelId(salesChannelId string) *CreateCart {
	c.SalesChannelId = salesChannelId
	return c
}
func (c *CreateCart) SetCountryCode(countryCode string) *CreateCart {
	c.CountryCode = countryCode
	return c
}
func (c *CreateCart) SetItems(items []*CartItem) *CreateCart {
	c.Items = items
	return c
}
func (c *CreateCart) SetContext(context CartContext) *CreateCart {
	c.Context = context
	return c
}

func (c *CreateCart) Create(config *medusa.Config) ([]byte, error) {
	const path = `/store/carts`
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetData(c).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
