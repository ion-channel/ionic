package ionic_test

import (
	"fmt"

	"github.com/ion-channel/ionic"
)

func ExampleIonClient_new() {
	client, err := ionic.New("https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	// Client can then call for the different actions
	client.GetProject("projectID", "teamID", "someapikey")
}
