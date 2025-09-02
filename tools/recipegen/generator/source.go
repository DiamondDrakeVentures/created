package generator

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

type Source struct {
	KeyDict []string

	src       *os.File
	csvReader *csv.Reader
}

func (src *Source) read() (record []string, err error) {
	return src.csvReader.Read()
}

func (src *Source) loadHeaders() error {
	rec, err := src.read()
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	src.KeyDict = make([]string, 0)
	for _, head := range rec {
		head = strings.TrimSpace(head)
		if head != "" {
			src.KeyDict = append(src.KeyDict, head)
		}
	}

	return nil
}

func (src *Source) Next() (entry Entry, err error) {
	rec, err := src.read()
	if err != nil {
		return
	}

	entry = make(Entry)
	for i, c := range rec {
		entry[src.KeyDict[i]] = strings.TrimSpace(c)
	}
	return
}

func (src *Source) Close() error {
	err := src.src.Close()
	src.src = nil
	src.csvReader = nil
	return err
}

func OpenSource(path string) (src *Source, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	src = new(Source)
	src.src = f
	src.csvReader = csv.NewReader(f)

	err = src.loadHeaders()
	return
}
