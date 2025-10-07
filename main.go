package main

import "time"

func main() {
	Sleep(5 * time.Second)
}

func Sleep(d time.Duration) {
	if d <= 0 {
		return
	}
	<-time.After(d)
}
