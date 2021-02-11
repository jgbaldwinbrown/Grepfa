package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "regexp"
)

func search_and_print_entry(header string, seq string, qual string, search_token *regexp.Regexp) {
    submatch_strings := search_token.FindAllStringSubmatch(seq, -1)
    submatch_indices := search_token.FindAllStringSubmatchIndex(seq, -1)
    for i := 0; i < len(submatch_strings); i++ {
        fmt.Printf("%s", header[1:])
        for j := 0; j < len(submatch_indices[i]); j+=2 {
            fmt.Printf("\t%d\t%d\t%s", submatch_indices[i][j], submatch_indices[i][j+1], submatch_strings[i][j/2])
        }
        fmt.Printf("\n")
    }
}

func main() {
    var header = ""
    var seq = ""
    var qual = ""
    var line_counter = 0
    search_token, _ := regexp.Compile(os.Args[1])
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        switch line_counter % 4 {
        case 0:
            header = scanner.Text()
        case 1:
            seq = scanner.Text()
        case 3:
            qual = scanner.Text()
            search_and_print_entry(header, seq, qual, search_token)
        }
        line_counter++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
