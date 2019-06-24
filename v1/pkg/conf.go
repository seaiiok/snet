package pkg

import (
	"encoding/json"
	"io/ioutil"

	"github.com/seaiiok/utils/files"
)

func LoadConfigFile(path string) (config map[string]string) {
	if !files.IsExists(path) {
		return nil
	}
	config = make(map[string]string, 0)
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(conf, &config)
	if err != nil {
		return nil
	}
	return
}
