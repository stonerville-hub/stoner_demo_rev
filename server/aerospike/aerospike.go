package aerospike

import (
	"net/http"
	"strconv"
	"strings"

	as "github.com/aerospike/aerospike-client-go/v5"
	util "github.com/my/repo/utility"
)

var db *util.DBAero = util.GetConnection()

func HomePage(w http.ResponseWriter, r *http.Request) {
	recordset, err := db.Client.ScanAll(nil, db.Namespace, db.Set)
	util.PanicOnError(err)
	WriteMessage(w, "<h3>List of Test Record(s)</h3>")
	WriteMessage(w, "<h3>------------------------------</h3>")
	recordsExists := false
	// consume recordset and check errors
	for rec := range recordset.Results() {
		if rec.Err != nil {
			util.PanicOnError(rec.Err)
		}
		WriteMessage(w, "<h4>"+util.ConvertToJson(rec.Record.Bins)+"</h4>")
	}

	if !recordsExists {
		WriteMessage(w, "<h2>No records found.  For loading test data, click <a href="+r.RemoteAddr+">here</a></h2>")
	}
}

func LoadNewCustomers(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 10; i++ {
		b := insertRecord(i)
		WriteMessage(w, util.ConvertToJson(b))
	}
	http.Redirect(w, r, util.HOME, http.StatusOK)
}

func GetRecordByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filters, present := query[util.APIKEY]
	if !present || len(filters) == 0 {
		WriteMessage(w, util.APIKEY+" is not present")
	} else {
		rec := readRecord(filters[0])
		if rec == nil {
			WriteMessage(w, "Api key "+filters[0]+" was not found")
		} else {
			WriteMessage(w, util.ConvertToJson(rec.Bins))
		}
	}
}

func insertRecord(recnum int) as.BinMap {
	apikey := util.GetUUID()
	key, err := as.NewKey(db.Namespace, db.Set, apikey)
	util.PanicOnError(err)
	r := strconv.Itoa(recnum)

	bins := as.BinMap{
		util.APIKEY:    apikey,
		util.FIRSTNAME: "Rev_" + r,
		util.LASTNAME:  "Content_" + r,
		util.COMPANY:   "Revcontent",
	}
	err = db.Client.Put(nil, key, bins)
	if err != nil {
		util.LogMessage("Failed inserting api_key "+apikey)
	}
	return bins
}

func readRecord(_key string) *as.Record {
	util.LogMessage("Reading record")
	key, err := as.NewKey(db.Namespace, db.Set, _key)
	util.PanicOnError(err)
	rec, err := db.Client.Get(nil, key)
	if err != nil {
		rec = nil
	}
	return rec
}

func CheckDBConnection() {
	util.LogMessage("Checking database connection...")
	db.Client.IsConnected()
}

func WriteMessage(w http.ResponseWriter, msg ...string) {
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(msg, ",")))
}
