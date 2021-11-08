package lab2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCompute_Correct_1(t *testing.T) {
	output := new(bytes.Buffer)
	handler := ComputeHandler{
		Input:  strings.NewReader("2 6 +"),
		Output: output,
	}
	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "8", output.String())
	}
}

func TestCompute_Correct_2(t *testing.T) {
	output := new(bytes.Buffer)
	handler := ComputeHandler{
		Input:  strings.NewReader("4 1 - 5 *"),
		Output: output,
	}
	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "15", output.String())
	}
}

func TestCompute_Correct_3(t *testing.T) {
	output := new(bytes.Buffer)
	handler := ComputeHandler{
		Input:  strings.NewReader("20 4 / 5 -"),
		Output: output,
	}
	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "0", output.String())
	}
}

func TestCompute_Incorrect_3(t *testing.T) {
	output := new(bytes.Buffer)
	handler := ComputeHandler{
		Input:  strings.NewReader("string"),
		Output: output,
	}
	err := handler.Compute()
	assert.Error(t, err)
}
