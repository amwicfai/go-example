package main

import (
	"517na.com/demo-thrift-idl/blog"
	"517na.com/demo-thrift-idl/blogservice"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"runtime"
)

const (
	NetworkAddr = "127.0.0.1:19090"
)

type BlogServiceImpl struct {
}

func (this *BlogServiceImpl) TestCase1(num1 int32, num2 int32, num3 string) (r int32, err error) {
	fmt.Println("-->TestCase1:")
	return num1 * num2, nil
}

func (this *BlogServiceImpl) TestCase2(num1 map[string]string) (r []string, err error) {
	fmt.Println("-->TestCase2:")
	return nil, nil
}

func (this *BlogServiceImpl) TestCase3() (err error) {
	fmt.Println("-->TestCase3:")
	return nil
}

func (this *BlogServiceImpl) TestCase4(blog []*blog.Blog) (err error) {
	fmt.Println("-->TestCase4:")
	return nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	//不同的goroutine都使用同一个handler实例，不用锁么，测试一下便知
	handler := &BlogServiceImpl{}
	processor := thrift.NewTMultiplexedProcessor()
	processor.RegisterProcessor("BlogService", blogservice.NewBlogServiceProcessor(handler))

	//TFramedTransport性能测试
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
