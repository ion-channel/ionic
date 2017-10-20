package ionic_test

import (
	"fmt"

	"github.com/ion-channel/ionic"
)

func ExampleIonClient_GetProduct() {
	client, err := ionic.New("apikey", "https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	product, err := client.GetProduct("cpe:/a:ruby-lang:ruby:1.8.7")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Product: %v\n", product)
}

func ExampleIonClient_GetRawProduct() {
	client, err := ionic.New("apikey", "https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	bodyBytes, err := client.GetRawProduct("cpe:/a:ruby-lang:ruby:1.8.7")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Product: %v\n", string(bodyBytes))
}
