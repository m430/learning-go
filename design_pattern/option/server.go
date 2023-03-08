package option

import (
	"crypto/tls"
	"time"
)

// 服务器初始化参数可能很多
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

// 定义选项函数类型
type Option func(*Server)

// 这里利用闭包的特性，每次都返回Option类型
func WithProtocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}

func WithTLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

// 然后我们用一个工厂方法来创建Server，这里Addr和port通常为必填参数，所以构造时候必须指定
// 然后利用可变参数可以对可选参数进行分别设置
func NewServer(addr string, port int, options ...Option) (*Server, error) {
	// 设置必填参数和可选参数默认值
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	// 对外部可选项进行设置
	for _, opt := range options {
		opt(&srv)
	}
	return &srv, nil
}
