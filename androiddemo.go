package main

import (
	"fmt"
	"github.com/indranshastri/sendNotification"
)


func main(){

	fmt.Println("sending notificaion on android")
	serverKey := "AAAAjUV9kgs:APA91bGViKn11qyYuKUgY7AwpOmbDubCD96cf2_KYntRxo58MFC_Vy3rdnVTo-CiYr9h36K51emf0ZD9UIjRlpWZX0jScRUYQT_bZq_BxE9FeONC-myLlzAGYyMxXhuVIZTtGsx5ob39"
	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	ids := []string{
		"ce4zCvgurDw:APA91bEJqLvRbktQBA5GDarPKWVMmAiCT94lVLT9Odd8PLPFEczmIM544BMmZkwGsxgXOb0EdcReBRhlUg45J6mcDDJe6uOonlu-J5JMAiklEmzknN83HaGcu8CnxhTbrBQlTWL2zDK8",
	}

	xds := []string{
		"ce4zCvgurDw:APA91bEJqLvRbktQBA5GDarPKWVMmAiCT94lVLT9Odd8PLPFEczmIM544BMmZkwGsxgXOb0EdcReBRhlUg45J6mcDDJe6uOonlu-J5JMAiklEmzknN83HaGcu8CnxhTbrBQlTWL2zDK8",
	}
	sendNotification.SendPushToAndroidClient(serverKey,ids,data,xds)


}