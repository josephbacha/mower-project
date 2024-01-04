package service

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"mower-project/internal/domain/model"
	"os"
)

// Execute program execution function that read the input file data, apply the logic and print the end results.
func Execute(config *viper.Viper) []model.Mower {
	log.Info("Execute Program")

	file, mowers, done := openInputFile(config)
	defer file.Close()

	if done {
		return mowers
	}

	log.Info("Input file located and read properly")

	lawn := readLawnDimensions(file)

	// Create a scanner to read mower instructions
	scanner := bufio.NewScanner(file)

	var result []model.Mower
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
		result = append(result, ProcessMower(lawn, initialPosition, instructions))
	}
	return result
}

// readLawnDimensions read lawn dimensions from input file.
func readLawnDimensions(file *os.File) model.Lawn {
	// Read the first line to get the lawn dimensions
	var lawn model.Lawn
	fmt.Fscanf(file, "%d %d\n", &lawn.Width, &lawn.Height)

	log.Debug("lawn dimensions ", map[string]interface{}{
		"Width":  lawn.Width,
		"Height": lawn.Height,
	})
	return lawn
}

// openInputFile open input file with path located in the config.
func openInputFile(config *viper.Viper) (*os.File, []model.Mower, bool) {
	// Open the input file
	pwd, err := os.Getwd()
	file, err := os.Open(pwd + config.GetString("filePath"))
	if err != nil {
		log.Error("Cannot read or locate input file")
		panic("Error opening file:" + err.Error())
		return nil, nil, true
	}
	return file, nil, false
}

// ProcessMower process the instructions for a mower on the given lawn.
func ProcessMower(lawn model.Lawn, initialPosition string, instructions string) model.Mower {
	mower := model.Mower{Orientation: initialPosition[4:5]} // Initialize the mower with the Orientation
	fmt.Sscanf(initialPosition, "%d %d", &mower.X, &mower.Y)

	for _, instruction := range instructions {
		mower.Move(string(instruction))
		// Check if the new position is within the lawn
		if mower.X < 0 || mower.X > lawn.Width || mower.Y < 0 || mower.Y > lawn.Height {
			// If outside the lawn, revert the move and break the loop
			mower.MoveBackward()
			continue
		}
	}

	return mower
}
