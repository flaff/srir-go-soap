package main

import (
	"./webservice"
	"fmt"
)

func getMessage () string {
	return "hello didney warld";
}

func getUrl () string {
	return "http://127.0.0.1:8080";
}

func main () {
	url := getUrl()
	service := webservice.NewHello_PortType(url, false, nil)

	message := getMessage()
	response, err := service.SayHello(&message)

	fmt.Println("response");
	fmt.Println(&response);

	if (err != nil) {
		fmt.Println("err");
		fmt.Println(err);
	}
}