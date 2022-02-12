package aerospike

import (
	"fmt"
	"net/http"

	as "github.com/aerospike/aerospike-client-go/v5"
	util "github.com/my/repo/utility"
)

var dbconfig *util.Dbconfig = util.GetConnection()

func PreloadCustomers(w http.ResponseWriter, r *http.Request) {
	WriteRecord(10)
}

func GetRecord(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// pathVar := r.URL.Path[len("/user/"):]
	query := r.URL.Query()
	filters, present := query[util.APIKEY] //filters=["color", "price", "brand"]
	if !present || len(filters) == 0 {
		util.LogMessage(util.APIKEY + " is not present")
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
		key, err := as.NewKey(dbconfig.Namespace, dbconfig.Set, apikey)
		util.PanicOnError(err)
		bins := as.BinMap{
			util.APIKEY:    apikey,
			util.FIRSTNAME: "Rev" + apikey,
			util.LASTNAME:  "Content" + apikey,
			util.COMPANY:   "Revcontent",
		}
		// var json = jsoniter.ConfigCompatibleWithStandardLibrary
		// json.Marshal(bins)
		// util.LogMessage("New Record created: " + bins.Bins["api_key"])
		// write the bins
		err = dbconfig.Client.Put(nil, key, bins)
		util.PanicOnError(err)

		rec, err := dbconfig.Client.Get(nil, key)
		if err != nil {
			util.PanicOnError(err)
		}
		fmt.Println(rec.Bins[util.APIKEY])
	}

	util.LogMessage("\nTest Records created.\n")
}

func ReadRecord(_key string) {
	util.LogMessage("Reading record")
	key, err := as.NewKey(dbconfig.Namespace, dbconfig.Set, _key)
	if err != nil {
		// read it
		rec, err := dbconfig.Client.Get(nil, key)
		util.PanicOnError(err)
		util.LogMessage("Record: " + rec.String())
	}
	util.LogMessage("\nEnd Processing\n\n")
}

func CheckDBConnection() {
	util.LogMessage("Checking database connection...")
	// Create a new client and connect to the server
	// Ping the primary
	dbconfig.Client.IsConnected()
}
