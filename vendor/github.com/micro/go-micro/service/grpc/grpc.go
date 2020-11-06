package grpc

import (
	"time"

	"github.com/micro/go-micro"
	broker "github.com/micro/go-micro/broker"
	client "github.com/micro/go-micro/client/grpc"
	server "github.com/micro/go-micro/server/grpc"
)

/*
// Registry sets the registry for the service
// and the underlying components
func Registry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r

		// Update Client and Server
		//Client是抽象的接口
		o.Client.Init(client.Registry(r))
		o.Server.Init(server.Registry(r))
		// Update Selector
		o.Client.Options().Selector.Init(selector.Registry(r))
		// Update Broker
		o.Broker.Init(broker.Registry(r))
	}
}

// Selector sets the selector for the service client
func Selector(s selector.Selector) Option {
	return func(o *Options) {
		o.Client.Init(client.Selector(s))
	}
}
*/
// NewService returns a grpc service compatible with go-micro.Service
//创建GRPC服务
func NewService(opts ...micro.Option) micro.Service {
	//既创建client,也创建server,估计是调用其他，或被调用
	// our grpc client
	c := client.NewClient()
	// our grpc server
	s := server.NewServer()
	// our grpc broker
	b := broker.NewBroker()

	// create options with priority for our opts
	//重新设置service的option
	options := []micro.Option{
		micro.Client(c),
		micro.Server(s),
		micro.Broker(b),
	}

	// append passed in opts
	options = append(options, opts...)

	// generate and return a service
	return micro.NewService(options...)
}

// NewFunction returns a grpc service compatible with go-micro.Function
func NewFunction(opts ...micro.Option) micro.Function {
	// our grpc client
	c := client.NewClient()
	// our grpc server
	s := server.NewServer()
	// our grpc broker
	b := broker.NewBroker()

	// create options with priority for our opts
	options := []micro.Option{
		micro.Client(c),
		micro.Server(s),
		micro.Broker(b),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second * 30),
	}

	// append passed in opts
	options = append(options, opts...)

	// generate and return a function
	return micro.NewFunction(options...)
}
