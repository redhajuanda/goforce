# goforce
A library to interface with Salesforce APIs

## Install 

```go get github.com/redhajuanda/goforce```

## Make a connection

```
const (
  testClientId      = "xxx"
  testClientSecret  = "xxx"
  testUserName      = "xxx"
  testPassword      = "xxx"
  testSecurityToken = "xxx"
  testEnvironment   = "sandbox"
)
forceAPI, err := goforce.NewClient(testClientId, testClientSecret, testUserName, testPassword, testSecurityToken, testEnvironment)
if err != nil {
  log.Fatalf("Failed connecting to salesforce: %v", err)
}
```

## Query SOQL

```
result, err := forceAPI.Query("SELECT Id FROM Contact LIMIT 10")
if err != nil {
  log.Fatalf("Query failed: %v", err)
}
// do something with response
```

## Insert Record 
```
insert := &goforce.InsertData{
		SObject: "Contact",
		Fields:  map[string]interface{
        		"LastName": "Redha",
    		},
}

resp, err := forceAPI.Insert(insert)
if err != nil {
  log.Fatalf("Insert failed: %v", err)
}
// do something with response
```

## Update Record 
```
update := &goforce.UpdateData{
		SObject: "Contact",
    		ID: "xxcontact_idxx",
		Fields:  map[string]interface{
        		"LastName": "Redha update",
    		},
}

resp, err := forceAPI.Update(update)
if err != nil {
  log.Fatalf("Update failed: %v", err)
}
// do something with response
```

## Delete Record 
```
delete := &goforce.UpdateData{
		SObject: "Contact",
    		ID: "xxcontact_idxx",
}

resp, err := forceAPI.Delete(delete)
if err != nil {
  log.Fatalf("Delete failed: %v", err)
}
// do something with response
```

## Request API
```
data_request := &goforce.Request{
		ServiceURL: "/services/apexrest/amalia/v2/contact",
		Fields:     map[string]interface{
        		"Name": "Redha",
    		},
}
resp, err := forceAPI.Request("POST", data_request)
if err != nil {
  log.Fatalf("Request failed: %v", err)
}
// do something with response
```
