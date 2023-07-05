package whmcs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	assert := assert.New(t)

	inputTime := time.Date(2023, time.July, 5, 16, 21, 0, 0, time.UTC)
	expectedFormattedTime := formattedTime("2023-07-05 16:21:00")

	formatted := Time(inputTime)
	assert.NotNil(formatted)
	assert.Equal(expectedFormattedTime, *formatted)

	parsedTime := formatted.Parse()
	assert.Equal(inputTime, parsedTime)
}
func TestBoolInt(t *testing.T) {
	assert := assert.New(t)

	inputBool := true
	expectedBoolInt := boolean(1)

	boolInt := BoolInt(inputBool)
	assert.NotNil(boolInt)
	assert.Equal(expectedBoolInt, *boolInt)

	boolValue := boolInt.Bool()
	assert.Equal(inputBool, boolValue)

	inputBool = false
	expectedBoolInt = boolean(0)

	boolInt = BoolInt(inputBool)
	assert.NotNil(boolInt)
	assert.Equal(expectedBoolInt, *boolInt)

	boolValue = boolInt.Bool()
	assert.Equal(inputBool, boolValue)
}

func TestBool(t *testing.T) {
	assert := assert.New(t)

	inputBool := true

	boolPtr := Bool(inputBool)
	assert.NotNil(boolPtr)
	assert.Equal(inputBool, *boolPtr)

	inputBool = false

	boolPtr = Bool(inputBool)
	assert.NotNil(boolPtr)
	assert.Equal(inputBool, *boolPtr)
}

func TestString(t *testing.T) {
	assert := assert.New(t)

	inputString := "Hello, World!"

	stringPtr := String(inputString)
	assert.NotNil(stringPtr)
	assert.Equal(inputString, *stringPtr)

	inputString = ""

	stringPtr = String(inputString)
	assert.NotNil(stringPtr)
	assert.Equal(inputString, *stringPtr)
}

func TestInt(t *testing.T) {
	assert := assert.New(t)

	inputInt := 21

	intPtr := Int(inputInt)
	assert.NotNil(intPtr)
	assert.Equal(inputInt, *intPtr)

	inputInt = -21

	intPtr = Int(inputInt)
	assert.NotNil(intPtr)
	assert.Equal(inputInt, *intPtr)

	inputInt = 0

	intPtr = Int(inputInt)
	assert.NotNil(intPtr)
	assert.Equal(inputInt, *intPtr)
}
