package main

import (
	"github.com/diguacheng/csvtransporter/csvvv"
	"github.com/diguacheng/csvtransporter/udppp"
	"flag"
	"fmt"
)

//39.97.229.151:9982
func main() {
	var path = flag.String("filename", "movie_25.csv ", "name of the file to be sent")
	var address = flag.String("address", "127.0.0.1:9982", "the remote address")
	flag.Parse()
	fmt.Println(*path, *address)
	go csvvv.SavetoFile("./return"+*path)
	go csvvv.ReadCSVFile("./" + *path)
	udppp.SendMessage(*address)
}
