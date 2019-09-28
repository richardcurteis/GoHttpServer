package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// TLS on by default
	tlsEnabled := true
	
	if len(os.Args) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	// User arguments
	listenPort := flag.Int("-port", 8000, "Local port to listen for connections")
	tlsCert := flag.String("-cert", "", "Certificate file for TLS connection")
	preserveTls := flag.Bool("-no-tls", false, "Disable TLS comms. Example: -no-tls")
	flag.Parse()
	
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "-no-tls" {
			*tlsEnabled = preserveTls
		}
	})

	// Choose between secure or unsecure channel
	if *tlsEnabled == true && *tlsCert != "" {
		// Unsecured function
	} else {
		// Secured function
	}
	
	if *tlsCert

	mux := http.NewServeMux() 
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("It works !!\n"))
	})

	b, _ := ioutil.ReadFile(*tlsCert)
	var pemBlocks []*pem.Block
	var v *pem.Block
	var pkey []byte

	for {
		v, b = pem.Decode(b)
		if v == nil {
			break
		}
		if v.Type == "RSA PRIVATE KEY" {
			if x509.IsEncryptedPEMBlock(v) {
				pkey, _ = x509.DecryptPEMBlock(v, []byte("xxxxxxxxx"))
				pkey = pem.EncodeToMemory(&pem.Block{
					Type:  v.Type,
					Bytes: pkey,
				})
			} else {
				pkey = pem.EncodeToMemory(v)
			}
		} else {
			pemBlocks = append(pemBlocks, v)
		}
	}
	c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		Certificates: []tls.Certificate{c},
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", *listenPort),
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Fatal(srv.ListenAndServeTLS("", ""))
}
