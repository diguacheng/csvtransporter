package main
import (
	"fmt"
	"net"
)
func main() {
	ip := net.ParseIP("0.0.0.0")
	listener,err:=net.ListenUDP("udp",&net.UDPAddr{IP:ip,Port: 9982})
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	data:=make([]byte, 1024)
	for {
		n,remote,err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ip:%s data:%s",remote,string(data[:n]))
		newData:=append([]byte("hwllo wordld"),data[:n]...)
		_,err=listener.WriteToUDP(newData, remote)
		if err != nil {
			fmt.Println(err)
		}
		
	}
}