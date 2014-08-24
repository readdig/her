package her

type MergedConfig struct {
	config map[string]interface{}
}

func loadConfig(conf map[string]interface{}) *MergedConfig {
	mergedConfig := &MergedConfig{}

	if mergedConfig.String("Address") == "" {
		conf["Address"] = "0.0.0.0"
	}

	if mergedConfig.Int("Port") == 0 {
		conf["Port"] = 8080
	}

	if mergedConfig.String("TemplatePath") == "" {
		conf["TemplatePath"] = "view"
	}

	if mergedConfig.String("CookieSecret") == "" {
		conf["CookieSecret"] = "kN)A/hJ]ZsmHk#5'=Q'88zv6]vqfzS"
	}

	mergedConfig.config = conf

	return mergedConfig
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
