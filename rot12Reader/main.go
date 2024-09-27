package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var lookup = map[byte]byte{
	'A': 'N',
	'B': 'O',
	'C': 'P',
	'D': 'Q',
	'E': 'R',
	'F': 'S',
	'G': 'T',
	'H': 'U',
	'I': 'V',
	'J': 'W',
	'K': 'X',
	'L': 'Y',
	'M': 'Z',
	'N': 'A',
	'O': 'B',
	'P': 'C',
	'Q': 'D',
	'R': 'E',
	'S': 'F',
	'T': 'G',
	'U': 'H',
	'V': 'I',
	'W': 'J',
	'X': 'K',
	'Y': 'L',
	'Z': 'M',
	'a': 'n',
	'b': 'o',
	'c': 'p',
	'd': 'q',
	'e': 'r',
	'f': 's',
	'g': 't',
	'h': 'u',
	'i': 'v',
	'j': 'w',
	'k': 'x',
	'l': 'y',
	'm': 'z',
	'n': 'a',
	'o': 'b',
	'p': 'c',
	'q': 'd',
	'r': 'e',
	's': 'f',
	't': 'g',
	'u': 'h',
	'v': 'i',
	'w': 'j',
	'x': 'k',
	'y': 'l',
	'z': 'm',
}

// https://go.dev/tour/methods/23
func (reader rot13Reader) Read(out []byte) (int, error) {

	n, err := reader.r.Read(out)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		newValue, exists := lookup[out[i]]
		if exists {
			// Ignore spaces, and punctuation symbols.
			out[i] = newValue
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	// s2 := strings.NewReader("You cracked the code!")
	// rInv := rot13Reader{s2}


	// Manually read everything.
	// buffer := make([]byte, 8)
	// var n int;
	// var err error;
	// n, err = r.Read(buffer)
	// n, err = r.Read(buffer)
	// n, err = r.Read(buffer)
	// fmt.Println(n, err)
	// fmt.Printf("'%v'", string(buffer))

	io.Copy(os.Stdout, &r)
}
