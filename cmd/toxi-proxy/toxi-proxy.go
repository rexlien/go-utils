package main

import (
	"encoding/json"
	toxiproxy "github.com/Shopify/toxiproxy/client"
	"github.com/gin-gonic/gin"
	"github.com/rexlien/go-utils/go-utils/xln-proto/build/gen/proxypb/proxypb"

	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

import _ "github.com/gin-gonic/gin"

func main() {

	path := os.Getenv("XLN_TOXIPROXY_CONFIG_PATH")
	if len(path) == 0 {
		path = "config.json"
	}

	httpPort, err := strconv.Atoi(os.Getenv("XLN_TOXI_PROXY_HTTP_PORT"))
	if err != nil {
		httpPort = 38474
	}


	proxy := proxypb.Config{}
	proxy.String()

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

	err = command.Start()
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





//	<- finished
	r := gin.Default()
	r.Run(":" + strconv.Itoa(httpPort))



}
