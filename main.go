package main

import (
	"bufio"
	"flag"
	"fmt"
	bin "github.com/weregoat/amazingoat/simulations/binary"
	j "github.com/weregoat/amazingoat/simulations/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var asJSON = flag.Bool("json", false, "STDIN stream is to be processed as JSON")
	var debug = flag.Bool("debug", false, "prints logs about the moving cursor")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	// Used to verify results. To pipe a binary stream is a bit a pain.
	//bs := []byte{ 0x4, 0x0, 0x4, 0x0, 0x2, 0x0, 0x2, 0x0 , 0x1, 0x4, 0x1, 0x3, 0x2, 0x3, 0x2, 0x4, 0x1, 0x0}
	//bs = []byte(`{"Width":4,"Height":4, "X":2, "Y":2, "Commands": ["f", "l", "f", "r", "b", "r", "b", "l", "f", "q"]}`)
	//*asJSON = true
	//*debug = true
	if *asJSON {
		s, err := j.New(bs)
		if err != nil {
			log.Fatal(err)
		}
		s.Debug(*debug)
		output, err := s.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stdout, string(output))
	} else { // Defaults to binary
		s, err := bin.New(bs)
		if err != nil {
			log.Fatal(err)
		}
		s.Debug(*debug)
		stdout := io.Writer(os.Stdout)
		output, err := s.Run()
		if err != nil {
			log.Fatal(err)
		}
		stdout.Write(output)
	}
}
