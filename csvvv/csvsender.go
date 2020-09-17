package csvvv

import (
	"bufio"
	"github.com/diguacheng/csvtransporter/udppp"
	"time"

	"fmt"
	"io"
	"os"
)




func ReadCSVFile(path string) {
	fileobj, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file failed: %v", err)
		return 
	}
	// 关闭文件
	defer fileobj.Close()
	// 创建一个从文件中读内容的对象
	Reader := bufio.NewReader(fileobj)
	_, err = Reader.ReadBytes('\n')
	// 排除第一行表头 
	if err == io.EOF {
		fmt.Println("已经发送完 ！ ")
		return 
	}
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	var data []byte

	for {
		data, err = Reader.ReadBytes('\n')
		udppp.Streamsend<- data
		if err == io.EOF {
			close(udppp.Streamsend)
			time.Sleep(time.Duration(time.Second*1))
			fmt.Println("已经发送完 ！")
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			return 
		}
	}
	_=fileobj.Close()
	time.Sleep(2)


}


