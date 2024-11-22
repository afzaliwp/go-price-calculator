package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func NewFileManager(inputPath string, outputPath string) *FileManager {
	return &FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

func (fm *FileManager) ReadFile() (data []float64, err error) {
	file, err := os.Open(fm.InputPath)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to open resource %s", fm.InputPath)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		priceInFloat, err := strconv.ParseFloat(text, 64)
		if err != nil {
			_ = file.Close()
			fmt.Println(err)
			return nil, errors.New("failed to read the file")
		}

		data = append(data, priceInFloat)
	}

	_ = file.Close()

	return data, nil
}

func (fm *FileManager) OutputJsonFile(data interface{}) error {
	time.Sleep(3 * time.Second) //only to see how routines are effecting
	file, err := os.Create(fm.OutputPath)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to save resource %s", fm.OutputPath)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	//This code is optional, it is used just to make the json formatted well
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to save resource %s", fm.OutputPath)
	}

	return err
}
