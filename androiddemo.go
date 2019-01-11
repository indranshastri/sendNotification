package main

import (
	"fmt"
	"github.com/indranshastri/sendNotification"
)


func main(){

	fmt.Println("sending notificaion on android")
	serverKey := "key"
	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	ids := []string{
		"token1",
	}

	xds := []string{
		"another tokens",
	}
	sendNotification.SendPushToAndroidClient(serverKey,ids,data,xds)


}