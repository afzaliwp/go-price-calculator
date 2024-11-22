package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(filepath string) (data []float64, err error) {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to open resource %s", filepath)
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

func SaveJson(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to save resource %s", path)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	//This code is optional, it is used just to make the json formatted well
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to save resource %s", path)
	}

	return err
}
