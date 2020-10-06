package shortshare

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"os"
)

func main() {
	unshare := flag.Bool("u", false, "If set will unshare stdin")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	bin := new(bytes.Buffer)
	io.Copy(bin, reader)

	if *unshare {
		var s string
		err := Unshare(bin, &s)
		if err != nil {
			panic(err)
		}
		writer.WriteString(s)
	} else {
		b, err := Share(bin.String())
		if err != nil {
			panic(err)
		}
		io.Copy(writer, b)
	}
	writer.Flush()
}
