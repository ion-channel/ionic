package ionic_test

import (
	"fmt"

	"github.com/ion-channel/ionic"
)

func ExampleIonClient_Login() {
	// In theory you should not have an API key yet, so providing blank will
	// work just fine
	client, err := ionic.New("https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	sess, err := client.Login("someusername", "supersecretpassword")
	if err != nil {
		fmt.Println(err.Error())
	}

	// Use the bearer token in subsequent calls
	vuln, _ := client.GetVulnerability("CVE-1234-1234", sess.BearerToken)
	fmt.Printf("Vulns: %v\n", vuln)
}
