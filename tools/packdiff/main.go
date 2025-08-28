package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/diamonddrakeventures/created/tools/packdiff/crashassistant"
	"github.com/diamonddrakeventures/created/tools/packdiff/list"
)

const flagsFmt = "  %s        %s\n"

func main() {
	var (
		showRemovals  = false
		showAdditions = false
		showUpdates   = false
		showUnchanged = false
		showStats     = true
	)
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: packdiff [options...] <a> <b>")
		fmt.Println()
		fmt.Printf(flagsFmt, "-r",
			"Show removal(s).")
		fmt.Printf(flagsFmt, "-a",
			"Show addition(s).")
		fmt.Printf(flagsFmt, "-u",
			"Show update(s).")
		fmt.Printf(flagsFmt, "-n",
			"Show unchanged.")
		fmt.Printf(flagsFmt, "-q",
			"Hide statistics.")
		os.Exit(1)
	}

	var aFile, bFile string

	for i, arg := range args {
		if i == len(args)-2 {
			aFile = arg
		} else if i == len(args)-1 {
			bFile = arg
		} else {
			switch arg {
			case "-r":
				showRemovals = true

			case "-a":
				showAdditions = true

			case "-u":
				showUpdates = true

			case "-n":
				showUnchanged = true

			case "-q":
				showStats = false

			default:
				fmt.Printf("Unknown argument: %s\n", arg)
				os.Exit(1)
			}
		}
	}

	var err error

	listA := list.NewModList()
	err = parseCAModList(aFile, listA)
	if err != nil {
		panic(err)
	}

	listB := list.NewModList()
	err = parseCAModList(bFile, listB)
	if err != nil {
		panic(err)
	}

	diff := listA.Diff(listB)

	if !showRemovals && !showAdditions && !showUpdates && !showUnchanged {
		diff.Results().Print()
	}
	if showRemovals {
		diff.Rems().Print()
	}
	if showAdditions {
		diff.Adds().Print()
	}
	if showUpdates {
		diff.Updates().Print()
	}
	if showUnchanged {
		diff.Unchanged().Print()
	}

	if showStats {
		diff.PrintStats()
	}
}

func parseCAModList(modlistFile string, modList *list.ModList) (err error) {
	f, err := os.Open(modlistFile)
	if err != nil {
		return
	}
	defer f.Close()

	var caModList crashassistant.ModList
	err = json.NewDecoder(f).Decode(&caModList)
	if err != nil {
		return
	}

	for _, mod := range caModList {
		modList.AddFromCrashAssistant(mod)
	}

	return
}
