package main

import (
	"time"
	"github.com/zyl0501/go-push-client/push"
	"fmt"
)

func main() {
	connectClient := push.ConnectClient{}
	pushClient := push.NewPushClient()
	pushClient.Init()
	pushClient.Start()

	FakeBizProcess(pushClient)

	defer connectClient.Close()
}

func FakeBizProcess(pushClient *push.PushClient) {
	bind := false
	fast := false
	for {
		time.Sleep(time.Second * 5)

		if !bind {
			bind = true
			pushClient.BindUser("user-0", "")
		}

		if !fast {
			fmt.Println("<<<<<<wait for test fast connect")
			time.Sleep(time.Second * 5)
			fast = true
			pushClient.TestFastConnect()
		}
	}
}
