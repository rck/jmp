package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/rck/jmp/jumpdb"
)

const jmpdb = ".jmpdb"

func printResult(entries []jumpdb.DBEntry) {
	if len(entries) == 0 {
		fmt.Println(".")
		return
	}

	if *flagC {
		for _, e := range entries {
			fmt.Println(e.Path)
		}
		return
	}

	// if there is an exact match, prefer that one (e.g., from completion)
	if flag.NArg() == 1 {
		input := flag.Args()[0]
		for _, e := range entries {
			if e.Path == input {
				fmt.Println(e.Path)
				return
			}
		}
	}
	fmt.Println(entries[0].Path)
}

var (
	flagS = flag.Int("s", 0, "Set weight of current path (0 is ignored, < 0 delete from DB)")
	flagC = flag.Bool("c", false, "Complete paths")
	flagL = flag.Bool("l", false, "List DB")
)

// TODO(rck): This code is a lot uglier than I would like to have it...
func main() {
	flag.Parse()
	bestPath := []jumpdb.DBEntry{{Path: ".", Weight: 0}}

	var complete bool

	if flag.NFlag() == 0 && flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "To add a path:\n\tcd PATH\n\tj .\n")
		printResult(bestPath)
		os.Exit(1)
	}
	if *flagC && flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Specify at least one argument")
		printResult(bestPath)
		os.Exit(1)
	}
	if (*flagC && flag.NArg() > 0) || (flag.NFlag() == 0 && flag.Args()[0] != ".") {
		complete = true
	}

	dbPath := path.Join(os.Getenv("HOME"), jmpdb)
	db := jumpdb.NewDB()
	dbloaded := db.LoadAsync(dbPath)

	var regex *regexp.Regexp
	/* we can prepare the regex before loading is finished */
	if complete {
		if flag.Args()[0] != "." {
			reg := ".*" + strings.Join(flag.Args(), ".*") + ".*"
			rgex, err := regexp.Compile(reg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not compile %s", reg)
				printResult(bestPath)
				os.Exit(1)
			}
			regex = rgex
		}
		if regex == nil {
			fmt.Fprintf(os.Stderr, "Something is wrong with your regex\n")
			printResult(bestPath)
			os.Exit(1)
		}
	}
	if err := <-dbloaded; err != nil {
		log.Fatal(err)
	}

	saveDB := func() {
		if err := db.Save(dbPath); err != nil {
			fmt.Fprintf(os.Stderr, "Could not save DB: %v\n", err)
		}
	}

	addCwdSavePrint := func(weight int64) {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get PWD: %v\n", err)
			return
		}
		if err := db.SetEntry(wd, weight); err != nil {
			fmt.Fprintf(os.Stderr, "Could not call SetEntry: %v\n", err)
			return
		}
		saveDB()
		printResult(bestPath)
	}

	if complete {
		entries := db.Complete(regex)
		if len(entries) > 0 && !*flagC {
			if err := db.IncEntry(entries[0].Path); err != nil {
				fmt.Fprintf(os.Stderr, "Could not increase weight: %v\n", err)
			}
			saveDB()
		}
		printResult(entries)
	} else if *flagL {
		db.List()
	} else if *flagS != 0 {
		addCwdSavePrint(int64(*flagS))
	} else if flag.NArg() == 1 && flag.Args()[0] == "." {
		addCwdSavePrint(1)
	}
}
