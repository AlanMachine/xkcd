package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	modes := map[string]bool{"parse": true, "search": true}
	mode := flag.String("m", "", "Mode: parse | search")
	flag.Parse()

	if _, ok := modes[*mode]; !ok {
		log.Fatalln("Unknown mode")
	}

	switch *mode {
	case "parse":
		{
			parse()
		}
	case "search":
		{
			search(flag.Args())
		}
	default:
		{
			log.Fatalln("Unknown mode")
		}
	}
}

func parse() {
	var wc webComic

	attempt := 0
	for i := 1; attempt != 5; i++ {
		if dbCheck(i) {
			continue
		}

		data, err := request("https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			attempt++
			continue
		}

		if err := json.Unmarshal([]byte(data), &wc); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err = wc.write(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Printf("%d\t%s\n", wc.Num, wc.Title)
	}
}

func search(args []string) {
	query := strings.Join(args, " ")
	if len(query) == 0 {
		fmt.Fprintln(os.Stdout, "Empty search request")
		os.Exit(1)
	}

	result, err := dbSearch(query)
	if err != nil {
		log.Fatal(err)
	}

	if result == nil {
		fmt.Fprintln(os.Stdout, "WebComics not found")
		os.Exit(1)
	}

	var wc webComic
	var webComics []webComic
	for _, num := range result {
		wc.read(num)
		webComics = append(webComics, wc)
	}

	if err := formatOutput(webComics, 0); err != nil {
		log.Fatalln(err)
	}
}
