package schema

import "time"

type Item struct {
	Name string `json:"name"`
	Brand string `json:"brand"`
	Value int64 `json:"value"`
	Created time.Time `json:"created"`
	Expiry time.Time `json:"expiry"`
}

type ResponseJson struct {
	Err bool `json:"err"`
	Data map[string]interface{} `json:"data"`
}