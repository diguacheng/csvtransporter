package csvvv

import (
	"bufio"
	"github.com/diguacheng/csvtransporter/udppp"
	"fmt"
	"os"
)


func checkFileisExist(pth string)bool{
	if _,err:=os.Stat(pth);os.IsExist(err){
		return true
	}
	return false
}


func SavetoFile(pth string) {
	var fileobj *os.File
	var err error
	if !checkFileisExist(pth){
		fileobj,err = os.Create(pth)
		if err!=nil{
			fmt.Println(err.Error())
			return
		}
	}else{
		fileobj,err = os.Open(pth)
		if err!=nil{
			fmt.Println(err.Error() )
		}
	}

	writer:=bufio.NewWriter(fileobj)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	for {
		select {
		case data,ok := <-udppp.Streamreceived:
			// fmt.Println(string(data),len(data),ok,"11111")
			if ok == false {
				return
			}
			_, err := writer.WriteString(string(data))
			_ = writer.Flush()
			if err != nil {
				return
			}
			fmt.Printf("写入数据%s \n", string(data))

		}
	}
}

// func save(ctx context.Context,writer *bufio.Writer){
// 	for {
// 		select {
// 		case data ,ok:=<-udppp.Streamreceived:
// 			if ok==false{
// 				return
// 			}
// 			_,err:=writer.WriteString( string(data) )
// 			_=writer.Flush()
// 			if err!=nil{
// 				return
// 			}
// 			fmt.Printf("写入数据%s \n",string(data))
//
// 	}
// }
