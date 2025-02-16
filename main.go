package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

const (
	TYPE      string = "todos" // post, users, todos, comments
	FILE_NAME string = "data_" + TYPE + ".json"
	URL_GET   string = "https://dummyjson.com/" + TYPE
)

func writeJson(data any, filename string) error {
	// write to files
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("error marchal %s: %w", filename, err)
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error write file %s: %w", filename, err)
	}

	fmt.Printf("JSON data write to file ~ %s\n", filename)
	return nil
}

func loadJson(filename string) (any, error) {
	// loading files
	plan, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(plan, &data)
	if err != nil {
		return nil, err
	}

	dataInterFace, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("want type map[string]interface{};  got %T", data)
	}

	for k, v := range dataInterFace {
		fmt.Println(k, "=>", v)
	}

	return dataInterFace, nil
}

func main() {
	resp, err := http.Get(URL_GET)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// add write files ~ 1
	err = writeJson(data, FILE_NAME)
	if err != nil {
		fmt.Print(err)
		return
	}

	// loading data files ~ 2
	dif, err := loadJson(FILE_NAME)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("type data ~ %s\n", reflect.TypeOf(dif))
}
