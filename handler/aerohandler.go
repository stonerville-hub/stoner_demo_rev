package handler

import (
	"net/http"
	"strconv"

	as "github.com/aerospike/aerospike-client-go/v5"
	"github.com/gin-gonic/gin"
	"github.com/my/repo/model"
	s "github.com/my/repo/server"
	u "github.com/my/repo/utility"
)

func HomePage(c *gin.Context) {
	recordset, err := s.Client.ScanAll(nil, s.Connectaero.Namespace, s.Connectaero.Set)
	u.PanicOnError(err)
	data := make([]model.Customer, 0)
	for rec := range recordset.Results() {
		if rec.Err != nil {
			u.PanicOnError(rec.Err)
		}
		data = append(data, u.ConvertFromJson(u.ConvertToJson(rec.Record.Bins), model.Customer{}))
	}
	c.HTML(http.StatusOK, "home.html", data)
}

func LoadNewData(c *gin.Context) {
	for i := 0; i < 10; i++ {
		insertRecord(i)
	}
	c.Redirect(http.StatusSeeOther, u.HOME)
}

func GetCustomerByID(c *gin.Context) {
	customerID := c.Param("id")
	query := c.Query(u.APIKEY)
	if len(query) == 0 {
		c.String(http.StatusNoContent, "Customer "+customerID+" with api key "+u.APIKEY+" is not present")
	} else {
		rec := readRecord(query)
		if rec == nil {
			c.String(http.StatusOK, "Customer "+customerID+" with api key "+query+" was not found")
		} else {
			c.String(http.StatusOK, u.ConvertToJson(rec.Bins))
		}
	}
}

func insertRecord(recnum int) {
	apikey := u.GetUUID()
	key, err := as.NewKey(s.Connectaero.Namespace, s.Connectaero.Set, apikey)
	u.PanicOnError(err)
	r := strconv.Itoa(recnum)

	bins := as.BinMap{
		u.APIKEY:    apikey,
		u.FIRSTNAME: "Rev_" + r,
		u.LASTNAME:  "Content_" + r,
		u.COMPANY:   "Revcontent",
	}

	err = s.Client.Put(nil, key, bins)
	if err != nil {
		u.LogMessage("Failed inserting api_key " + apikey)
	}
}

func readRecord(_key string) *as.Record {
	u.LogMessage("Reading record")
	key, err := as.NewKey(s.Connectaero.Namespace, s.Connectaero.Set, _key)
	u.PanicOnError(err)
	rec, err := s.Client.Get(nil, key)
	if err != nil {
		rec = nil
	}
	return rec
}
