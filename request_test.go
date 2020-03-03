package goforce

import "testing"

import "fmt"

func TestRequestPOST(t *testing.T) {
	force := NewClientTest()
	params := Request{
		ServiceURL: "/services/apexrest/amalia/v1/contact",
		Fields: map[string]interface{}{
			"FirstName":           "Redha",
			"MiddleName":          "",
			"LastName":            "Capung",
			"Email":               "email@gma.com",
			"MobilePhone":         "08",
			"Territory":           "ID",
			"Curr":                "IDR",
			"Rererral_Program_Id": "",
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
			"OpportunityId":       "0060w0000045V5vAAE",
			"Outstanding_Amount":  998877,
			"Monthly_Installment": 112233,
			"Account_Number":      "123456789",
		},
	}

	response, err := force.Request("PUT", &params)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println(response.Message)
}
