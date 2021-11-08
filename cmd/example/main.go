package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/multipleton/sa-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File where to write the result")
)

func main() {
	flag.Parse()
	var reader io.Reader
	var writer io.Writer
	err := handleIO(&reader, &writer)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	handler := &lab2.ComputeHandler{
		Reader: reader,
		Writer: writer,
	}
	err = handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func handleIO(input *io.Reader, output *io.Writer) error {
	if *inputExpression != "" && *inputFile != "" || *inputExpression == "" && *inputFile == "" {
		return errors.New("invalid arguments")
	}
	if *inputExpression != "" {
		*input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		res, err := os.Open(*inputFile)
		if err != nil {
			return err
		}
		*input = res
	}
	if *outputFile != "" {
		res, err := os.OpenFile(*outputFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
		*output = res
	} else {
		*output = os.Stdout
	}
	return nil
}
