package main

import (
	"github.com/marciogualtieri/paymentsapi/configuration"
	"github.com/marciogualtieri/paymentsapi/controllers"
	"github.com/marciogualtieri/paymentsapi/repositories"
	"github.com/marciogualtieri/paymentsapi/serializers"
)

func main() {
	config := configuration.GetConfiguration()
	repository := repositories.NewSqlitePaymentRepository(config)
	serializer := serializers.NewDefaultJSONSerializer()
	controller := controllers.NewGinController(serializer, repository, config)
	controller.Run()
}
