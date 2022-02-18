package server

import (
	"strconv"
	"github.com/my/repo/model"
	as "github.com/aerospike/aerospike-client-go/v5"
	utils "github.com/my/repo/utility"
)

var Connectaero model.Aerodb
var Client *as.Client

func LoadConnection() {
	Connectaero.Host = utils.GetEnv(utils.HOST, utils.DefaultHost)
	Connectaero.Port = utils.GetEnv(utils.PORT, utils.DefaultPort)
	Connectaero.Namespace = utils.GetEnv(utils.NAMESPACE, utils.DefaultNS)
	Connectaero.Set = utils.GetEnv(utils.SET, utils.DefaultSet)
	if !(Client != nil && Client.IsConnected()) {
		setClient()
	}
}

func setClient() {
	errDBDown := "Database is DOWN"
	_port, err := strconv.Atoi(Connectaero.Port)
	utils.PanicOnError(err)
	_client, err := as.NewClient(Connectaero.Host, _port)
	if err != nil {
		utils.LogMessage(errDBDown + " for " + Connectaero.Host + ":" + Connectaero.Port)
	}
	utils.LogMessage("Database is UP for " + Connectaero.Host + ":" + Connectaero.Port)
	
	Client = _client
}

func CheckDBConnection() {
	if !(Client.IsConnected()){
		setClient()
	}
}