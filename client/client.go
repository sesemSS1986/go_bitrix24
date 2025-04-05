package client

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"

	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

const (
	BatchLimit     = 45
	ResponseOffset = 50
)

type Client struct {
	AuthClient *resty.Client
	oAuth      *OAuthData
	Url        string
	WebHookUrl string
}

type Parameters struct {
	Value map[string][]string
	File  map[string][]*os.File
}

type FAuthData struct {
	ClientId string `valid:"required"`
	Secret   string `valid:"alphanum,required"`
}

type OAuthData struct {
	AuthToken    string `valid:"alphanum,required"`
	RefreshToken string `valid:"alphanum,required"`
}

type WebhookAuthData struct {
	UserID string `valid:"required"`
	Secret string `valid:"alphanum,required"`
}

type MethodParametr struct {
	Method   string
	Parametr string
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewClientWithOAuth(url, authToken, refreshToken string) (*Client, error) {

	auth := &OAuthData{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}

	_, err := govalidator.ValidateStruct(auth)
	if err != nil {
		return nil, errors.Wrap(err, "Auth params validation failed")
	}

	return &Client{
		AuthClient: resty.DefaultClient,
		Url:        url,
		oAuth:      auth,
	}, nil
}

func NewClientWithWebhookInCome(WebHookUrl string) (*Client, error) {

	return &Client{
		AuthClient: resty.DefaultClient,
		Url:        WebHookUrl,
	}, nil
}

func NewEnvClientWithOauth() (*Client, error) {
	return NewClientWithOAuth(
		os.Getenv("BITRIX_URL"),
		os.Getenv("BITRIX_AUTH_TOKEN"),
		os.Getenv("BITRIX_REFRESH_TOKEN"))
}

func (c *Client) SetInsecureSSL(v bool) {
	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: v})
}

func (c *Client) SetDebug(v bool) {
	resty.SetDebug(v)
}

// You need to use the method only if the endpoint you want to call is not implemented in this package.
// If possible, it is necessary to use more specific methods defined in this package to work with Bitrix24 entities.
// */
func (c Client) Request(function string, parameters Parameters) (result map[string]interface{}, err error) {
	responseByte, err := c.callMethod(function, parameters)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return
}

// The method that actually executes the request to Bitrix24 and passes the response to the calling method.
func (c Client) callMethod(function string, params Parameters) ([]byte, error) {
	url := c.WebHookUrl + function + ".json"
	body, writer, err := createFormFields(params)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST", url, bytes.NewReader(body),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка")
		return nil, err
	}

	return bytes, nil
}

// Creates form fields from the "Parameters" type in the "multipart/form-data" format.
func createFormFields(p Parameters) ([]byte, *multipart.Writer, error) {
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	for key, values := range p.Value {
		for _, value := range values {
			fw, err := writer.CreateFormField(key)
			if err != nil {
				return nil, nil, err
			}
			_, err = io.Copy(fw, strings.NewReader(value))
			if err != nil {
				return nil, nil, err
			}
		}

	}

	for key, files := range p.File {
		for _, file := range files {
			fw, err := writer.CreateFormFile(key, file.Name())
			if err != nil {
				return nil, nil, err
			}
			_, err = io.Copy(fw, file)
			if err != nil {
				return nil, nil, err
			}
		}

	}

	writer.Close()
	return buffer.Bytes(), writer, nil
}
