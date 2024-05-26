/*
package main

import (

	"fmt"
	"time"

)

	func tik() {
		for i := 0; i < 30; i++ {
			fmt.Print(".")
			time.Sleep(time.Millisecond)
		}
	}

	func tok() {
		for i := 0; i < 30; i++ {
			fmt.Print("o")
			time.Sleep(time.Millisecond)
		}
	}

	func main() {
		go tik()
		go tok()
		time.Sleep(500 * time.Millisecond)
	}
*/
package main

import (
	"fmt"
	"time"
)

func tak(s string) {
	for i := 0; i < 30; i++ {
		fmt.Print(s)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	go tak(".")
	go tak("o")
	time.Sleep(2000 * time.Millisecond)
}
