package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"zinx_study1/znet"
)

func main() {
	fmt.Println("client start ...")
	time.Sleep(1 * time.Second)
	// 连接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:3001")
	if err != nil {
		log.Printf("client start error: %s\n", err)
		return
	}
	for {
		// 调用 write 写数据
		err = SendPackedData(&conn, 2, []byte("zinx v0.5.1 test packed message"))
		if err != nil {
			log.Printf("client write data error: %s\n", err)
			continue
		}
		// 调用完 wirte 后，可以接着从连接中读取数据
		err = ReceivePackedData(&conn)
		if err != nil {
			log.Printf("client read data error: %s\n", err)
			continue
		}
		time.Sleep(2 * time.Second)
	}
}

// 普通的发送数据到服务端
func NormalSend(conn *net.Conn, data []byte) error {
	if _, err := (*conn).Write(data); err != nil {
		return err
	}
	return nil
}

func ReceiveNormalData(conn *net.Conn) error {
	buf := make([]byte, 512)
	if _, err := (*conn).Read(buf); err != nil {
		return err
	}
	log.Printf("ReceivePackedData: %s\n", string(buf))
	return nil
}

// 发送封包消息，到服务端
func SendPackedData(conn *net.Conn, msgId uint32, data []byte) error {
	dp := znet.NewDataPacker()
	binaryMsg, err := dp.Pack(znet.NewMessage(msgId, data))
	if err != nil {
		return err
	}
	_, err = (*conn).Write(binaryMsg)
	if err != nil {
		return err
	}
	return nil
}

// 接收并**解包**数据
func ReceivePackedData(conn *net.Conn) error {
	// 接收 head msg
	dp := znet.NewDataPacker()
	headBuf := make([]byte, dp.GetHeaderLen())
	if _, err := io.ReadFull(*conn, headBuf); err != nil {
		return err
	}
	msg, err := dp.UnPack(headBuf)
	if err != nil {
		return err
	}
	msgObj := msg.(*znet.Message)
	msgBuf := make([]byte, msgObj.DataLen)
	if _, err := io.ReadFull(*conn, msgBuf); err != nil {
		return err
	}
	log.Printf("ReceivePackedData-receive data from server. Id: %d\t %s\n", msgObj.GetMsgId(), string(msgBuf))
	return nil
}
