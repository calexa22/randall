
# randall
A [Harvest Time Tracking API](https://help.getharvest.com/api-v2/) client, written in Go.

![randall logo](https://raw.githubusercontent.com/calexa22/gorandall/main/.github/randall.gif)

## Features

 * A Go client interface with support for (almost) all Harvest V2 API REST endpoints (`/v2/reports/*` endpoints may be implemented at a later date)
 * Quick auth setup via your Harvest API Personal Access Token (visit https://help.getharvest.com/api-v2/authentication-api/authentication/authentication/ for more info)
 * Pagination support for GET collection endpoints
 * JSON serialization/deserialization
 * `multipart/form-data` request support for endpoints that accept files

## Install
Run `go get github.com/calexa22/randall`

## Requirements 
The Randall module requires Go version `>=1.19`

## Usage
```go
package main

import (
	"time"
	
	"github.com/calexa22/randall"
	"github.com/shopspring/decimal"
)

func main() {

    // Create randall.Client instance
    client :=  randall.NewClient(
		"MyHarvestAccountId",
		"MyApiToken",
		"MyAppName", // for the User-Agent
		"MyEmail", // for the User-Agent
	)

    // Retrieve the currently authenticated user object
    myUser, err := client.Users.Me()

    if err != nil {
		panic(err)
	}

    // Retrives all time entries accessible to the authenticated user
	// starting from 11-24-2022, with pagination (page 2, 20 entries per page)
	timeEntriesParams := randall.GetTimeEntriesParams{
		FromDate: randall.OptionalTime(time.Date(2022, time.November, 24, 0, 0, 0, 0, time.UTC)),
		Page: randall.OptionalInt(2),
		PerPage: randall.OptionalInt(20),
	}

    if err != nil {
		panic(err)
	}
	
	timeEntries, err := client.TimeEntries.GetAll(timeEntriesParams)

    // Create a new time entry under the given project/task
	// via hours spent
	hours := decimal.NewFromInt32(8)
	durationEntry := randall.CreateTimeEntryViaDurationRequest{
		TaskId:    2222222,
		ProjectId: 1111111,
		SpentDate: time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC),
		Hours:     randall.OptionalDecimal(hours),
	}
	
	newEntry, err := client.TimeEntries.CreateViaDuration(durationEntry)
	
	if err != nil {
		panic(err)
	}
	
	// ...
}
```

## Notes

* To avoid precision loss, decimal properties in randall are serliaized/deserialized as strings and implemented via the [shopspring/decimal](https://github.com/shopspring/decimal#readme) go library.
* The majority of requests sent to the Harvest API are sent as JSON. However in the case of endpoints that may take a file, if a file is specified the entire request body is encoded as `multipart/form-data` per the documentation.
* Detailed explanations of every endpoint and their requests, as well as example CURL requests can be found in the official [Harvest documentation](https://help.getharvest.com/api-v2/).

Happy Tracking!
