package endpointsHandler

import (
	"../coreFunctions"
	"errors"
)
// HandlerCreate check the parameters,run the CreateUpdate and return the json response
func HandlerCreateUpdate(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if CheckKeys(d,[]string{"name","brand","value","createdAt","expiry","update"})==true {
		e := coreFunctions.CreateUpdate(d["name"].(string),d["brand"].(string),
			d["value"].(int64),d["createdAt"].(string), d["expiry"].(string),d["update"].(bool),false)
		err = e
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return ComposeJson(dict,err)
}

// HandlerUpdate check the parameters,run the CreateUpdate and return the json response
func HandlerRetrieve(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if CheckKeys(d,[]string{"name"})==true {
		fullList,e := coreFunctions.Retrieve(d["name"].(string), d["active"].(bool), d["debug"].(bool))
		err = e
		dict = map[string]interface{}{"list": fullList}
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return ComposeJson(dict,err)
}