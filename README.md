# CSD (CBOR Stream Decoder)

This updated version CBOR Stream Decoder with float16 support.

CSD decodes a stream of CBOR bytes into JSON. CBOR stream could be copressed by gzip.
``` sh
Usage:

    csd [-in inputFile] [-out outputFile] [-follow]

Run `csd -h` for a list of supported options and usage.

If `-in` is omitted, csd reads from stdin.

If `-out` is omitted, csd writes to stdout.

```

## Example

Suppose CBOR encoded data is present in file cbor.log, you could do one of 
these methods to convert to JSON using csd:

Input from file, output to stdout :

    $ csd -in cbor.log
    {"level":"error","Fault":41650,"time":"2018-03-31T07:17:19-07:00","message":"TCA:"}
    {"level":"error","Fault":41654,"time":"2018-03-31T07:17:19-07:00","message":"TCA:"}
    ...


Input from stdin, output to stdout

    $ cat cbor.log | csd 
    {"level":"error","Fault":41650,"time":"2018-03-31T07:17:19-07:00","message":"TCA:"}
    {"level":"error","Fault":41654,"time":"2018-03-31T07:17:19-07:00","message":"TCA:"}
    ...


Input from file, output to file

    $ csd -in cbor.log -out json.txt
    $ cat json.txt
    {"level":"error","Fault":41650,"time":"2018-03-31T07:17:19-07:00","message":"TCA:"}
    {"level":"error","Fault":41654,"time":"2018-03-31T07:17:19-07:00","message":"TCA:"}
    ...



## APIs

For documentation of APIs used to decode, see: https://godoc.org/github.com/toravir/csd/libs/

## Limitations

The input is expected to be CBOR data (either gzib-compressed or not). It is NOT possible to
detect JSON (text) output reliably since binary format spans over JSON character set also.

