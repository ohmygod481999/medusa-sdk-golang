package customers

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ResetPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func NewResetPassword() *ResetPassword {
	return new(ResetPassword)
}

func (r *ResetPassword) SetEmail(email string) *ResetPassword {
	r.Email = email
	return r
}

func (r *ResetPassword) SetPassword(password string) *ResetPassword {
	r.Password = password
	return r
}

func (r *ResetPassword) SetToken(token string) *ResetPassword {
	r.Token = token
	return r
}

func (r *ResetPassword) Reset(config *medusa.Config) (*Customer, error) {
	path := "/store/customers/password-reset"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(r).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(ResponseBody)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}

	return respBody.Customer, nil
}