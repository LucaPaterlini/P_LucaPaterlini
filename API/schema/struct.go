package schema

import (
	"time"
)


//Item struct of the items of the perkstable
type Item struct {
	Name    string    `json:"name"`
	Brand   string    `json:"brand"`
	Value   int64     `json:"value"`
	Created time.Time `json:"created"`
	Expiry  time.Time `json:"expiry"`
}
//ResponseJson tamplete for every json answer of the api
type ResponseJson struct {
	Err  bool                   `json:"err"`
	Data map[string]interface{} `json:"data"`
}
