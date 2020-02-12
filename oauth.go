package goforce

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type oauth struct {
	ClientId      string
	ClientSecret  string
	Username      string
	Password      string
	SecurityToken string
	Environment   string
}

// Method for authenticate client
func (o *oauth) Authenticate() (client *http.Client, instanceURL string, err error) {
	loginURL := "https://test.salesforce.com/services/oauth2/token"
	if o.Environment == "production" {
		loginURL = "https://login.salesforce.com/services/oauth2/token"
	}
	// Using Google Oauth2 to authenticate
	conf := &oauth2.Config{
		ClientID:     o.ClientId,
		ClientSecret: o.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: loginURL,
		},
	}
	// Get the token
	token, err := conf.PasswordCredentialsToken(oauth2.NoContext, o.Username, o.Password)
	// fmt.Printf("token = %+v\n", token)
	if err != nil {
		return nil, "", fmt.Errorf("Authentication Error %+v", err)
	}
	// Get the instace URL
	instanceURL, _ = token.Extra("instance_url").(string)
	// Build HTTP Client
	client = conf.Client(oauth2.NoContext, token)
	return client, instanceURL, nil
}
