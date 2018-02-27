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

	json_data, err := ToJSON(testStruct)

	assertNil(t, err)

	result := TestStruct{}

	err = FromJSON(json_data, &result)

	assertNil(t, err)

	fmt.Printf("json: %s\n", json_data)

	fmt.Printf("decoded: %#v\n",result )

	assertEqual(t, testStruct, result)

}

func TestToJSON(t *testing.T) {

	testStruct := TestStruct{Foo: "Foo", Bar: "Bar", Cat: 1, Dog: 1}

	jsonData, err := ToJSON(testStruct)
	
	jsonData = strings.TrimSpace(jsonData)

	assertNil(t,err)

	assertEqual(t, jsonData, "{\"Foo\":\"Foo\",\"Bar\":\"Bar\",\"Cat\":1,\"Dog\":1}")
	

}

func TestEncodeArray(t *testing.T) {

	testStruct := TestStruct{Foo: "Foo", Bar: "Bar", Cat: 1, Dog: 1}
	structs := make([]TestStruct, 0)

	structs = append(structs, testStruct)
	assertEqual(t, len(structs), 1)

	jsonString, err := ToJSON(structs)

	assertNil(t, err)

	assertEqual(t, jsonString, "[{\"Foo\":\"Foo\",\"Bar\":\"Bar\",\"Cat\":1,\"Dog\":1}]")

}