package main

import (
	"bufio"
	"fmt"
	"os"

	"keesvv.nl/morse/decoder"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	decoder := decoder.NewDecoder(decoder.InternationalMorse)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if dec := decoder.Decode(scanner.Bytes()); dec != 0 {
			fmt.Print(string(dec))
		}
	}
}
