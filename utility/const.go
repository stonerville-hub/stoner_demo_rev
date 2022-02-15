package utility

const (
	HOME         = "/"
	LOADNEWDATA  = "/loadNewData"
	GET_USERBYID = "/user/{id}"
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
	defaultHost = "0.0.0.0"
	defaultPort = "3000"
	defaultNS   = "test"
	defaultSet  = "users"
)

type User struct {
	APIKEY    string `json:"api_key"`
	COMPANY   string `json:"company"`
	FIRSTNAME string `json:"first_name"`
	LASTNAME  string `json:"last_name"`
}