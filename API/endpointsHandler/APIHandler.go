package endpointsHandler


import (
	"../coreFunctions"
	"errors"
)

func HandlerCreate(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(d,[]string{"name","brand","value","createdAt","expiry"})==true {
		e := coreFunctions.CreateUpdate(d["name"].(string),d["brand"].(string),
			d["value"].(string),d["createdAt"].(string), d["expiry"].(string),false)
		err = e
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}

func HandlerUpdate(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(d,[]string{"name","brand","value","createdAt","expiry"})==true {
		e := coreFunctions.CreateUpdate(d["name"].(string),d["brand"].(string),
			d["value"].(string),d["createdAt"].(string), d["expiry"].(string),true)
		err = e
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}

func HandlerRetrieve(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(d,[]string{"name"})==true {
		fullList,e := coreFunctions.Retrieve(d["name"].(string))
		err = e
		dict = map[string]interface{}{"list": fullList}
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}


func HandlerList(_ map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	fullList, e := coreFunctions.List()
	err = e
	dict = map[string]interface{}{"list": fullList}
	return composeJson(dict,err)
}
