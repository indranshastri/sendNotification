package sendNotification
 
import (
    "fmt"
    "github.com/anachronistic/apns"
    "github.com/NaySoftware/go-fcm"
)


func SendPushToIosClient(pushText string,pushToken string,files string)  {
       fmt.Println("SendPushToClient")
       fmt.Println("pushText: ", pushText)
       fmt.Println("pushToken: ", pushToken)
       payload := apns.NewPayload()
       payload.Alert = pushText
       pn := apns.NewPushNotification()
       pn.DeviceToken = pushToken
       pn.AddPayload(payload)
       client := apns.NewClient("gateway.sandbox.push.apple.com:2195",files , files)
       resp := client.Send(pn)
       alert, _ := pn.PayloadString()
       fmt.Println("Alert:", alert)
       fmt.Println("Success:", resp.Success)
       fmt.Println("Error:", resp.Error)
}

func SendPushToAndroidClient(serverKey string,ids []string,data map[string]string,xds []string) {
  c := fcm.NewFcmClient(serverKey)
  c.NewFcmRegIdsMsg(ids, data)
  c.AppendDevices(xds)
  
  status, err := c.Send()

	if err == nil {
    	status.PrintResults()
	} else {
		fmt.Println(err)
	}
}


func ShowTheMessage(message string) {
  fmt.Println(message)
}

