package goforce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Requester interface {
	URL() string
	RequestBody() map[string]interface{}
}

type Response struct {
	Status    string
	ErrorCode string
	Message   string
	Data      []map[string]interface{}
}

type ResponseError struct {
	ErrorCode string
	Message   string
}

// Represents a request
type Request struct {
	ServiceURL string
	Fields     map[string]interface{}
}

// Get request body from Request struct
func (r *Request) RequestBody() map[string]interface{} {
	return r.Fields
}

func (r *Request) URL() string {
	return r.ServiceURL
}

func (force *ForceAPI) Request(method string, requester Requester) (*Response, error) {
	if method == "POST" {
		req, err := force.createPostRequest(requester)
		if err != nil {
			return nil, err
		}
		resp, err := force.createResponse(req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else if method == "PUT" {
		req, err := force.createPutRequest(requester)
		if err != nil {
			return nil, err
		}
		resp, err := force.createResponse(req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else {
		return nil, fmt.Errorf("Method is not supported")
	}

}

func (force *ForceAPI) createPostRequest(requester Requester) (*http.Request, error) {

	url := force.instanceURL + requester.URL()
	body, err := json.Marshal(requester.RequestBody())
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	// r.dml.session.AuthorizationHeader(request)
	return request, nil
}

func (force *ForceAPI) createPutRequest(requester Requester) (*http.Request, error) {

	url := force.instanceURL + requester.URL()
	body, err := json.Marshal(requester.RequestBody())
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	// r.dml.session.AuthorizationHeader(request)
	return request, nil
}

func (force *ForceAPI) createResponse(request *http.Request) (*Response, error) {

	response, err := force.client.Do(request)
	decoder := json.NewDecoder(response.Body)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var respErr []ResponseError
		err = decoder.Decode(&respErr)
		var errMsg error
		if err == nil {
			for _, respErr := range respErr {
				errMsg = fmt.Errorf("request response err: %s: %s", respErr.ErrorCode, respErr.Message)
			}
		} else {
			errMsg = fmt.Errorf("request response err: %d %s", response.StatusCode, response.Status)
		}
		return nil, errMsg
	}

	var resp []Response
	err = decoder.Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp[0], nil
}
