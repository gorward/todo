package app

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// GenerateUUID generates uuid using nu7hatch/guuids
func GenerateUUID() string {
	uuid, _ := uuid.NewV4()

	return uuid.String()
}

// StructToBson converts structs to bson.M (map[string]interface{})
func StructToBson(data interface{}) bson.M {
	newData := bson.M{}

	out, _ := bson.Marshal(data)
	bson.Unmarshal(out, newData)

	return newData
}

// SetExisting format bson.M keys to only update existing keys
func SetExisting(prefix string, data bson.M) bson.M {
	newData := bson.M{}

	for k, v := range data {
		newData[prefix+"."+k] = v
	}

	return newData
}

// RespondErr helper to respond with error
func RespondErr(w http.ResponseWriter, status int, data string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{
		"error": %s
	}`, data)
}

// RespondData helper to respond with data
func RespondData(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{
		"data": %s
	}`, string(data))
}
