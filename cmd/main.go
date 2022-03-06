package main

import (
	"flag"

	"DailyReportAPI/pkg/server"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":1323", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Serve(addr)
}
