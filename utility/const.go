package utility

const (
    HOME           = "/"
    LOADNEWDATA    = "/loadNewData"
    GET_USERBYID   = "/user/{id}"
)

const (
	HOST        = "HOST"
	PORT        = "PORT"
	NAMESPACE   = "NAMESPACE"
	SET         = "SET"
	APIKEY      = "api_key"
	FIRSTNAME   = "first_name"
	LASTNAME    = "last_name"
	COMPANY     = "company"
	defaultHost = "aerospike"
	defaultPort = "3000"
	defaultNS   = "test"
	defaultSet  = "users"
)

type User struct {
	APIKEY    string
	FIRSTNAME string
	LASTNAME  string
	COMPANY   string
}
