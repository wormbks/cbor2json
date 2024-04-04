package main

import (
	"flag"
	"io"
	"log"
	"os"
	"time"

	iot "github.com/wormbks/dry/ioutils"

	"github.com/wormbks/cbor2json/csd"
)

// main is the entry point of the program.
//
// No parameters.
// No return values.
func main() {
	inFile := flag.String("in", "<stdin>", "Input File (cbor Encoded)")
	outFile := flag.String("out", "<stdout>", "Output File to which decoded JSON will be written to (WILL overwrite if already present).")

	flag.Parse()

	csd.DecodeTimeZone, _ = time.LoadLocation("UTC")
	var in io.Reader = os.Stdin
	var out io.Writer = os.Stdout

	if *inFile != "<stdin>" {
		zin, err := iot.NewGzipReader(*inFile)
		if err != nil {
			log.Fatal(err)
		}
		in = zin
		defer zin.Close()
	}

	if *outFile != "<stdout>" {
		f, err := os.OpenFile(*outFile, os.O_RDWR|os.O_CREATE, 0o644)
		if err != nil {
			log.Fatal(err)
		}
		out = f
		defer f.Close()
	}

	err := csd.Cbor2JsonManyObjects(in, out)
	if err != nil {
		log.Fatal(err)
	}

}
