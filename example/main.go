package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/calexa22/randall"
)

func main() {

	accountId, err := GetEnvVariable("HARVEST_ACCOUNT_ID")

	if err != nil {
		log.Fatal(err)
	}

	token, err := GetEnvVariable("HARVEST_ACCESS_TOKEN")

	if err != nil {
		log.Fatal(err)
	}

	app, err := GetEnvVariable("USER_AGENT_APP")

	if err != nil {
		log.Fatal(err)
	}

	email, err := GetEnvVariable("USER_AGENT_EMAIL")

	if err != nil {
		log.Fatal(err)
	}

	randall.NewClient(
		accountId,
		token,
		app,
		email,
	)

	// client.Roles.CreateRole()

	//fmt.Println(reflect.TypeOf(randall.HarvestDate{}))

	// 1941661671

	// resp, err := client.Users.Me()

	// if err != nil {
	// 	log.Panic(err)
	// }

	// PrintResponse(resp, "/v2/users/me")

	// resp, err = client.Company.MyCompany()

	// if err != nil {
	// 	log.Panic(err)
	// }

	// PrintResponse(resp, "/v2/company")

	// from := time.Date(2022, time.November, 24, 0, 0, 0, 0, time.UTC)

	// resp, err := client.TimeEntries.All(randall.GetTimeEntriesParams{
	// 	FromDate: &from,
	// })

	// per := 1
	// resp, err := client.TimeEntries.All(randall.GetTimeEntriesParams{
	// 	PerPage: &per,
	// })

	// if err != nil {
	// 	log.Panic(err)
	// }

	// PrintResponse(resp, "/v2/time_entries")

	// resp, err := client.TimeEntries.CreateViaDuration(randall.TimeEntryViaDurationRequest{
	// 	TaskId:    8627145,
	// 	ProjectId: 26374788,
	// 	SpentDate: randall.HarvestDate(time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC)),
	// 	Hours:     8,
	// })

	// if err != nil {
	// 	log.Panic(err)
	// }

	// PrintResponse(resp, "POST /v2/time_entries")

	// resp, err := client.TimeEntries.DeleteTimeEntry(1941676550)

	// if err != nil {
	// 	log.Panic(err)
	// }

	// PrintResponse(resp, "DELETE /v2/time_entries/{id}")
}

func GetEnvVariable(key string) (string, error) {
	v, exists := os.LookupEnv(key)

	if !exists || v == "" {
		return "", fmt.Errorf("Unable to retrieve value for %s from .env file\n", key)
	}

	return v, nil
}

func PrintResponse(resp randall.HarvestResponse, endpoint string) {
	bytes, _ := json.MarshalIndent(resp.Data, "", "\t")

	fmt.Println()
	fmt.Printf("%s Response:\n", endpoint)
	fmt.Println()
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println()
	fmt.Println(string(bytes))
}
