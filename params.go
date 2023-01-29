package ginkit

import "github.com/gin-gonic/gin"

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

func NewParams(p gin.Params) Params {
	var params Params
	for _, v := range p {
		params = append(params, Param{
			Key:   v.Key,
			Value: v.Value,
		})
	}

	return params
}
