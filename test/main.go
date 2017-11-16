package main

import (
	"time"
	"github.com/zyl0501/go-push-client/push"
	"fmt"
)

func main() {
	connectClient := push.ConnectClient{}
	pushClient := push.PushClient{}
	pushClient.Init()
	pushClient.Start()

	FakeBizProcess(&pushClient)

	defer connectClient.Close()
}

func FakeBizProcess(pushClient *push.PushClient) {
	for {
		time.Sleep(time.Second * 5)
		fmt. Println("tick...")
	}
}
