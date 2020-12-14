package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/docs"
	_ "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/docs"
	"git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/domain/currency"
	"git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/domain/rate"
	currencyRepo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/currency"
	rateRepo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/rate"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
	"github.com/dimiro1/health"
	kitlog "github.com/go-kit/kit/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger lmgp-rates-svc API
// @version 1.0
// @description This is the documentation from lmgp-rates-svc service.
// termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://www.lifemiles.dev/support
// @contact.email jose.regalado@lifemiles.com

// license.name Apache 2.0
// license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath
func main() {

	// Creat read configuration dependency
	config := configuration.GetInstance(configuration.NewDefaultSettingWithoutVault())
	docs.SwaggerInfo.Host = "localhost:" + config.GetString("swagger-port")
	port := config.GetString("server.port")
	swaggerPort := config.GetString("swagger-port")
	// logger
	// Create log configuration dependency
	var kitlogger kitlog.Logger
	kitlogger = kitlog.NewJSONLogger(os.Stderr)
	kitlogger = kitlog.With(kitlogger, "time", kitlog.DefaultTimestamp)

	// Create log configuration dependency
	logger := log.NewLogger()
	defer logger.Sync()

	//create documentation handler
	swaggerHandler := httpSwagger.Handler(httpSwagger.URL("http://localhost:" + swaggerPort + "/swagger/doc.json"))

	// create HTTP Server
	mux := http.NewServeMux()

	// create health check handler
	handler := health.NewHandler()

	CurrencyPS := currencyRepo.NewCurrencyPS(config, logger)
	var serviceCurrency currency.Service
	serviceCurrency = currency.NewCurrencyService(CurrencyPS)
	serviceCurrency = currency.NewServiceLogger(logger, serviceCurrency)

	endpointCurrency := currency.Endpoints{
		Currency: currency.MakeGetCurrencyEndpoint(serviceCurrency),
	}

	mux.Handle("/getCurrencies", currency.NewHTTPCurrencyHandler(endpointCurrency))
	rateIfly := rateRepo.NewRateIfly(config, logger)
	var serviceRate rate.Service
	serviceRate = rate.NewMilesRateService(rateIfly)
	serviceRate = rate.NewServiceLogger(logger, serviceRate)

	endpoint := rate.Endpoints{
		MilesRate: rate.MakeGetMilesRatesEndpoint(serviceRate),
	}
	mux.Handle("/milesRate", rate.NewHTTPHandler(endpoint))
	mux.Handle("/health", handler)
	mux.Handle("/swagger/", swaggerHandler)

	errs := make(chan error, 2)
	go func() {
		//currencyRepo.NewCurrencyPS(config, logger).GetCurrency()
		//rateRepo.NewRateIfly(config, logger).GetRatePerStore("TOTCO", "0010077923")
		kitlogger.Log("transport", "http", "address", port, "msg", "listening")
		errs <- http.ListenAndServe(":"+port, mux)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		logger.Sync()
	}()

	kitlogger.Log("terminated", <-errs)

}
