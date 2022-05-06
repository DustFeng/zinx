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
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addt error : ", err)
		return
	}

	//2 监听服务器的地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen ", s.IPVersion, "err ", err)
		return
	}

	fmt.Println("start Zinx server succ, ", s.Name, " succ, Listenning...")

	//3 阻塞的等待客户端进行连接，处理客户端链接业务
	

	for {
		//如果有客户端链接过来，阻塞会返回
		conn, err := listenner.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		//已经与客户端建立间接，做一些业务，做一个最基本的最大512字节长度的回显业务
		go func ()  {
			for {
				buf := make([]byte, 512)
				cnt, err ：=conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err ", err)
					continue
				}

				//回显功能
				if _, err := conn.write(buf[:cnt]); err != nil {
					fmt.Println("write back buf err", err)
					continue
				}
			}	
		}()
	}

}

func (s *Server) Stop() {
	//将一些服务器的资源,
}

func (s *Server) Serve() {
	//启动server的服务功能
	s.Start()

	//做一些启动服务器之后的额外业务

	//阻塞状态
	select{}
}

//初始化Server模块的方法

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:    		"0.0.0.1"
		Port:      8999,
	}
	return s
}
