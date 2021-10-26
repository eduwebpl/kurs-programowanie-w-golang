package main

import "fmt"

func main() {
	completionChannel := make(chan bool)
	go runOnSeparatedThread(completionChannel)
	<-completionChannel
	// lub 	value := <- completionChannel

}

func runOnSeparatedThread(completionChannel chan bool) {
	fmt.Println("Opened on separeted thread!!!")
	completionChannel <- true
}

// ---- Race condition ----

// package main

// import "fmt"
// import "time"

// var sharedInt int = 0
// var unusedValue int = 0

// func runSimpleReader() {
// 	for {
// 		var value int = sharedInt
// 		if value % 10 == 0 {
// 			unusedValue = unusedValue + 1
// 		}
// 		fmt.Println(sharedInt, unusedValue)
// 	}
// }

// func runSimpleWriter() {
// 	for {
// 		sharedInt = sharedInt + 1
// 	}
// }

// func main() {
// 	go runSimpleReader()
// 	go runSimpleWriter()
// 	time.Sleep(time.Second)
// }

// ---- Thread safety ----

// package main

// import (
// 	"sync/atomic"
// 	"time"
// 	"fmt"
// )

// var sharedInt int64 = 0
// var unusedValue int = 0

// func runReader() {
// 	for {
// 		var value int64 = atomic.LoadInt64(&sharedInt)
// 		if value % 10 == 0 {
// 			unusedValue = unusedValue + 1
// 		}
// 		fmt.Println(sharedInt, unusedValue)
// 	}
// }

// func runWriter() {
// 	for {
// 		atomic.AddInt64(&sharedInt, 1)
// 	}
// }

// func main() {
// 	go runReader()
// 	go runWriter()
// 	time.Sleep(time.Second)
// }
