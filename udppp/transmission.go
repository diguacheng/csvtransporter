package udppp

import (
	"fmt"
	"net"
	"time"
)



var Streamreceived = make(chan []byte, 0)
var Streamsend =make(chan []byte,0)

func SendMessage(address string) {
	remoteAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("address is not valid")
		return
	}
	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(remoteAddr, "正常")
	// fmt.Printf("Local: <%s> \n", conn.LocalAddr().String())
	receivedata:= make([]byte,1024)
	for {
		select {
		case data,ok := <-Streamsend:
			if len(data)==0||ok==false{
				close(Streamreceived)
				time.Sleep(time.Second)
				return
			}
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("已发送：%s\n", string(data))
			n, err := conn.Read(receivedata)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("已收到：%s\n", string(receivedata[:n]))
			Streamreceived <- receivedata[:n]
		}

	}


}
