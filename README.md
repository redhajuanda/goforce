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
  fmt.Println(err)
}
```

## Query SOQL

```
result, err := forceAPI.Query("SELECT Id FROM Contact LIMIT 10")
if err != nil {
  fmt.Println(err)
}
fmt.Println(result)
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
  fmt.Println(err)
}
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
  fmt.Println(err)
}
```

## Delete Record 
```
delete := &goforce.UpdateData{
		SObject: "Contact",
    ID: "xxcontact_idxx",
}

resp, err := forceAPI.Delete(delete)
if err != nil {
  fmt.Println(err)
}
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
  fmt.Println(err)
}
```
