package smooch

import (
	"net/http"

	"fmt"

	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/imroc/req"
)

type Smooch struct {
	Req     *req.Req
	Headers req.Header
	domain  string
}

func NewSmooch() *Smooch {
	return &Smooch{
		Req:     req.New(),
		Headers: req.Header{},
		domain:  "https://api.smooch.io/v1/",
	}
}

func (s *Smooch) Auth(keyID, secret string) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"scope": "app",
	})

	token.Header["kid"] = keyID

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	s.Headers["authorization"] = "Bearer " + tokenString

	return nil
}

type ErrorResponse struct {
	Error struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"error"`
}

func (s *Smooch) request(method string) (*req.Resp, error) {
	resp, err := s.Req.Get(s.domain+method, s.Headers)
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != http.StatusOK {
		errorResponse := &ErrorResponse{}

		err = resp.ToJSON(errorResponse)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Code: %v, description: %v", errorResponse.Error.Code, errorResponse.Error.Description)
	}

	return resp, nil
}

func (s *Smooch) requestPost(method string, body interface{}) (*req.Resp, error) {
	resp, err := s.Req.Post(s.domain+method, req.BodyJSON(body), s.Headers)
	if err != nil {
		return nil, err
	}

	log.Println("Code:", resp.Response().StatusCode)

	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusCreated {
		errorResponse := &ErrorResponse{}

		err = resp.ToJSON(errorResponse)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Code: %v, description: %v", errorResponse.Error.Code, errorResponse.Error.Description)
	}

	return resp, nil
}
