package her

type MergedConfig struct {
	config map[string]interface{}
}

func LoadConfig(conf map[string]interface{}) *MergedConfig {
	return &MergedConfig{config: conf}
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
