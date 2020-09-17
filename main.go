package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/diguacheng/csvtransporter/csvvv"
	"github.com/diguacheng/csvtransporter/udppp"
)


func main() {
	var path = flag.String("filename", "movie_25.csv ", "name of the file to be sent")
	var address = flag.String("address", "127.0.0.1:9982", "the remote address")
	flag.Parse()
	fmt.Println(*path, *address)
	var wg sync.WaitGroup
	wg.Add(2)
	go csvvv.SavetoFile("./return"+*path)
	go csvvv.ReadCSVFile("./" + *path)
	udppp.SendMessage(*address)
	
}
