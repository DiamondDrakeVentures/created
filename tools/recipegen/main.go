package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/diamonddrakeventures/created/tools/common/resourceloc"
	"github.com/diamonddrakeventures/created/tools/recipegen/generator"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var scanQ []string
	if len(os.Args) > 1 {
		scanQ = os.Args[1:]
	} else {
		scanQ = []string{"."}
	}

	workQ := make([]string, 0)

	for _, dir := range scanQ {
		if dir == "." {
			dir = cwd
		}

		err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("Error accessing %s: %v\n", path, err)
				return nil
			}

			if d.IsDir() || !generator.IsConfigFile(path) {
				return nil
			}

			workQ = append(workQ, path)
			return nil
		})
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", dir, err)
		}
	}

	errCnt := 0
	for _, configFile := range workQ {
		dir := filepath.Dir(configFile)

		config, err := generator.LoadConfig(configFile)
		if err != nil {
			fmt.Printf("Cannot read %q: %v\n", configFile, err)
			errCnt++
			continue
		}

		config.Template = filepath.FromSlash(config.Template)
		config.Source = filepath.FromSlash(config.Source)

		template, err := generator.LoadTemplate(filepath.Join(dir, config.Template))
		if err != nil {
			fmt.Printf("Cannot load template %q: %v\n", filepath.Join(dir, config.Template), err)
			errCnt++
			continue
		}

		src, err := generator.OpenSource(filepath.Join(dir, config.Source))
		if err != nil {
			fmt.Printf("Cannot open source %q: %v\n", filepath.Join(dir, config.Source), err)
			errCnt++
			continue
		}
		defer src.Close()

		for entry, err := src.Next(); !errors.Is(err, io.EOF); entry, err = src.Next() {
			generated := entry.ReplaceVars(template)
			output := filepath.FromSlash(config.Output)
			output = entry.ReplaceVars(output)
			output = resourceloc.SubPathToUnderscore(output)
			output = filepath.Join(dir, output)

			fmt.Printf("Generating %s ", output)

			outputDir := filepath.Dir(output)
			err = os.MkdirAll(outputDir, 0744)
			if err != nil && !errors.Is(err, os.ErrExist) {
				fmt.Printf("Error: %v\n", err)
				errCnt++
				continue
			}

			f, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				errCnt++
				continue
			}
			defer f.Close()
			_, err = f.WriteString(generated)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				errCnt++
				continue
			}

			fmt.Println("Done")
		}
	}

	if errCnt > 0 {
		fmt.Printf("Completed with %d error(s)\n", errCnt)
		os.Exit(1)
	} else {
		fmt.Println("Done!")
	}
}
