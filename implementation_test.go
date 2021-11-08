package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePostfix_EmptyParameter(t *testing.T) {
	_, err := CalculatePostfix("")
	assert.Error(t, err, "Calling CalculatePostfix with empty string should return error")
}

func TestCalculatePostfixn_InvalidParameters(t *testing.T) {
	_, err := CalculatePostfix("abc")
	assert.Error(t, err, "Calling CalculatePostfix with invalid characters should return error")
}

func TestCalculatePostfix_Adding_Integers(t *testing.T) {
	res, err := CalculatePostfix("4 7 + 6 +")
	if assert.Nil(t, err) {
		assert.Equal(t, 17., res)
	}
}

func TestCalculatePostfix_Adding_Floats(t *testing.T) {
	res, err := CalculatePostfix("2.5 5.1 +")
	if assert.Nil(t, err) {
		assert.Equal(t, 7.6, res)
	}
}

func TestCalculatePostfix_Adding_Negatives(t *testing.T) {
	res, err := CalculatePostfix("-3 -11 +")
	if assert.Nil(t, err) {
		assert.Equal(t, -14., res)
	}
}

func TestCalculatePostfix_Subtracting_Integers(t *testing.T) {
	res, err := CalculatePostfix("5 18 -")
	if assert.Nil(t, err) {
		assert.Equal(t, -13., res)
	}
}

func TestCalculatePostfix_Subtracting_Floats(t *testing.T) {
	res, err := CalculatePostfix("3.4 5.1 - 2.6 -")
	if assert.Nil(t, err) {
		assert.Equal(t, -4.3, res)
	}
}

func TestCalculatePostfix_Subtracting_Negatives(t *testing.T) {
	res, err := CalculatePostfix("-6 -3 - -5 -")
	if assert.Nil(t, err) {
		assert.Equal(t, 2., res)
	}
}

func TestCalculatePostfix_Multiplying_Integers(t *testing.T) {
	res, err := CalculatePostfix("5 18 * 2 *")
	if assert.Nil(t, err) {
		assert.Equal(t, 180., res)
	}
}

func TestCalculatePostfix_Multiplying_Floats(t *testing.T) {
	res, err := CalculatePostfix("6.25 5.8 *")
	if assert.Nil(t, err) {
		assert.Equal(t, 36.25, res)
	}
}

func TestCalculatePostfix_Multiplying_Negatives(t *testing.T) {
	res, err := CalculatePostfix("-9 11 *")
	if assert.Nil(t, err) {
		assert.Equal(t, -99., res)
	}
}

func TestCalculatePostfix_Multiplying_ByZero(t *testing.T) {
	res, err := CalculatePostfix("13 0 *")
	if assert.Nil(t, err) {
		assert.Equal(t, 0., res)
	}
}

func TestCalculatePostfix_Dividing_Integers(t *testing.T) {
	res, err := CalculatePostfix("18 5 / 2 /")
	if assert.Nil(t, err) {
		assert.Equal(t, 1.8, res)
	}
}

func TestCalculatePostfix_Dividing_Floats(t *testing.T) {
	res, err := CalculatePostfix("9.5 8.0 /")
	if assert.Nil(t, err) {
		assert.Equal(t, 1.1875, res)
	}
}

func TestCalculatePostfix_Dividing_Negatives(t *testing.T) {
	res, err := CalculatePostfix("-45 -15 / -3 /")
	if assert.Nil(t, err) {
		assert.Equal(t, -1., res)
	}
}

func TestCalculatePostfix_Dividing_ByZero(t *testing.T) {
	_, err := CalculatePostfix("13 0 /")
	assert.Error(t, err, "String with dividing by zero should return error")
}

func TestCalculatePostfix_Exponentiating_IntegerExponent(t *testing.T) {
	res, err := CalculatePostfix("6 3 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, 216., res)
	}
}

func TestCalculatePostfix_Exponentiating_FloatExponent(t *testing.T) {
	res, err := CalculatePostfix("0.64 0.5 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, 0.8, res)
	}
}

func TestCalculatePostfix_Exponentiating_NegativeExponent(t *testing.T) {
	res, err := CalculatePostfix("2 -5 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, 0.03125, res)
	}
}

func TestCalculatePostfix_Exponentiating_ZeroExponent(t *testing.T) {
	res, err := CalculatePostfix("14 0 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, 1., res)
	}
}

func TestCalculatePostfix_ComplexExpression_1(t *testing.T) {
	res, err := CalculatePostfix("6 10 + 4 - 1 1 2 * + / 1 +")
	if assert.Nil(t, err) {
		assert.Equal(t, 5., res)
	}
}

func TestCalculatePostfix_ComplexExpression_2(t *testing.T) {
	_, err := CalculatePostfix("3 7 - 6 + 2 2 - /")
	assert.Error(t, err, "String with dividing by zero should return error")
}

func TestCalculatePostfix_ComplexExpression_3(t *testing.T) {
	res, err := CalculatePostfix("5 8 + 3 - 2 ^ 5 2 * /")
	if assert.Nil(t, err) {
		assert.Equal(t, 10., res)
	}
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("7 5 - 8 + -1 -5 * /")
	fmt.Println(res)
	// Output:
	// 2
}
