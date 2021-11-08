package lab2

import (
	"bytes"
	"fmt"
	"io"
)

type ComputeHandler struct {
	io.Reader
	io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(ch.Reader)
	if err != nil {
		return err
	}
	expression := buffer.String()
	res, err := CalculatePostfix(expression)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(ch.Writer, fmt.Sprint(res))
	return err
}
