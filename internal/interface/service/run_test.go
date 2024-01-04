package service_test

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"mower-project/internal/domain/model"
	"mower-project/internal/interface/service"
	"os"
	"strings"
	"testing"
)

const tempFileName = "test_input.txt"

// TestExecute main test function to test the temp file creation with data allocation, config reading, the execution logic and the temp file deletion
func TestExecute(t *testing.T) {
	// Create a temporary file with test data
	inputData := "5 5\n1 2 N\nLFLFLFLFF\n3 3 E\nFFRFFRFRRF\n"
	tempFile := createTempFile(tempFileName, inputData)
	defer tempFile.Close()

	// Set up the Viper configuration
	config := viper.New()
	config.Set("filePath", "/"+tempFileName)

	// Call the Execute function with the test configuration
	mowers := service.Execute(config)

	// Build the result assert comparison object
	var result strings.Builder
	for _, mower := range mowers {
		result.WriteString(fmt.Sprintf("%d %d %s\n", mower.X, mower.Y, mower.Orientation))
	}

	// Expected output
	expectedOutput := "1 3 N\n5 1 E\n"

	// Exception on results compare
	assert.Equal(t, result.String(), expectedOutput)

	// Delete temp file
	deleteTempFile(tempFileName)
}

// createTempFile helper function to create a temporary file for testing
func createTempFile(filename, content string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic("Error creating temporary file: " + err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic("Error writing to temporary file: " + err.Error())
	}

	return file
}

// deleteTempFile helper function to delete the temporary testing file
func deleteTempFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		panic("Error deleting temporary file: " + err.Error())
	}
}

// TestProcessMower functional testing
func TestProcessMower(t *testing.T) {
	lawn := model.Lawn{Width: 5, Height: 5}
	initialPosition := "2 2 N"
	instructions := "LFFRRFFLF"
	result := service.ProcessMower(lawn, initialPosition, instructions)
	expectedResult := model.Mower{
		Orientation: "N",
		X:           2,
		Y:           3,
	}
	if !assert.Equal(t, result, expectedResult) {
		t.Errorf("ProcessMower returned %v, expected %v", result, expectedResult)
	}

	initialPosition = "1 1 N"
	instructions = "FFLFFRF"
	result = service.ProcessMower(lawn, initialPosition, instructions)
	expectedResult = model.Mower{
		Orientation: "N",
		X:           0,
		Y:           4,
	}
	if !assert.Equal(t, result, expectedResult) {
		t.Errorf("ProcessMower returned %v, expected %v", result, expectedResult)
	}

	initialPosition = "1 1 N"
	instructions = "FFLFFLF"
	result = service.ProcessMower(lawn, initialPosition, instructions)
	expectedResult = model.Mower{
		Orientation: "S",
		X:           0,
		Y:           2,
	}
	if !assert.Equal(t, result, expectedResult) {
		t.Errorf("ProcessMower returned %v, expected %v", result, expectedResult)
	}

	initialPosition = "1 1 N"
	instructions = "FFLFFLLF"
	result = service.ProcessMower(lawn, initialPosition, instructions)
	expectedResult = model.Mower{
		Orientation: "E",
		X:           1,
		Y:           3,
	}
	if !assert.Equal(t, result, expectedResult) {
		t.Errorf("ProcessMower returned %v, expected %v", result, expectedResult)
	}

	initialPosition = "3 3 E"
	instructions = "FFLFFRFFLL"
	result = service.ProcessMower(lawn, initialPosition, instructions)
	expectedResult = model.Mower{
		Orientation: "W",
		X:           5,
		Y:           5,
	}
	if !assert.Equal(t, result, expectedResult) {
		t.Errorf("ProcessMower returned %v, expected %v", result, expectedResult)
	}
}
