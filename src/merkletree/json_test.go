package merkletree

import (
	"testing"
	"fmt"
	"encoding/json"
	"bytes"
	"strings"
)

type TestStruct struct {

	Foo string

	Bar string

	Cat int

	Dog int

}

func ToJSON(obj interface{}) string {

	data := make([]byte, 0)

	buffer := bytes.NewBuffer(data)

	encoder := json.NewEncoder(buffer)

	encoder.Encode(obj)

	return buffer.String()

}

func FromJSON(str string, v interface{}) interface{} {

	data := make([]byte, 0)

	buffer := bytes.NewBuffer(data)

	buffer.WriteString(str)

	decoder := json.NewDecoder(buffer)

	return decoder.Decode(v)

}

func TestJSONEncoding(t *testing.T) {

	testStruct := TestStruct{Foo: "Foo", Bar: "Bar", Cat: 1, Dog: 1}

	data := make([]byte, 0)

	buffer := bytes.NewBuffer(data)

	encoder := json.NewEncoder(buffer)

	encoder.Encode(testStruct)

	json_data := buffer.String()

	fmt.Println(json_data)

}

func TestJSONEncodingFunctions(t *testing.T) {

	testStruct := TestStruct{Foo: "Foo", Bar: "Bar", Cat: 1, Dog: 1}

	json_data := ToJSON(testStruct)

	result := TestStruct{}

	FromJSON(json_data, &result)

	fmt.Printf("json: %s\n", json_data)

	fmt.Printf("decoded: %#v\n",result )

}

func TestToJSON(t *testing.T) {

	testStruct := TestStruct{Foo: "Foo", Bar: "Bar", Cat: 1, Dog: 1}

	jsonData := strings.TrimSpace(ToJSON(testStruct))

	assertEqual(t, jsonData, "{\"Foo\":\"Foo\",\"Bar\":\"Bar\",\"Cat\":1,\"Dog\":1}")
	

}