package connectpool

import (
	"bufio"
	"encoding/binary"
	"fmt"
	silenceperpool "github.com/silenceper/pool"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func get() []byte {
	buf := []byte("IDENTIFY\n")
	bufs := []byte(`{"client_id":"test"}`)

	var l = make([]byte, 4)
	binary.BigEndian.PutUint32(l, uint32(len(bufs)))
	buf = append(buf, l...)

	buf = append(buf, bufs...)
	return buf
}

func getV2() []byte {
	buf := []byte("  V2")
	return buf
}

func Benchmark_silenceper_pool(b *testing.B) {

	//factory 创建连接的方法
	factory := func() (interface{}, error) {
		return net.Dial("tcp", "10.60.6.35:4160")
	}

	//close 关闭连接的方法
	close := func(v interface{}) error {
		return v.(net.Conn).Close()
	}

	//ping 检测连接的方法
	//ping := func(v interface{}) error { return nil }

	//创建一个连接池： 初始化5，最大空闲连接是20，最大并发连接30
	poolConfig := &silenceperpool.Config{
		InitialCap: 5,  //资源池初始连接数
		MaxIdle:    20, //最大空闲连接数
		MaxCap:     30, //最大并发连接数
		Factory:    factory,
		Close:      close,
		//Ping:       ping,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}

	p, err := silenceperpool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("err=", err)
	}

	//从连接池中取得一个连接
	v, err := p.Get()
	if err != nil {
		fmt.Println("err=", err)
		os.Exit(1)
	}

	//do something
	//conn=v.(net.Conn)
	_, err = v.(net.Conn).Write([]byte("  V2"))
	if err != nil {
		fmt.Println("err=", err)
	}
	//将连接放回连接池中
	p.Put(v)

	//释放连接池中的所有连接
	//p.Release()

	//查看当前连接中的数量
	current := p.Len()
	fmt.Println("current=", current)

	for i := 0; i < b.N; i++ {
		v, err := p.Get()
		if err != nil {
			continue
		}
		var tmp = make([]byte, 0, 1024)

		fmt.Println(string(tmp))
		_, err = v.(net.Conn).Write(get())
		if err != nil {
			continue
		}

		_, err = v.(net.Conn).Read(tmp)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(tmp))
		p.Put(v)
	}

}
func Benchmark_without_silenceper_pool(b *testing.B) {
	// 连接服务器
	conn, err := net.Dial("tcp", "10.60.6.35:4150")
	if err != nil {
		fmt.Println("Connect to TCP server failed ,err:", err)
		return
	}
	defer conn.Close()
	// 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)

	// 一直读取直到遇到换行符
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Read from console failed,err:", err)
			return
		}

		// 读取到字符"Q"退出
		str := strings.TrimSpace(input)
		if str == "Q" {
			break
		}

		// 响应服务端信息
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Write failed,err:", err)
			break
		}
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("Write failed,err:", err)
			break
		}
		fmt.Println(n)

		//fmt.Println(string(buf))
	}

}
