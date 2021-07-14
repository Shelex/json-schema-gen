package fixture

import (
	"encoding/json"
	"errors"
	"log"
	"path/filepath"
	"strings"

	"github.com/Shelex/json-schema-gen/pkg/prompt"
	"github.com/spf13/afero"
)

var fixtureFolder = filepath.Join("cypress", "fixtures")

// File - file required for sending request
type File struct {
	Content map[string]interface{}
	Path    string
}

// Select - prompt with content of cypress/fixtures folder to select fixture with request data
func Select(prompt prompt.Executor, os afero.Afero, selected string) (File, error) {
	currentPath := filepath.Join("cypress", "fixtures")
	var fixture File
	for selected == "" {
		files, err := os.ReadDir(currentPath)
		if err != nil {
			return fixture, err
		}
		filePaths := make([]string, 0, len(files)+1)
		if !strings.HasSuffix(currentPath, fixtureFolder) {
			filePaths = append(filePaths, "..")
		}
		for _, f := range files {
			filePaths = append(filePaths, f.Name())
		}
		result, err := prompt.Select("Path to fixture with response data", filePaths, 10)
		if err != nil {
			return fixture, err
		}

		currentPath = filepath.Join(currentPath, result)
		stat, err := os.Fs.Stat(currentPath)
		if err != nil {
			return fixture, err
		}
		if !stat.IsDir() {
			selected = currentPath
		}
	}
	log.Println("selected fixture:", selected)

	fixture, err := read(selected, os)
	if err != nil {
		return fixture, err
	}
	return fixture, nil
}

// Read - read fixture content
func read(path string, os afero.Afero) (File, error) {
	var fixture File
	byteValue, err := os.ReadFile(path)

	if err != nil {
		log.Print(err)
		return fixture, err
	}

	if err := json.Unmarshal(byteValue, &fixture.Content); err != nil {
		return fixture, errors.New("fixture reading failed")
	}
	fixture.Path = path
	return fixture, nil
}

func overwrite(prompt prompt.Executor) (bool, string, error) {
	confirmOverwrite, err := prompt.Confirm("File already exist. Overwrite file")
	if err != nil {
		return false, "", err
	}
	if confirmOverwrite {
		return true, "", nil
	}
	newName, err := prompt.Input("Please enter file name for schema")
	if err != nil {
		return false, "", err
	}
	return false, newName, nil
}

// Validate - all text inputs to avoid dots and dashes
func Validate(s string) error {
	if strings.Contains(s, "-") || strings.Contains(s, ".") || len(s) == 0 {
		return errors.New("invalid value provided")
	}
	return nil
}
