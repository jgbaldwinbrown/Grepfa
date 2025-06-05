package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/jgbaldwinbrown/fastats/pkg"
)

func search_and_print_entry(entry fastats.FaEntry, search_token *regexp.Regexp) {
	submatch_strings := search_token.FindAllStringSubmatch(entry.Seq, -1)
	submatch_indices := search_token.FindAllStringSubmatchIndex(entry.Seq, -1)
	for i := 0; i < len(submatch_strings); i++ {
		if _, e := fmt.Printf("%s", entry.Header[1:]); e != nil {
			log.Fatal(e)
		}
		for j := 0; j < len(submatch_indices[i]); j+=2 {
			if _, e := fmt.Printf("\t%d\t%d\t%s", submatch_indices[i][j], submatch_indices[i][j+1], submatch_strings[i][j/2]); e != nil {
				log.Fatal(e)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	search_token, e := regexp.Compile(os.Args[1])
	if e != nil {
		log.Fatal(e)
	}
	for entry, e := range fastats.ParseFasta(os.Stdin) {
		if e != nil {
			log.Fatal(e)
		}
		search_and_print_entry(entry, search_token)
	}
}
