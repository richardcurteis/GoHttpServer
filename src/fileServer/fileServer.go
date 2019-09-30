package fileServer

import (
	"../configs"
	"crypto/tls"
	"log"
	"net/http"

)

func Run(port string, directory string) {
	keyPath := "certificates/simpleServer/"
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(directory)))

	config := configs.GetTlsConfigNoCer()

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		TLSConfig:    config,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS(keyPath + "tls.cert", keyPath +  "tls.key"))
}
