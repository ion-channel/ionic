package ionic_test

import (
	"fmt"

	"github.com/ion-channel/ionic"
)

func ExampleIonClient_getProduct() {
	client, err := ionic.New("https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	ps, err := client.GetProducts("cpe:/a:ruby-lang:ruby:1.8.7", "someapikey")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Products: %v\n", ps)
}

func ExampleIonClient_getRawProduct() {
	client, err := ionic.New("https://api.test.ionchannel.io")
	if err != nil {
		panic(fmt.Sprintf("Panic creating Ion Client: %v", err.Error()))
	}

	bodyBytes, err := client.GetRawProducts("cpe:/a:ruby-lang:ruby:1.8.7", "someapikey")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Raw Products: %v\n", string(bodyBytes))
}
