package processor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatRequestDate(t *testing.T) {
	date := time.Date(2020, 10, 15, 9, 20, 40, 0, time.UTC)
	formatedDate := formatRequestDate(date)
	assert.Equal(t, "09h20 - 15/10/2020", formatedDate)
}

func TestConvertFloatToString(t *testing.T) {
	value := float32(3.14)
	parsedValue := floatToString(value)
	assert.Equal(t, "3.14", parsedValue)
}
