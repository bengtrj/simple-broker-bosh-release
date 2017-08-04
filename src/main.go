package main

import (
	"log"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"

	"github.com/pivotal-cf/brokerapi"
)

func main() {
	serviceBroker := &FakeServiceBroker{}

	// result, err := serviceBroker.Bind(nil, "nil", "nil", details)

	// if err == nil {
	// 	fmt.Printf(result.SyslogDrainURL)
	// }

	logger := lager.NewLogger("Log")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	handler := brokerapi.New(serviceBroker, logger, brokerapi.BrokerCredentials{"user", "password"})
	log.Fatal(http.ListenAndServe(":8080", handler))

}

//service_broker
