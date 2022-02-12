package utility

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	as "github.com/aerospike/aerospike-client-go/v5"

	"github.com/google/uuid"
)

type Dbconfig struct {
	Environment string
	Host        string
	Port        string
	Namespace   string
	Set         string
	Client      *as.Client
}

const (
	ENVIROMENT  = "ENVIROMENT"
	HOST        = "HOST"
	PORT        = "PORT"
	NAMESPACE   = "NAMESPACE"
	SET         = "SET"
	APIKEY      = "api_key"
	FIRSTNAME   = "first_name"
	LASTNAME    = "last_name"
	COMPANY     = "company"
	defaultEnv  = "dev"
	defaultHost = "aerospike"
	defaultPort = "3000"
	defaultNS   = "test"
	defaultSet  = "users"
)

func ConvertToJson(bins as.BinMap) []byte {
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

func GetConnection() *Dbconfig {
	host := GetEnv(HOST, defaultHost)
	port := GetEnv(PORT, defaultPort)
	client := retrieveClient(host, port)
	dbconfig := &Dbconfig{GetEnv(ENVIROMENT, defaultEnv), host, port, GetEnv(NAMESPACE, defaultNS), GetEnv(SET, defaultSet), client}
	return dbconfig
}

func retrieveClient(host string, port string) *as.Client {
	errDBDown := "Database is DOWN"
	_port, err := strconv.Atoi(port)
	PanicOnError(err)
	client, err := as.NewClient(host, _port)
	if err != nil {
		LogMessage("*******WARNING*************")
		LogMessage(errDBDown+" for "+host+":"+port)
		LogMessage("*******WARNING*************")
		panic(errDBDown)
	}
	LogMessage("Database is UP for "+host+":"+port)

	return client
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
