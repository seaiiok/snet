package pkg

import (
	"encoding/json"
	"io/ioutil"
)

func loadConfigFile(configFile string) (config map[string]string) {

	config = make(map[string]string, 0)
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(conf, &config)
	if err != nil {
		return nil
	}

	return nil
}
