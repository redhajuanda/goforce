package goforce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Requester interface {
	URL() string
	RequestBody() map[string]interface{}
}

type Response struct {
	Status    string
	ErrorCode string
	Message   string
	Data      interface{}
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
	if method == "GET" {
		req, err := force.createGETRequest(requester)
		if err != nil {
			return nil, err
		}
		resp, err := force.createResponse(req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else if method == "POST" {
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

func (force *ForceAPI) createGETRequest(requester Requester) (*http.Request, error) {

	v := url.Values{}
	var q string
	for key, val := range requester.RequestBody() {
		fmt.Println(key, val)
		v.Add(key, fmt.Sprintf("%v", val))
		q = v.Encode()
	}

	url := force.instanceURL + requester.URL() + "?" + q
	fmt.Println(url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("createGETRequest, Error on request: %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	return request, nil
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
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error when reading body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		var respErr []ResponseError
		err = json.Unmarshal(data, &respErr)
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

	var resp Response
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, fmt.Errorf("Error when unmarshalling json response: %v", err)
	}

	return &resp, nil
}
