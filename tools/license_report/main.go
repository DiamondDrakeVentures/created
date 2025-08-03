package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/diamonddrakeventures/created/tools/license_report/internal/modrinth"
	"github.com/diamonddrakeventures/created/tools/license_report/internal/packwiz"
)

func main() {
	fmt.Printf("%s v%s\n", AppName, Version)

	const modsDir = "mods"
	api := modrinth.New(modrinth.Config{
		UserAgent: UserAgent,
		Timeout:   15 * time.Second,
	})

	// mods := make([][]string, 0)
	// mods = append(mods, []string{
	// 	"Mod",
	// 	"License ID",
	// 	"License URL",
	// })

	queue := make([]packwiz.Mod, 0)
	err := filepath.Walk(modsDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var mod packwiz.Mod
		err = toml.Unmarshal(file, &mod)
		if err != nil {
			return err
		}

		queue = append(queue, mod)

		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d mods!\n", len(queue))

	f, err := os.Create("licenses.csv")
	if err != nil {
		fmt.Printf("Cannot encode licenses: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	err = w.Write([]string{
		"Mod",
		"Filename",
		"Modrinth ID",
		"License ID",
		"License URL",
	})
	if err != nil {
		fmt.Printf("Cannot write record: %v\n", err)
		os.Exit(1)
	}
	w.Flush()

	for i, mod := range queue {
		fmt.Printf("Fetching info for %s (%s) from Modrinth [%d/%d]\n",
			mod.Name, mod.Update.Modrinth.ModID, i, len(queue))

		project, err := api.Get(mod.Update.Modrinth.ModID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		err = w.Write([]string{
			mod.Name,
			mod.Filename,
			mod.Update.Modrinth.ModID,
			project.License.ID,
			project.License.URL,
		})
		if err != nil {
			fmt.Printf("Cannot write record: %v\n", err)
		}

		w.Flush()
		if err := w.Error(); err != nil {
			fmt.Printf("Cannot write record: %v\n", err)
		}
	}
}
