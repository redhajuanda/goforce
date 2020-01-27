package goforce

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

const (
	RequestURL = "/services/data"
	Version    = "v40.0"
)

type ForceAPI struct {
	client      *http.Client
	instanceURL string
}

func (force *ForceAPI) getServiceURL() string {
	return fmt.Sprintf("%v%v/%v", force.instanceURL, RequestURL, Version)
}

func NewClient(clientId string, clientSecret string, userName string, password string, securityToken string, environment string) (*ForceAPI, error) {
	oauth := oauth{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Username:     userName,
		Password:     password + securityToken,
		Environment:  environment,
	}
	httpClient, instanceURL, err := oauth.Authenticate()
	if err != nil {
		return nil, fmt.Errorf("Authentication Error %+v", err)
	}
	forceAPI := &ForceAPI{
		client:      httpClient,
		instanceURL: instanceURL,
	}
	return forceAPI, nil
}

func NewClientTest() *ForceAPI {
	loadEnv()
	var (
		testClientId      = os.Getenv("testClientId")
		testClientSecret  = os.Getenv("testClientSecret")
		testUserName      = os.Getenv("testUserName")
		testPassword      = os.Getenv("testPassword")
		testSecurityToken = os.Getenv("testSecurityToken")
		testEnvironment   = os.Getenv("testEnvironment")
	)
	client, err := NewClient(testClientId, testClientSecret, testUserName, testPassword, testSecurityToken, testEnvironment)

	if err != nil {
		log.Panicf("Cannot Create Test Client\n %+v", err)
	}
	return client
}

func loadEnv(){
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}
