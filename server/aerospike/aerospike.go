package aerospike

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	as "github.com/aerospike/aerospike-client-go/v5"
	jsoniter "github.com/json-iterator/go"
	util "github.com/my/repo/utility"
)

var (
	client = retrieveClient()
	// _enviroment  = os.Args[1]
	_host       = os.Args[2]
	_port       = os.Args[3]
	_namespace  = os.Args[4]
	_set        = os.Args[5]
	_api_key    = "api_key"
	_first_name = "first_name"
	_last_name  = "last_name"
	_company    = "company"
)

func PreloadCustomers(w http.ResponseWriter, r *http.Request) {
	WriteRecord(10)
}

func GetRecord(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// pathVar := r.URL.Path[len("/user/"):]
	query := r.URL.Query()
	filters, present := query[_api_key] //filters=["color", "price", "brand"]
	if !present || len(filters) == 0 {
		util.LogMessage(_api_key + " is not present")
	}
	ReadRecord(filters[0])
	util.LogMessage("testing URL Path=" + r.URL.Path)
}

func WriteRecord(numberOfRecords int) {
	util.LogMessage("Generating Test Records\n\n")
	util.LogMessage("Here are a few api_keys for testing:\n")

	for i := 0; i < numberOfRecords; i++ {
		// apikey := util.RandomInt64(int64(i))
		apikey := util.GetUUID()
		key, err := as.NewKey(_namespace, _set, apikey)
		util.PanicOnError(err)
		bins := as.BinMap{
			_api_key:    apikey,
			_first_name: "Rev" + apikey,
			_last_name:  "Content" + apikey,
			_company:    "Revcontent",
		}
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Marshal(bins)
		// util.LogMessage("New Record created: " + bins.Bins["api_key"])
		// write the bins
		err = client.Put(nil, key, bins)
		util.PanicOnError(err)

		rec, err := client.Get(nil, key)
		if err != nil {
			util.PanicOnError(err)
		}
		fmt.Println(rec.Bins[_api_key])
	}

	util.LogMessage("\nTest Records created.\n")
}

func ReadRecord(_key string) {
	util.LogMessage("Reading record")
	key, err := as.NewKey(_namespace, _set, _key)
	if err != nil {
		// read it
		rec, err := client.Get(nil, key)
		util.PanicOnError(err)
		util.LogMessage("Record: " + rec.String())
	}
	util.LogMessage("\nEnd Processing\n\n")
}

func CheckDBConnection() {
	util.LogMessage("Checking database connection...")
	// Create a new client and connect to the server
	// Ping the primary
	client.IsConnected()
}
func retrieveClient() *as.Client {
	port, err := strconv.Atoi(_port)
	util.PanicOnError(err)
	client, err := as.NewClient(_host, port)
	if err != nil {
		util.LogMessage("*******WARNING*************")
		util.LogMessage("Database is DOWN")
		util.LogMessage("*******WARNING*************")
	} else {
		util.LogMessage("Database is UP")
	}
	util.PanicOnError(err)

	return client
}
