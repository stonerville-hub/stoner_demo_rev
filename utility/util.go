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

var client *as.Client
var db *DBAero

type DBAero struct {
	Host      string
	Port      string
	Namespace string
	Set       string
	Client    *as.Client
}

func ConvertToJson(bins as.BinMap) string {
	jsonBytes, err := json.Marshal(bins)
	PanicOnError(err)
	return string(jsonBytes[:])
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

func GetConnection() *DBAero {
	host := GetEnv(HOST, defaultHost)
	port := GetEnv(PORT, defaultPort)
	client := getClient(host, port)
	db = &DBAero{host, port, GetEnv(NAMESPACE, defaultNS), GetEnv(SET, defaultSet), client}
	return db
}

func getClient(host string, port string) *as.Client {
	if client != nil && client.IsConnected() {
		return client
	}
	return retrieveClient(host, port)
}

func retrieveClient(host string, port string) *as.Client {
	errDBDown := "Database is DOWN"
	_port, err := strconv.Atoi(port)
	PanicOnError(err)
	_client, err := as.NewClient(host, _port)
	if err != nil {
		panic(errDBDown+" for " + host + ":" + port)
	}
	LogMessage("Database is UP for " + host + ":" + port)
	return _client
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && len(value)>0  {
		return value
	}
	return fallback
}
