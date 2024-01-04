package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"mower-project/internal/domain/model"
	"os"
)

// Execute program execution function that read the input file data, apply the logic and print the end results.
func Execute(config *viper.Viper) ([]*model.Mower, error) {
	log.Info("Execute Program")

	file, err := openInputFile(config)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	log.Info("Input file located and read properly")

	lawn, err := readLawnDimensions(file)

	if err != nil {
		return nil, err
	}

	// Create a scanner to read mower instructions
	scanner := bufio.NewScanner(file)

	var result []*model.Mower
	mowerCount := 0

	// Process each mower's instructions
	for scanner.Scan() {
		initialPosition := scanner.Text()
		scanner.Scan()
		instructions := scanner.Text()

		mowerCount++
		log.Debug("Mower", mowerCount, map[string]interface{}{
			"InitialPosition": initialPosition,
			"Instructions":    instructions,
		})

		// Process the mower and print its final position
		mower, err := ProcessMower(lawn, initialPosition, instructions)
		if err != nil {
			return nil, err
		}
		result = append(result, mower)
	}
	return result, nil
}

// readLawnDimensions read lawn dimensions from input file.
func readLawnDimensions(file *os.File) (model.Lawn, error) {
	// Read the first line to get the lawn dimensions
	var lawn model.Lawn
	fmt.Fscanf(file, "%d %d\n", &lawn.Width, &lawn.Height)

	if lawn.Width < 0 || lawn.Height < 0 {
		log.Error("Lawn Dimensions are wrong")
		return lawn, errors.New("WRONG LAWN DIMENSIONS")
	}
	log.Debug("lawn dimensions ", map[string]interface{}{
		"Width":  lawn.Width,
		"Height": lawn.Height,
	})
	return lawn, nil
}

// openInputFile open input file with path located in the config.
func openInputFile(config *viper.Viper) (*os.File, error) {
	// Open the input file
	pwd, err := os.Getwd()
	file, err := os.Open(pwd + config.GetString("filePath"))
	if err != nil {
		log.Error("Cannot read or locate input file")
		return nil, err
	}
	return file, nil
}

// ProcessMower process the instructions for a mower on the given lawn.
func ProcessMower(lawn model.Lawn, initialPosition string, instructions string) (*model.Mower, error) {
	mower := model.Mower{Orientation: initialPosition[4:5]} // Initialize the mower with the Orientation
	fmt.Sscanf(initialPosition, "%d %d", &mower.X, &mower.Y)

	if mower.X < 0 || mower.X > lawn.Width || mower.Y < 0 || mower.Y > lawn.Height {
		log.Error("Mower position is outside the lawn")
		return nil, errors.New("WRONG MOWER INIT POSITION")
	}
	for _, instruction := range instructions {
		mower.Move(string(instruction))
		// Check if the new position is within the lawn
		if mower.X < 0 || mower.X > lawn.Width || mower.Y < 0 || mower.Y > lawn.Height {
			// If outside the lawn, revert the move and break the loop
			mower.MoveBackward()
			continue
		}
	}

	return &mower, nil
}
