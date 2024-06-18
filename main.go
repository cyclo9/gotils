package utils

import (
	"encoding/json"
	"log"
	"os"
)

// taskDesc: task description; should be in present continuous tense
func HandleErr(err error, taskDesc ...string) {
	if err != nil {
		if len(taskDesc) > 0 {
			log.Fatalf("Error %s: %s", taskDesc[0], err)
		} else {
			log.Fatalf("Error: %s", err)
		}
		os.Exit(1)
	}
}

func ZeroValue[T any]() T {
	var zeroValue T
	return zeroValue
}

func WriteJSONToFile(filepath string, data interface{}) {
	file, err := os.Create(filepath)
	HandleErr(err, "creating file")
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	HandleErr(err, "writing to file")
}
