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
