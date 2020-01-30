package main

import (
	"flag"
	"github.com/kevin-monteiro/http-stub-provider/pkg/server"
	"github.com/kevin-monteiro/http-stub-provider/pkg/types"
)

func main() {
	addr := flag.String("addr", "", "Adress the admin service will bind to. Default to localhost, set to 0.0.0.0 to use from another machine")
	port := flag.Int("port", 4770, "Port of stub admin service")
	stub := flag.String("stub", "", "Path where the stub files are (Optional)")

	flag.Parse()
	// run admin stub server
	server.StartStubServer(&types.Server{Addr: *addr, Port: *port, StubPath: *stub})
}
