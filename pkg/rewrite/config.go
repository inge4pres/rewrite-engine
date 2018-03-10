package rewrite

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// JSONConfig represents the configuartion as read from the JSON
type JSONConfig struct {
	Rules []*Rule `json:"rules"`
}

// ParseJSONConfig tries to read the config from the io.Reader provided as input
// Will return nil and the error if the parsing fails
func ParseJSONConfig(input io.Reader) (*JSONConfig, error) {
	b, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	conf := new(JSONConfig)
	err = json.Unmarshal(b, conf)
	return conf, err
}
