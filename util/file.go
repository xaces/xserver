package util

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// JSONFile
func JSONFile(filename string, v interface{}) error {
	jsonFp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jsonFp.Close()
	var jsString string
	iReader := bufio.NewReader(jsonFp)
	for {
		tString, err := iReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		jsString = jsString + tString
	}
	return json.Unmarshal([]byte(jsString), v)
}


func YamlFile(filename string, v interface{}) error {
	yfile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yfile, v)
}