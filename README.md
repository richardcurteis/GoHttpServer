1. Python script to:
1.1 Compile static binary server for target architecture
1.2 Remove debug info from binary and pack as small as possible

2. Simple Golang HTTP Server # https://gist.github.com/paulmach/7271283

3. Project to create slightly more compl,ex server utilising TLS encryption. WIP.

4. goCat - Netcat copy but in Go. basic server to listen on a port and execute commands

Fix:

1. Compilation Helper not creating 64 bit for windows. All coming out as 32 bit
