package main

import (
	"encoding/json"
	"fmt"
	toxiproxy "github.com/Shopify/toxiproxy/client"
	xln_utils "github.com/rexlien/go-utils/xln-utils"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	xln_utils.Hello()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	finished := make(chan bool, 1)
	go func() {
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)
		finished <- true
	}()

	command := exec.Command("toxiproxy-server")
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
	case <-time.After(5 * time.Second):
		panic("timeout")
	case err := <-done:
		panic(err)
	default:

	}

	client := toxiproxy.NewClient("127.0.0.1:8474")
	var config []toxiproxy.Proxy
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &config)
	proxies, err := client.Populate(config)
	if err != nil {
		panic(err)
	}
	log.Println(len(proxies))


	<- finished



}
