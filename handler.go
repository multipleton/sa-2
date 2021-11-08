package lab2

import (
	"io"
)

type ComputeHandler struct {
	io.Reader
	io.Writer
}

func (ch *ComputeHandler) Compute() error {
	input := make([]byte, 64)
	for {
		_, err := ch.Reader.Read(input)
		if err == io.EOF {
			break
		}
	}
	expression := string(input)
	_, err := CalculatePostfix(expression)
	return err
}
