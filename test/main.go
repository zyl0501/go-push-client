package main

import (
	"time"
	"github.com/zyl0501/go-push-client/push"
)

func main() {
	connectClient := push.ConnectClient{}
	pushClient := push.PushClient{ConnClient: connectClient}
	pushClient.Start()

	FakeBizProcess(&pushClient)

	defer connectClient.Close()
}

func FakeBizProcess(pushClient *push.PushClient) {
	for {
		time.Sleep(time.Second * 5)
		pushClient.Send()
	}
}
