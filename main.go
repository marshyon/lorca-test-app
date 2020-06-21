package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/zserge/lorca"
)

func main() {
	args := []string{}
	prepareArgsForLorcaBootstrap(args)

	// create and launch the app
	ui, err := lorca.New("", "", 1100, 800, args...)
	genericErrHandler(err, "initializing the app UI")
	defer ui.Close()

	err = ui.Bind("onStart", func(dateInString string) {
		// perform the server side task here
		fmt.Printf(">>>>>>> app started and recieved message from client [%s]\nsending timestamp back to client [%v]\n", dateInString, time.Now().Unix())
		dteStr := strconv.Itoa(int(time.Now().Unix()))
		ui.Eval(fmt.Sprintf(`onStartAck(%s);`, dteStr))
	})

	genericErrHandler(err, "binding event listener - app started")

	err = ui.Bind("onActivate", func(dateInString string) {
		// perform the server side task here
		fmt.Printf(">>>>> message recieved from front end : [%s]\n", dateInString)

		dteStr := strconv.Itoa(int(time.Now().Unix()))

		sillyName := randomdata.SillyName()
		fmt.Printf(">>>>>> sending time stamp to client of [%s] and random name of [%s]\n", dteStr, sillyName)
		ui.Eval(fmt.Sprintf(`onMessage("%v","%s");`, dteStr, sillyName))
	})

	genericErrHandler(err, "binding event listener - activate")

	// connect to FS (fileServer pointing to folder www)
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	genericErrHandler(err, "connecting to the fileServer (e.g. www folder)")
	defer listener.Close()

	// start the server by binding the listener
	go http.Serve(listener, http.FileServer(FS))

	err = ui.Load(fmt.Sprintf("http://%s", listener.Addr()))
	genericErrHandler(err, "load the index.html")

	// os signal handling
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	// can exit now
	fmt.Println("Thanks for using the app!")
}

func genericErrHandler(err error, description ...string) {
	if err != nil {
		if description != nil {
			fmt.Println(fmt.Sprintf("oops! something is wrong! %v\n", description[0]))
		}
		panic(err)
	}
}

/**
 *	prepare bootstrap arguments for different OS (for the moment, only Linux)
 */
func prepareArgsForLorcaBootstrap(args []string) []string {
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	return args
}
