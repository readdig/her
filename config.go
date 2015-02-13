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

	if mergedConfig.GetString("Address") == "" {
		mergedConfig.config["Address"] = "0.0.0.0"
	}

	if mergedConfig.GetInt("Port") == 0 {
		mergedConfig.config["Port"] = 8080
	}

	if mergedConfig.GetString("TemplatePath") == "" {
		mergedConfig.config["TemplatePath"] = "view"
	}

	if mergedConfig.GetString("TemplateExt") == "" {
		mergedConfig.config["TemplateExt"] = ".html"
	}

	if mergedConfig.GetString("CookieSecret") == "" {
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

func (c *MergedConfig) GetString(key string) string {
	val, ok := c.config[key]
	if !ok {
		return ""
	}
	return val.(string)
}

func (c *MergedConfig) GetInt(key string) int {
	val, ok := c.config[key]
	if !ok {
		return -1
	}
	return int(val.(float64))
}

func (c *MergedConfig) GetBool(key string) bool {
	val, ok := c.config[key]
	if !ok {
		return false
	}
	return val.(bool)
}

func (c *MergedConfig) GetFloat(key string) float64 {
	val, ok := c.config[key]
	if !ok {
		return -1
	}
	return val.(float64)
}

func (c *MergedConfig) GetMap(key string) map[string]interface{} {
	val, ok := c.config[key]
	if !ok {
		return nil
	}
	return val.(map[string]interface{})
}

func (c *MergedConfig) GetArray(key string) []interface{} {
	val, ok := c.config[key]
	if !ok {
		return []interface{}(nil)
	}
	return val.([]interface{})
}
