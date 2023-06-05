package main

import (
	"fmt"
	"github.com/mark2185/pomogoro/network"
	"github.com/mark2185/pomogoro/timer"
	"os"
	"os/exec"
	"time"
)

func initTimer() {
	var t timer.Timer = timer.MakeTimer()
	t.Pause()

	go network.Listen(&t)

	for {
		fmt.Println(t.ToString())
		time.Sleep(time.Second)
		if t.GetMinutes() == 0 && t.GetSeconds() == 0 {
			time.Sleep(1 * time.Second)
			fmt.Println(t.ToString())
			time.Sleep(1 * time.Second)
			t.Pause()
			msg := ""
			switch t.GetState() {
			case timer.Work:
				msg = "Take a break!"
			case timer.Break:
				msg = "Break over!"
			}

			if err := exec.Command("canberra-gtk-play", "-i", "complete").Run(); err != nil {
				panic(err.Error())
			}
			if err := exec.Command("dunstify", "-u", "normal", msg).Run(); err != nil {
				panic(err.Error())
			}

			t.Switch()
			fmt.Println(t.ToString())
			t.Pause()
			time.Sleep(3 * time.Second)
			t.Resume()
		}
		if t.IsRunning() {
			t.Tick()
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		initTimer()
	} else {
		switch os.Args[1] {
		case "start":
			initTimer()
		default:
			conn := network.Connect()
			fmt.Fprintf(conn, os.Args[1])
		}
	}
}
