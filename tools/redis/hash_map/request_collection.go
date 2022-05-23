package hash_map

type Collection struct {
	Request map[string]string
	Result  []string
}

func NewCollection(req map[string]string) *Collection {
	return &Collection{
		Request: req,
	}
}

func (c *Collection) GetArgs() (result []string) {
	for k, v := range c.Request {
		result = append(result, k, v)
	}
	return
}
