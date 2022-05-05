package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

//iServer的接口实现，定义一个Server的服务器模块

type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的IP版本
	IPVersion string
	//服务器监听的IP
	IP string
	//服务器监听的端口
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s, Port %d, is starting\n", s.IP, s.Port)
	//1 获取一个TCP的Addr
	addr, err := net.ResolveIPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addt error : ", err)
		return
	}

	//2 监听服务器的地址

	//3 阻塞的等待客户端进行连接，处理客户端链接业务
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

}

//初始化Server模块的方法

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8889,
	}
	return s
}
