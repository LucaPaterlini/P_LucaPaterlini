package endpointsHandler

import "encoding/json"
import "../schema"

func composeJson(params map[string]interface{},err error)(s string){
	objR := &schema.ResponseJson{}
	objR.Err=err!=nil
	objR.Data=make(map[string]interface{})
	if err==nil{
		for k,v := range params{objR.Data[k]=v}
	}else {
		objR.Data["errMsg"]=err.Error()
	}
	jsonString, _ := json.MarshalIndent(objR,"","\t")
	return string(jsonString)
}

func checkKeys(d map[string]interface{},key[]string)bool{
	for _,v := range key {
		_,ok:=d[v]
		if ok==false{return false}
	}
	return true
}

