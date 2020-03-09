package goforce

import "testing"

import "fmt"

func TestRequestGET(t *testing.T) {
	force := NewClientTest()
	params := Request{
		ServiceURL: "/services/apexrest/amalia/v2/opportunity",
		Fields: map[string]interface{}{
			"ContactId": "0030w000008aoTEAAY",
		},
	}

	response, err := force.Request("GET", &params)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println(response.Message)
	fmt.Println("=============================")
	hehe := response.Data.([]interface{})
	fmt.Println(hehe[0])
}

func TestRequestPOST(t *testing.T) {
	force := NewClientTest()
	params := Request{
		ServiceURL: "/services/apexrest/amalia/v1/contact",
		Fields: map[string]interface{}{
			"FirstName":           "Redha",
			"LastName":            "Capung",
			"Email":               "email@gma.com",
			"MobilePhone":         "08",
			"Country":             "ID",
			"Referrer_Program_Id": "",
			"Referrer_Id":         "",
			"Referrer_Email":      "",
			"Whatsapp_Opt_In":     "Yes",
		},
	}

	response, err := force.Request("POST", &params)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println(response.Message)
}

func TestRequestPUT(t *testing.T) {
	force := NewClientTest()
	params := Request{
		ServiceURL: "/services/apexrest/amalia/v2/opportunity",
		Fields: map[string]interface{}{
			"OpportunityId":       "0060w0000043jD9AAI",
			"Outstanding_Amount":  998877,
			"Monthly_Installment": 112233,
			"Account_Number":      "123456789",
			"Last_Payment":        "10/10/2019",
		},
	}

	response, err := force.Request("PUT", &params)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println(response.Message)
	fmt.Printf("%+v", response.Data)
}
