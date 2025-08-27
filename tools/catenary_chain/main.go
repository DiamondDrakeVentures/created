package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const template = `{
  "textures": {
    "chain": "{chain}",
    "knot": "{knot}"
  }
}
`

const (
	resourcepackBaseDir = "resourcepacks"
	datapackBaseDir     = "datapacks"
	resourcepackID      = "created"
	datapackID          = "created"
)

var blacklist = []string{
	"minecraft:chain",
}

type Tag struct {
	Replace bool     `json:"replace"`
	Values  []string `json:"values"`
}

func main() {
	f, err := os.Open(datapackPath("created:tags/item/chains.json"))
	if err != nil {
		panic(err)
	}

	var tag Tag
	err = json.NewDecoder(f).Decode(&tag)
	if err != nil {
		panic(err)
	}
	f.Close()

	errCnt := 0
	for _, val := range tag.Values {
		if slices.Contains(blacklist, val) {
			fmt.Printf("Skipping %q: in blacklist\n", val)
			continue
		}

		namespace, resourcePath := resourcePath(val)
		resourceID := fmt.Sprintf("%s:models/entity/connectiblechains/%s.json", namespace, resourcePath)
		_, err = os.Stat(resourcepackPath(resourceID))
		if err == nil {
			fmt.Printf("Skipping %q: file already exist\n", val)
			continue
		} else if err != nil && !errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Error with %q: %v\n", val, err)
			errCnt++
			continue
		}

		fmt.Printf("Creating file %s\n", resourcepackPath(resourceID))
		f, err := os.OpenFile(resourcepackPath(resourceID), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 10644)
		if err != nil {
			fmt.Printf("Error with %q: %v\n", val, err)
			errCnt++
			continue
		}
		defer f.Close()

		chainPath := fmt.Sprintf("%s:block/materials/%s", namespace, resourcePath)
		knotPath := chainPath

		model := strings.ReplaceAll(template, "{chain}", chainPath)
		model = strings.ReplaceAll(model, "{knot}", knotPath)

		_, err = f.WriteString(model)
		if err != nil {
			fmt.Printf("Error with %q: %v\n", val, err)
			errCnt++
			continue
		}
	}

	if errCnt > 0 {
		fmt.Println("Completed with errors!")
		os.Exit(1)
	} else {
		fmt.Println("Done!")
	}
}

func resourcepackPath(resourceID string) string {
	namespace, resourcePath := resourcePath(resourceID)

	return filepath.Join(resourcepackBaseDir, resourcepackID, "assets", namespace, resourcePath)
}

func datapackPath(resourceID string) string {
	namespace, resourcePath := resourcePath(resourceID)

	return filepath.Join(datapackBaseDir, datapackID, "data", namespace, resourcePath)
}

func resourcePath(resourceID string) (namespace string, path string) {
	if strings.Contains(resourceID, ":") {
		fragments := strings.SplitN(resourceID, ":", 2)
		namespace = fragments[0]
		path = fragments[1]
	} else {
		namespace = "minecraft"
		path = resourceID
	}
	return
}
