package merkletree

import "encoding/json"

// ToJSON takes a struct and converts it to a JSON string.
func ToJSON(obj interface{}) (string, error) {

	//if data, err := json.MarshalIndent(obj, "", "  "); err == nil {
	if data, err := json.Marshal(obj); err == nil {
		return string(data), err
	} else {
		return "", err
	}

}

// FromJSON takes a string and converts it to a struct
func FromJSON(str string, v interface{}) error {

	data := []byte(str)

	err := json.Unmarshal(data, v)

	return err

}

