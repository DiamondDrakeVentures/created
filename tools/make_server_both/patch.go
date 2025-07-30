package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

var sidePattern = regexp.MustCompile(`(?m)^side\s*=\s*"(both|server|client)"$`)

func main() {
	const modsDir = "mods"
	err := filepath.Walk(modsDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		mod, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// sides := sidePattern.FindStringSubmatch(string(mod))
		// if len(sides) < 1 {
		// 	fmt.Printf("%s missing side!\n", path)
		// 	return nil
		// }
		// side := sides[0]
		// side := sidePattern.ReplaceAllString(string(mod), "$1")
		// fmt.Printf("%s is %q\n", path, side)
		sides := sidePattern.FindStringSubmatch(string(mod))
		if len(sides) < 2 {
			fmt.Printf("%s missing side!\n", path)
			return nil
		}
		side := sides[1]

		rewrite := false

		if side == "server" {
			newSide := "both"
			fmt.Printf("Rewriting %s side from %q to %q\n", path, side, newSide)
			mod = sidePattern.ReplaceAll(mod, []byte(`side = "`+newSide+`"`))
			rewrite = true
		}

		if rewrite {
			err = os.WriteFile(path, mod, info.Mode())
		}

		return err
	})
	if err != nil {
		panic(err)
	}
}
