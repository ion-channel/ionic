package ionic_test

import (
	"fmt"

	"github.com/ion-channel/ionic"
)

func ExampleIonClient_Login() {
	client, err := ionic.New("https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	sess, err := client.Login("someusername", "supersecretpassword")
	if err != nil {
		fmt.Println(err.Error())
	}

	vuln, _ := client.GetVulnerability("CVE-1234-1234", sess.BearerToken)
	fmt.Printf("Vulns: %v\n", vuln)
}
