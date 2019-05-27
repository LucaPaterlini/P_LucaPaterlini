package endpointsHandler


//ResponseJson template for every json answer of the api
type ResponseJson struct {
	Err  bool                   `json:"err"`
	Data map[string]interface{} `json:"data"`
}

