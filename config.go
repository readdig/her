package her

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type MergedConfig struct {
	config map[string]interface{}
}

func loadConfig(a ...interface{}) *MergedConfig {
	config := make(map[string]interface{})

	if len(a) > 0 {
		switch conf := a[0].(type) {
		case map[string]interface{}:
			config = conf
		case string:
			config = loadConfigJSON(conf)
		}
	} else {
		config = loadConfigJSON("config.json")
	}

	mergedConfig := &MergedConfig{config: config}

	if mergedConfig.String("Address") == "" {
		mergedConfig.config["Address"] = "0.0.0.0"
	}

	if mergedConfig.Int("Port") == 0 {
		mergedConfig.config["Port"] = 8080
	}

	if mergedConfig.String("TemplatePath") == "" {
		mergedConfig.config["TemplatePath"] = "view"
	}

	if mergedConfig.String("CookieSecret") == "" {
		mergedConfig.config["CookieSecret"] = "kN)A/hJ]ZsmHk#5'=Q'88zv6]vqfzS"
	}

	return mergedConfig
}

func loadConfigJSON(filename string) map[string]interface{} {
	var conf map[string]interface{}

	if pathExist(filename) {
		bytes, err := ioutil.ReadFile(filename)

		if err != nil {
			log.Println(err.Error())
			return nil
		}

		if err := json.Unmarshal(bytes, &conf); err != nil {
			log.Println(err.Error())
			return nil
		}
	} else {
		conf = make(map[string]interface{})
	}

	return conf
}

func (c *MergedConfig) String(key string) string {
	val, ok := c.config[key].(string)
	if !ok {
		return ""
	}
	return val
}

func (c *MergedConfig) Int(key string) int {
	val, ok := c.config[key].(int)
	if !ok {
		return 0
	}
	return val
}

func (c *MergedConfig) Int64(key string) int64 {
	val, ok := c.config[key].(int64)
	if !ok {
		return 0
	}
	return val
}

func (c *MergedConfig) Bool(key string) bool {
	val, ok := c.config[key].(bool)
	if !ok {
		return false
	}
	return val
}

func (c *MergedConfig) Float(key string) float64 {
	val, ok := c.config[key].(float64)
	if !ok {
		return 0.0
	}
	return val
}
