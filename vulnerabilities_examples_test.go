package ionic_test

import (
	"fmt"

	"github.com/ion-channel/ionic"
	"github.com/ion-channel/ionic/pagination"
)

func ExampleIonClient_GetVulnerabilities() {
	client, err := ionic.New("https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	vulns, err := client.GetVulnerabilities("jdk", "", "atoken", pagination.AllItems)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Vulnerabilities: %v\n", vulns)
}

func ExampleIonClient_GetVulnerabilities_version() {
	client, err := ionic.New("https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	vulns, err := client.GetVulnerabilities("jdk", "1.7.0", "atoken", pagination.AllItems)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Vulnerabilities: %v\n", vulns)
}

func ExampleIonClient_GetVulnerability() {
	client, err := ionic.New("https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	vuln, err := client.GetVulnerability("CVD-2014-0030", "atoken")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Vulnerability: %v\n", vuln)
}
