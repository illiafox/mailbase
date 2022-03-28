package config

import (
	"encoding/json"
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

// Package provide many formats

type Format int

const (
	// Count starts from 1 to make FormatMap check easier. Map returns 0 whether it has not required format
	// Example:
	//	if f := FormatMap["json"]; f == 0 { return error }
	_ = Format(iota)

	// JSON https://pkg.go.dev/encoding/json
	JSON

	// TOML https://pkg.go.dev/github.com/pelletier/go-toml
	TOML

	// YAML https://pkg.go.dev/gopkg.in/yaml.v3
	YAML
)

// FormatMap is a map implementation of all formats
// Can be used with flag package
var FormatMap = map[string]Format{
	"json": JSON,
	"toml": TOML,
	"yaml": YAML,
}

// Available is string implementation of all formats. Can be used with flag package
// Example:
//	'json toml yaml'
var Available = func() string {
	ret := make([]string, len(FormatMap))
	i := 0
	for k := range FormatMap {
		ret[i] = k
		i++
	}
	return strings.Join(ret, " ")
}()

// ReadConfig reads and parses config using available formats (JSON, TOML, YAML)
// Example:
//	config.ReadConfig("config.json",config.JSON)
func ReadConfig(filename string, format Format) (Config, error) {
	var conf Config

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return conf, err
	}
	switch format {
	default:
		return conf, fmt.Errorf("unknown format %d", format)

	case YAML:
		err = yaml.Unmarshal(data, &conf)
		if err != nil {
			return conf, fmt.Errorf("yaml parsing: %w", err)
		}

	case JSON:
		err = json.Unmarshal(data, &conf)
		if err != nil {
			return conf, fmt.Errorf("json parsing: %w", err)
		}

	case TOML:
		err = toml.Unmarshal(data, &conf)
		if err != nil {
			return conf, fmt.Errorf("toml parsing: %w", err)
		}
	}

	return conf, nil
}
