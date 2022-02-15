package aerospike

import (
	"encoding/json"
	"html/template"
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
	data := make([]util.User, 0)
	for rec := range recordset.Results() {
		if rec.Err != nil {
			util.PanicOnError(rec.Err)
		}
		jsonData := util.ConvertToJson(rec.Record.Bins)
		u := util.User{}
		if err := json.Unmarshal([]byte(jsonData), &u); err != nil {
			util.PanicOnError(err)
		}
		data = append(data, u)
	}
	tmpl := template.Must(template.ParseFiles("home.html"))
	tmpl.Execute(w, data)
}

func LoadNewCustomers(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 10; i++ {
		insertRecord(i)
	}
	http.Redirect(w, r, util.HOME, http.StatusSeeOther)
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

func insertRecord(recnum int) {
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
		util.LogMessage("Failed inserting api_key " + apikey)
	}
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
	db.Client.IsConnected()
}

func WriteMessage(w http.ResponseWriter, msg ...string) {
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(msg, ",")))
}
