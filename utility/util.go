package utility

import (
	"encoding/json"
	"fmt"
	"strings"

	aero "github.com/aerospike/aerospike-client-go/v5"

	"github.com/google/uuid"
)

func ConvertToJson(bins aero.BinMap) []byte {
	jsonValue, err := json.Marshal(bins)
	PanicOnError(err)
	return jsonValue
}

func GetUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func LogMessage(msg string) {
	fmt.Println(msg)
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

