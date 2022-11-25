package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/*
一个TCP客户端进行TCP通信的流程如下：

1.建立与服务端的链接
2.进行数据收发
3.关闭链接
*/

//1.与服务器建立连接
//2.进行连接数据收发
//2.1从终端读取信息
//2.2发送到服务器
//2.3接收服务器信息
func Tcp(s string) {
	//1.与服务器建立连接
	conn, err := net.Dial("tcp", s)
	if err != nil {
		fmt.Printf("dial failed,err:%v", err)
		return
	}
	//2.根据链接进行数据收发

	//2.1创建一个从终端读取的对象
	input := bufio.NewReader(os.Stdin)
	//for就是一直发送接收
	for {
		//2.2利用从终端读取信息的对象读到换行
		ss, _ := input.ReadString('\n')
		//2.3把读取的信息空格去掉
		strings.TrimSpace(ss)
		//2.4如果ss内容是大写就退出
		if strings.ToUpper(ss) == "Q" {
			return
		}
		//2.5如果ss内容不是q，就将终端的内容写入链接
		n, err := conn.Write([]byte(ss))
		if err != nil {
			fmt.Printf("send failed,err:%v\n", err)
		}
		//fmt.Println(n)
		//2.6从服务端读消息
		//2.6.1这只一次读多少内容
		var buf [1024]byte

		//2.6.2开始读取信息
		n, err = conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed,err:%v", err)
		}
		//fmt.Printf("读取了%d字节\n", n)
		fmt.Printf("收到服务端回复：%v", string(buf[:n]))
		//3.关闭链接

	}
}

func CheckAddrConnect(addr string) (bool, error) {
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return false, err
	}
	if conn != nil {
		_ = conn.Close()
		return true, nil
	}
	return false, errors.New("addr连接错误")
}
