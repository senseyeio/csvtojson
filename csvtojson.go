package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	disableFirstRowHeader := flag.Bool("n", false, "Do not use the first row of each file as column names")
	out := os.Stdout

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		err := processReader(os.Stdin, out, !*disableFirstRowHeader)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error processin stdin: %v", err)
			os.Exit(1)
		}
		return
	}

	for _, r := range args {
		err := processFile(r, out, !*disableFirstRowHeader)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error processing file %s: %v", r, err)
			os.Exit(1)
		}
	}
}

func processFile(fn string, out io.Writer, firstRowHeader bool) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	return processReader(f, out, firstRowHeader)
}

func processReader(r io.Reader, out io.Writer, firstRowHeader bool) error {
	cr := csv.NewReader(r)

	var header []string
	var err error

	if firstRowHeader {
		header, err = cr.Read()
		if err != nil {
			return err
		}
		for i := range header {
			header[i] = strings.TrimSpace(header[i])
		}
	}

	fieldName := func(idx int) string {
		if len(header) <= idx {
			return fmt.Sprintf("_column%v", idx)
		}
		return header[idx]
	}

	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		tmp := make(map[string]interface{})
		for i, v := range row {
			tmp[fieldName(i)] = v
		}

		js, err := json.Marshal(tmp)
		if err != nil {
			line, _ := cr.FieldPos(0)
			return fmt.Errorf("line %v: %w", line, err)
		}

		_, _ = fmt.Fprintln(out, string(js))
	}
	return nil
}
