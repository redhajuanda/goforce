package force

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	QueryURL = "/services/data/v20.0/query/?"
)

type BaseResponse struct {
	Done           bool   `json:"done"`
	TotalSize      int    `json:"totalSize"`
	NextRecordsURL string `json:"nextRecordsUrl"`
}

type QueryResponse struct {
	BaseResponse
	// Records        []map[string]interface{} `json:"records"`
	Records []map[string]interface{} `json:"records"`
}

// type Contact struct {
// 	Attributes Attributes `json:"attributes"`
// 	Name       string     `json:"Name"`
// }

// type Attributes struct {
// 	Type string `json:"type"`
// 	URL  string `json:"url"`
// }

func (force *ForceAPI) Query(query string) (*QueryResponse, error) {
	req, err := force.queryRequest(query)
	if err != nil {
		return nil, err
	}
	res, err := force.queryResponse(req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (force *ForceAPI) queryRequest(query string) (*http.Request, error) {
	params := url.Values{}
	params.Set("q", query)
	url := fmt.Sprintf("%v%v%v", force.instanceURL, QueryURL, params.Encode())
	// Build Request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (force *ForceAPI) queryResponse(request *http.Request) (QueryResponse, error) {
	response, err := force.client.Do(request)

	if err != nil {
		return QueryResponse{}, err
	}

	decoder := json.NewDecoder(response.Body)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var queryErrs []ResponseError
		err = decoder.Decode(&queryErrs)
		var errMsg error
		if err == nil {
			for _, queryErr := range queryErrs {
				errMsg = fmt.Errorf("insert response err: %s: %s", queryErr.ErrorCode, queryErr.Message)
			}
		} else {
			errMsg = fmt.Errorf("insert response err: %d %s", response.StatusCode, response.Status)
		}

		return QueryResponse{}, errMsg
	}

	var resp QueryResponse
	err = decoder.Decode(&resp)
	if err != nil {
		return QueryResponse{}, err
	}

	return resp, nil
}

// func (force *ForceAPI) Insert(object string, data map[string]interface{}) ([]byte, error) {
// 	url := fmt.Sprintf("%v/services/data/v20.0/sobjects/%v", force.instanceURL, object)
// 	body, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	result, err := force.httpRequest(http.MethodPost, url, bytes.NewReader(body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("INSERT")
// 	fmt.Println(string(result))
// 	return result, nil
// }

// func (force *ForceAPI) Update(object string, id string, data map[string]interface{}) ([]byte, error) {
// 	url := fmt.Sprintf("%v/services/data/v20.0/sobjects/%v/%v", force.instanceURL, object, id)
// 	body, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	result, err := force.httpRequest(http.MethodPatch, url, bytes.NewReader(body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("UPDATE")
// 	fmt.Println(string(result))
// 	return result, nil
// }

// func (force *ForceAPI) httpRequest(method, url string, body io.Reader) ([]byte, error) {
// 	// Build Request
// 	req, err := http.NewRequest(method, url, body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	// Do Request
// 	resp, err := force.client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp_body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	return resp_body, nil
// }
