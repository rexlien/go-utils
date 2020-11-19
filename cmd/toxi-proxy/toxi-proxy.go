package main

import (
	"context"
	"encoding/json"
	toxiproxy "github.com/Shopify/toxiproxy/client"
	"github.com/gin-gonic/gin"
	"github.com/rexlien/go-utils/go-utils/xln-proto/build/gen/proxypb/proxypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

import _ "github.com/gin-gonic/gin"



type RequestProcess struct {

	request *proxypb.ProxyMatchRequest
	response chan *proxypb.ProxyMatchResponse

}

type server struct {
	proxypb.UnimplementedProxyServiceServer
	reqChannel chan *RequestProcess
	//responseChannel chan *proxypb.ProxyMatchResponse
}

func (s* server) MatchProxy(context context.Context, req *proxypb.ProxyMatchRequest) (*proxypb.ProxyMatchResponse, error) {

	process := &RequestProcess{request: req, response: make (chan *proxypb.ProxyMatchResponse)}
	s.reqChannel <- process


	response := <- process.response
	return response, nil
}


func main() {

	path := os.Getenv("XLN_TOXIPROXY_CONFIG_PATH")
	if len(path) == 0 {
		path = "config.json"
	}

	httpPort := os.Getenv("XLN_TOXI_PROXY_HTTP_PORT")
	if httpPort == "" {
		httpPort = ":38474"
	}

	grpcPort := os.Getenv("XLN_TOXI_PROXY_GRPC_PORT")
	if grpcPort == "" {
		grpcPort = ":39090"
	}



	//appPort, err := strconv.Atoi(os.Getenv("XLN_TOXI_PROXY_APP_PORT"))
	//if(err != nil) {
	//	appPort = 8080
	//}



	//proxypb.Config{}

/*
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	finished := make(chan bool, 1)
	go func() {
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)
		finished <- true
	}()
*/
	command := exec.Command("cmd")

	//var outb, errb bytes.Buffer
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Start()
	if err != nil {
		panic(err)
	}
	defer func() {
		log.Println("do kill")
		err = command.Process.Kill()
		if err != nil {
			log.Println(err)
		}
	}()

	done := make(chan error, 1)
	go func() {
		done <- command.Wait()
	}()



	select {
	case <- time.After(3 * time.Second):
	case err := <-done:
		if  err != nil &&  err.(*exec.ExitError) != nil {
			//log.Println(errb.String())
			panic(err)
		}

	}
	//time.Sleep(3 * time.Second)

	client := toxiproxy.NewClient("127.0.0.1:8474")
	var config []toxiproxy.Proxy
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &config)
	proxies, err := client.Populate(config)
	if err != nil {
		panic(err)
	}
	log.Println(len(proxies))

	//client.Proxies().
	for _, proxy := range config  {
		proxyName := proxy.Name
		for _, toxics := range proxy.ActiveToxics {
			proxyMap, _ := client.Proxies()
			proxyMap[proxyName].AddToxic(toxics.Name, toxics.Type, toxics.Stream, toxics.Toxicity, toxics.Attributes)
		}

	}

	proxyReqChannel := make(chan *RequestProcess, 50)

	go func() {
		list, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()

		reflection.Register(s)
		proxypb.RegisterProxyServiceServer(s, &server{reqChannel: proxyReqChannel})
		if err:= s.Serve(list); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}

	}()




	go func() {

		client := toxiproxy.NewClient("127.0.0.1:8474")

		//freeList := list.New()
		portMap := make(map[string]string)

		initPort := 55000

		for {
			select {
				case process := <-proxyReqChannel:
					{
						//retProxy := make([]*proxypb.Proxy, 0)
						for _, proxy := range process.request.Proxies {
							for i, host:= range proxy.UpStreamHosts {
								nextListen := ""

								mapKey := proxy.Name + strconv.Itoa(i)

								nextListen, ok := portMap[proxy.Name + strconv.Itoa(i)]
								if  !ok {
									nextListen = "127.0.0.1:" + strconv.Itoa(initPort)
									portMap[mapKey] = nextListen
									initPort++
								}

								_, err := client.CreateProxy(proxy.Name, nextListen, host)
								if err != nil {
									log.Println("proxy create failed")
								}

								//fixed to matched listen proxy for response
								proxy.UpStreamHosts[i] = nextListen
							}
						}
						response := &proxypb.ProxyMatchResponse{Proxies: process.request.Proxies}
						process.response <- response

					}
			}
		}



	}()


//	<- finished
	r := gin.Default()
	r.Run(httpPort)



}
