package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Interface string`json:interface`
	Upstreams []Upstream`json:"upstreams"`
}

type Upstream struct {
	Path string`json:"path"`
	Method string`json:"method"`
	Backends []string`json:"backends"`
	ProxyMethod string`json:"proxyMethod"`
}

var (
	app = cli.NewApp()
	name = "config.json"
)

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "balancer"
	app.Version = "1.0.0"
}


func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "To run",
			Action: func(c *cli.Context) {
				if c.Args().First() != "" {
					name = c.Args().Get(0)
					file, err := ioutil.ReadFile(name)
					if err == nil {
						start(&file)
						fmt.Printf("started with %s configuration\n", c.Args().Get(0))
					} else {
						fmt.Printf("%s configuration file not found\n", c.Args().Get(0))
					}
					//fmt.Printf("started with %s configuration\n", c.Args().Get(0))
				} else {
					fmt.Printf("started with default configuration\n")
					file, err := ioutil.ReadFile(name)
					if err == nil {
						start(&file)
						fmt.Printf("started with %s configuration\n", name)
					} else {
						fmt.Printf("%s configuration file not found\n", name)
					}
				}

			},
		},
		{
			Name:    "reload",
			Aliases: []string{"re"},
			Usage:   "Reload program",
			Action: func(c *cli.Context) {
				fmt.Printf("reload with %s configuration\n", name)
				file, err := ioutil.ReadFile(name)
				if err == nil {
					start(&file)
					fmt.Printf("started with %s configuration\n", c.Args().Get(0))
				} else {
					fmt.Printf("%s configuration file not found\n", c.Args().Get(0))
				}
			},
		},
	}
}

func start(file *[]byte){

	data := Config{}
	_= json.Unmarshal(*file, &data)

	roundId := 0

	fmt.Printf("%s\n", data)
	rtr := mux.NewRouter()
	srv := &http.Server{
		Addr:fmt.Sprintf("127.0.0.1%s", data.Interface),
		Handler: rtr,
	}

	//srv.Shutdown(context.Background())

	for _, upstream := range data.Upstreams {
		upstr := upstream
		rtr.HandleFunc(fmt.Sprintf("/%s", upstr.Path), func(writer http.ResponseWriter, request *http.Request) {
			//writer.Write([]byte(fmt.Sprintf("%d\n", upstream)))
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in start", r)
				}
			}()
			ch := make(chan []byte)
			if upstr.ProxyMethod == "round-robin" {
				go serve(upstr.Backends[roundId], upstr.Method, ch)
				select {
				case d := <-ch:
					writer.Write(d)
				case <-time.After(time.Minute):
					fmt.Println("Time out: No news in one minute")
				}

				fmt.Printf("round-robin: %d\n", roundId)
				roundId++
				roundId %= len(upstr.Backends)

			} else {
				for _, url := range upstr.Backends {
					go serve(url, upstr.Method, ch)
				}
				select {
				case d := <-ch:
					fmt.Printf("")
					writer.Write(d)
				case <-time.After(time.Minute):
					fmt.Println("Time out: No news in one minute")
				}
			}

			defer close(ch)
		})
	}

	srv.ListenAndServe()

}


func serve(url string, method string, ch chan []byte) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in start", r)
		}
	}()
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	fmt.Println(url)
	ch <- f
	return nil
}