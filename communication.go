
//created by : OPH - PASTI
package main

import (
	"fmt"
	"time"
)

//Struktur Request 
type Request struct {
	ID      int
	Payload string
	Reply   chan<- string
}

//Menangani request dan mengirimkan balasan yang diterima
func Replier(requests <-chan Request) {
	for req := range requests {
		// Time Processing
		time.Sleep(time.Second)
		// Formulating balasan (response)
		reply := fmt.Sprintf("Request %d: Payload yg diterima - %s", req.ID, req.Payload)
		// Mengirimkan balasan melalui channel
		req.Reply <- reply
	}
}

func main() {
	// membuat channel request
	requests := make(chan Request)

	go Replier(requests)

	// membuat request
	for i := 1; i <= 5; i++ {
		// Membuat channel untuk menerima balasan
		replyChannel := make(chan string)
		// Membuat request
		request := Request{
			ID:      i,
			Payload: fmt.Sprintf("Payload Request %d", i),
			Reply:   replyChannel,
		}

		// Kirim request melalui channel
		requests <- request

		// Tunggu dan prin hasil balasan
		reply := <-replyChannel
		fmt.Println(reply)
	}

	// Close request untuk menghentikan replier
	close(requests)
}
