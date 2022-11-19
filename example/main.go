package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/calexa22/randall"
	"github.com/joho/godotenv"
)

func main() {
	headerValues := GetHeaderValuesFromEnv()

	fmt.Println("AccessToken: :)")
	fmt.Printf("AccountId: %s\n", headerValues.AccountId)
	fmt.Printf("UserAgentApp: %s\n", headerValues.UserAgentApp)
	fmt.Printf("UserAgentEmail: %s\n", headerValues.UserAgentEmail)
	fmt.Println()

	client := randall.New(headerValues)

	resp, err := client.Users.Me()

	if err != nil {
		log.Panic(err)
	}

	bytes, err := json.MarshalIndent(resp.Data, "", "\t")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("/v2/users/me Response:")
	fmt.Println()
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println()
	fmt.Println(string(bytes))
}

func GetHeaderValuesFromEnv() randall.HarvestHeaders {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	var hv randall.HarvestHeaders

	if token, exists := os.LookupEnv("HARVEST_ACCESS_TOKEN"); !exists || token == "" {
		log.Fatal("Unable to retrieve HARVEST_ACCESS_TOKEN from .env file")
	} else {
		hv.AccessToken = token
	}

	if id, exists := os.LookupEnv("HARVEST_ACCOUNT_ID"); !exists || id == "" {
		log.Fatal("Unable to retrieve HARVEST_ACCOUNT_ID from .env file")
	} else {
		hv.AccountId = id
	}

	if app, exists := os.LookupEnv("USER_AGENT_APP"); !exists || app == "" {
		log.Fatal("Unable to retrieve USER_AGENT_APP from .env file")
	} else {
		hv.UserAgentApp = app
	}

	if email, exists := os.LookupEnv("USER_AGENT_EMAIL"); !exists || email == "" {
		log.Fatal("Unable to retrieve USER_AGENT_EMAIL from .env file")
	} else {
		hv.UserAgentEmail = email
	}

	return hv
}
