package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
	"github.com/vchrisr/cf-static-broker/broker"
	"github.com/vchrisr/cf-static-broker/config"
	"github.com/vchrisr/cf-static-broker/static"
)

func loadCatalog() ([]brokerapi.Service, error) {
	var services []brokerapi.Service

	inBuf, err := ioutil.ReadFile("./catalog.json")
	if err != nil {
		return []brokerapi.Service{}, err
	}

	err = json.Unmarshal(inBuf, &services)
	if err != nil {
		return []brokerapi.Service{}, err
	}
	return services, nil
}

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	brokerCredentials := brokerapi.BrokerCredentials{
		Username: env.BrokerUsername,
		Password: env.BrokerPassword,
	}

	services, err := loadCatalog()
	if err != nil {
		panic(err)
	}

	logger := lager.NewLogger("static-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	static.LoadStaticStore("./staticCreds.json")

	serviceBroker := broker.New(services, logger, env)
	brokerAPI := brokerapi.New(serviceBroker, logger, brokerCredentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":"+env.Port, nil)
}
