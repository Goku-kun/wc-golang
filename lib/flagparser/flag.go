package flagparser

import "flag"

func ParseFlags() (map[string]bool, []string) {
	c := flag.Bool("c", false, "Outputs the bytes of the file")
	l := flag.Bool("l", false, "Outputs the number of lines in the file")
	w := flag.Bool("w", false, "Outputs the number of words in the file")

	flag.Parse()

	args := flag.Args()

	return map[string]bool{
		"c": *c,
		"l": *l,
		"w": *w,
	}, args
}

