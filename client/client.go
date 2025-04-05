package client

import (
	"fmt"
	"os"

	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type Client struct {
	OAuth *OAuthData
}

type Parameters struct {
	Value map[string][]string
	File  map[string][]*os.File
}

type OAuthData struct {
	AuthToken    string
	RefreshToken string
}

type MethodParametr struct {
	Method   string
	Parametr string
}

func (c Client) Request(root bool, url, function string, parameters Parameters) (result map[string]interface{}, err error) {
	responseByte, err := c.callMethod(root, url, function, parameters)
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
func (c Client) callMethod(root bool, url, function string, params Parameters) ([]byte, error) {
	u := url + function + ".json"
	body, writer, err := createFormFields(params)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST", u, bytes.NewReader(body),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if root == true {
		request.Header.Set("Content-Type", writer.FormDataContentType())
		param := map[string]string{
			"auth": c.OAuth.AuthToken,
		}

		q := request.URL.Query()
		for p, v := range param {
			q.Set(p, v)
		}
		request.URL.RawQuery = q.Encode()

	} else {
		request.Header.Set("Content-Type", writer.FormDataContentType())
	}

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
