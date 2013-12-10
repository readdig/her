package handy

type config struct {
	conf map[string]interface{}
	key  string
}

type configFuc interface {
	String(key string) string
	Int(key string) int
	Int64(key string) int64
	Bool(key string) bool
	Float(key string) float64
}

func LoadConfig(conf map[string]interface{}) *config {
	return &config{conf: conf}
}

func (c *config) Get(key string) *config {
	c.key = key
	return c
}

func (c *config) String() string {
	val, ok := c.conf[c.key].(string)
	if !ok {
		return ""
	}
	return val
}

func (c *config) Int() int {
	val, ok := c.conf[c.key].(int)
	if !ok {
		return 0
	}
	return val
}

func (c *config) Int64() int64 {
	val, ok := c.conf[c.key].(int64)
	if !ok {
		return 0
	}
	return val
}

func (c *config) Bool() bool {
	val, ok := c.conf[c.key].(bool)
	if !ok {
		return false
	}
	return val
}
func (c *config) Float() float64 {
	val, ok := c.conf[c.key].(float64)
	if !ok {
		return 0.0
	}
	return val
}
