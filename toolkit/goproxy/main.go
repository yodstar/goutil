package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)

	p := flag.Int("p", 0, "the port listening on.")
	t := flag.String("t", "", "then ip:port to connect.")
	o := flag.String("o", "", "the output file path.")
	flag.Parse()

	if *p == 0 || *t == "" {
		flag.Usage()
		return
	}

	var file *os.File
	if *o != "" {
		var err error
		file, err = os.OpenFile(*o, os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	lsnr, err := net.Listen("tcp", fmt.Sprintf(":%d", *p))
	if err != nil {
		log.Println(err)
		return
	}
	defer lsnr.Close()
	fmt.Println("listening on " + lsnr.Addr().String())

	for {
		conn, err := lsnr.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()
			dial, err := net.Dial("tcp", *t)
			if err != nil {
				log.Println(err.Error())
				return
			}
			defer dial.Close()
			fmt.Println("connecting to " + dial.RemoteAddr().String())

			go func() {
				for {
					data := make([]byte, 102400)
					n, err := dial.Read(data)
					if err != nil {
						if err != io.EOF {
							log.Println(err.Error())
						}
						time.Sleep(time.Millisecond * 10)
						continue
					}
					s := fmt.Sprintf("%s <- %s %v\n", dial.LocalAddr(), dial.RemoteAddr(), data[0:n])
					if *o != "" {
						file.WriteString(s)
					} else {
						fmt.Print(s)
					}
					if _, err = conn.Write(data[0:n]); err != nil {
						log.Println(err.Error())
						return
					}
				}
			}()

			for {
				data := make([]byte, 102400)
				n, err := conn.Read(data)
				if err != nil {
					if err != io.EOF {
						log.Println(err.Error())
					}
					time.Sleep(time.Millisecond * 10)
					continue
				}

				s := fmt.Sprintf("%s -> %s %v\n", dial.LocalAddr(), dial.RemoteAddr(), data[0:n])
				if *o != "" {
					file.WriteString(s)
				} else {
					fmt.Print(s)
				}
				if _, err = dial.Write(data[0:n]); err != nil {
					log.Println(err.Error())
					return
				}
			}
		}(conn)
	}
}
