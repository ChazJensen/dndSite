package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"server/characters"
	"sync"
)

func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}
//*
func getCharData(w http.ResponseWriter, r *http.Request) {
	charName := r.URL.Query().Get("name")
	usrName := r.URL.Query().Get("usr")

	// charInfo will store a server/characters Char data-type
	char := characters.GetChar(usrName, charName)

	char.ExecuteTemplate(w)
}/**/

func makeServer() *http.Server {
	srv := new(http.Server)

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/char", getCharData)

	srv.Addr = ":8080"
	srv.Handler = mux

	return srv
}

func shell() {
	exit := false
	stdin := bufio.NewReader(os.Stdin)
	for !exit {
		print(">> ") // signify user has entered shell

		command, err  := stdin.ReadString('\n')
		if err != nil {
			panic(err)
		}

		command = command[:len(command)-1]

		switch command {
			case "help": {
				println("exit - stop server and exit main")
			}
			case "char": {
				c := characters.GetChar("chris", "dan")
				// io.WriteString(os.Stdout, c.Sheet)
				c.ExecuteTemplate(os.Stdout)
			}
			case "exit": {
				println("stopping server")
				srv.Close()
				exit = true
			}
			default : {
				println("invalid command, use help if stuck")
				continue
			}
		}
	}
}

var srv = makeServer()
func main() {
	// use wait group to run server on different goroutine
	var wg sync.WaitGroup
	wg.Add(1) // tell wg it has an unfinished process
	go func(s *http.Server) {
		s.ListenAndServe()
		wg.Done() // tell wg process has finished 
	} (srv)

	// main loop: shell; allows for server control
	shell()

	wg.Wait()
}
