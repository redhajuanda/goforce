package force

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
