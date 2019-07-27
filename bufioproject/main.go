package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	f, err := os.Open("test")

	if err != nil {
		fmt.Printf("%s (test)\n", err)
	}

	wr, err := os.Create("answer.txt")

	if err != nil {
		fmt.Printf("%s (answer.txt)\n",err)
	}

	w := bufio.NewWriter(wr)
	r := bufio.NewReaderSize(f, 256)
	buf := make([]byte, 256)
	b, err := r.Read(buf)

	for err == nil {

		fmt.Printf("%s : %d, %q\n", time.Now() ,b, buf[:b])
		_, werr := w.Write([]byte(fmt.Sprintf("%s : %s\n", time.Now(), buf[:b])))

		if werr != nil {
			fmt.Printf("%s (write)\n", err)
		}

		time.Sleep(100)

		b, err = r.Read(buf)
	}

	if err == io.EOF {
		fmt.Printf("%s\n", err)
	}

	w.Flush()
	wr.Close()
	f.Close()
}
