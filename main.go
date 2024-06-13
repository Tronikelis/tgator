package main

import (
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	server := http.NewServeMux()

	server.HandleFunc("POST /agg/new", func(rw http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			// todo
		}

		ip, err := net.ResolveUDPAddr("udp", r.RemoteAddr)
		if err != nil {
			// todo
		}

		slog.Info(string(bodyBytes))
		slog.Info(ip.IP.String())
	})

	slog.Info("listening")

	if err := http.ListenAndServe("localhost:3000", server); err != nil {
		panic(err)
	}
}
