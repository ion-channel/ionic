package ionic_test

import (
	"fmt"

	"github.com/ion-channel/ionic"
)

func ExampleIonClient_New() {
	client, err := ionic.New("apikey", "https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	// Client can then call for the different actions
	client.GetProject("projectID", "teamID")
}
