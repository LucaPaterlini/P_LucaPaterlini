//endpointsHandler test the core handlers of the api
// checking for each return status
package endpointsHandler

import (
	"../coreFunctions"
	"errors"
	"github.com/globalsign/mgo"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)
//HandlerStruct handler for the endpoints
type HandlerStruct struct {
	Collection *mgo.Collection
}

// setDefaultHeaders set the default values of the header
func setDefaultHeaders(w http.ResponseWriter) {
	// allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

// HandlerCreate check the parameters,run the CreateUpdate and return the json response
func (ctx HandlerStruct) HandlerCreateUpdate(w http.ResponseWriter, r *http.Request) {
	// Consider only GET requests
	if r.Method == http.MethodGet {
		setDefaultHeaders(w)
		m, _ := url.ParseQuery(r.URL.RawQuery)
		dict := make(map[string]interface{})
		// converting args strings to the appropriate a types
		value, err := strconv.ParseInt(m.Get("value"), 10, 64)
		var update bool
		if err == nil {
			update, err = strconv.ParseBool(m.Get("update"))
		}
		if err != nil {
			log.Print(err.Error())
			err = errors.New(r.URL.Path + ": wrong set of input parameters")
			w.WriteHeader(http.StatusBadRequest)

		}
		// checking the correct execution of CreateUpdate
		if err == nil {
			err = coreFunctions.CreateUpdate(m.Get("name"), m.Get("brand"),
				value, m.Get("createdAt"), m.Get("expiry"), update, ctx.Collection)
			if err != nil {
				if mgo.IsDup(err) {
					err = errors.New(r.URL.Path + ": Entry Duplicate")
					w.WriteHeader(http.StatusBadRequest)
				} else {
					err = errors.New(r.URL.Path + ": InternalError")
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
		// preparing the Json Response
		_, err = w.Write([]byte(ComposeJson(dict, err)))
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandlerCreate check the parameters,run the CreateUpdate and return the json response
func (ctx HandlerStruct) HandlerRetrieve(w http.ResponseWriter, r *http.Request) {
	// Consider only GET requests
	if r.Method == http.MethodGet {
		setDefaultHeaders(w)
		m, _ := url.ParseQuery(r.URL.RawQuery)
		dict := make(map[string]interface{})
		// converting args strings to the appropriate a types
		active, err := strconv.ParseBool(m.Get("active"))
		if err != nil {
			log.Print(err.Error())
			err = errors.New(r.URL.Path + ": wrong set of input parameters")
			w.WriteHeader(400)
		}
		// checking the correct execution of Retrieve
		if err == nil {
			fullList, e := coreFunctions.Retrieve(m.Get("name"), active, time.Now(), ctx.Collection)
			err = e
			dict = map[string]interface{}{"list": fullList}
			if err != nil {
				log.Print(err.Error())
				err = errors.New(r.URL.Path + ": InternalError")
				w.WriteHeader(500)
			} else {
				dict = map[string]interface{}{"list": fullList}
			}
		}
		// preparing the Json Response
		_, err = w.Write([]byte(ComposeJson(dict, err)))
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(500)
		}
		return
	} else {
		w.WriteHeader(405)
	}
}
