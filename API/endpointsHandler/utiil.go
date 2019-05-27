package endpointsHandler

import (
	"encoding/json"
)



// ComposeJson return a formatted json ready to be sent as response message
// taking as input the parameters to return and if present the error
func ComposeJson(params map[string]interface{}, err error) (s string) {
	objR := &ResponseJson{}
	objR.Err = err != nil
	objR.Data = make(map[string]interface{})
	if err == nil {
		for k, v := range params {
			objR.Data[k] = v
		}
	} else {
		objR.Data["errMsg"] = err.Error()
	}
	jsonString, _ := json.MarshalIndent(objR, "", "\t")
	return string(jsonString)
}
