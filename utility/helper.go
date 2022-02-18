package utils

import (
	"log"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/my/repo/model"

	as "github.com/aerospike/aerospike-client-go/v5"
	"github.com/google/uuid"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ConvertToJson(bins as.BinMap) string {
	jsonBytes, err := json.Marshal(bins)
	PanicOnError(err)
	return string(jsonBytes[:])
}

func ConvertFromJson(jsonData string, u model.Customer) model.Customer {	
	if err := json.Unmarshal([]byte(jsonData), &u); err != nil {
		PanicOnError(err)
	}
	return u
}

func GetUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && len(value) > 0 {
		return value
	}
	return fallback
}

func LogMessage(msg string) {
	log.Output(303, msg)
}

func PanicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}