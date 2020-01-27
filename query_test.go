package goforce

import "testing"

import "fmt"

func TestQuery(t *testing.T) {
	force := NewClientTest()
	res, err := force.Query("SELECT Name FROM Contact LIMIT 10")
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Println(res)
}

// func TestUpdate(t *testing.T) {
// 	force := NewClientTest()
// 	data := map[string]interface{}{"FirstName": "Fresh"}
// 	_, err := force.Update("Contact", "0030w000006ujRdAAI", data)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}
// }

// func TestInsert(t *testing.T) {
// 	force := NewClientTest()
// 	data := map[string]interface{}{
// 		"LastName":        "Test",
// 		"Country_g__c":    "ID",
// 		"CurrencyIsoCode": "IDR",
// 		"AccountId":       "0010w000008c2saAAA",
// 	}
// 	_, err := force.Insert("Contact", data)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}
// }
