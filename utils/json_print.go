package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) ([]byte, error) {
	b, _ := json.Marshal(v)
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func PrettyPrint2(v interface{}) {
	b, _ := json.Marshal(v)
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out.Bytes()))
}
