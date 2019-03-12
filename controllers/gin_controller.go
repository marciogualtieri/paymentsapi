package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marciogualtieri/paymentsapi/configuration"
	"github.com/marciogualtieri/paymentsapi/errors"
	"github.com/marciogualtieri/paymentsapi/repositories"
	"github.com/marciogualtieri/paymentsapi/serializers"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/*
GinController is a controller implemented using the Gin framework.
*/
type GinController struct {
	JSONSerializer    serializers.JSONSerializer
	PaymentRepository repositories.PaymentsRepository
	Configuration     configuration.Configuration
}

/*
NewGinController creates a new Gin controller.
*/
func NewGinController(
	serializer serializers.JSONSerializer,
	repository repositories.PaymentsRepository,
	config configuration.Configuration) *GinController {
	return &GinController{
		JSONSerializer:    serializer,
		PaymentRepository: repository,
		Configuration:     config,
	}
}

func setupLogging(configuration configuration.Configuration) {
	requestsLogFile, _ := os.OpenFile(configuration.RequestsLogFile,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errorsLogFile, _ := os.OpenFile(configuration.ErrorsLogFile,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(requestsLogFile)
	gin.DefaultErrorWriter = io.MultiWriter(errorsLogFile)
}

func handleError(err error, context *gin.Context) {
	if err != nil {
		errorJSON := gin.H{"error": err.Error()}
		errorStatus := http.StatusInternalServerError
		switch err.(type) {
		case *errors.ErrRepositoryRecordNotFound:
			errorStatus = http.StatusNotFound
		case *errors.ErrRepository:
		case *errors.ErrParsingJSON:
			errorStatus = http.StatusBadRequest
		}
		context.AbortWithStatusJSON(errorStatus, errorJSON)
	}
}

func handlerForFetch(repository repositories.PaymentsRepository) func(c *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		payment, err := repository.Read(id)
		if err != nil {
			handleError(err, context)
		} else {
			context.JSON(http.StatusOK, payment)
		}
	}
}

func handlerForList(repository repositories.PaymentsRepository) func(c *gin.Context) {
	return func(context *gin.Context) {
		payments, err := repository.ReadAll()
		if err != nil {
			handleError(err, context)
		} else {
			context.JSON(http.StatusOK, payments)
		}
	}
}

func handlerForCreate(repository repositories.PaymentsRepository,
	serializer serializers.JSONSerializer) func(c *gin.Context) {
	return func(context *gin.Context) {
		body, _ := ioutil.ReadAll(context.Request.Body)
		payment, err := serializer.ParsePayment(string(body))
		if err != nil {
			handleError(err, context)
		} else {
			id, err := repository.Create(*payment)
			if err != nil {
				handleError(err, context)
			} else {
				context.JSON(http.StatusCreated, gin.H{"ID": id})
			}
		}
	}
}

func handlerForUpdate(paymentsRepository repositories.PaymentsRepository,
	serializer serializers.JSONSerializer) func(c *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		body, _ := ioutil.ReadAll(context.Request.Body)
		payment, err := serializer.ParsePayment(string(body))
		if err != nil {
			handleError(err, context)
		} else {
			payment.ID = id
			err = paymentsRepository.Update(*payment)
			if err != nil {
				handleError(err, context)
			} else {
				context.Status(http.StatusOK)
			}
		}
	}
}

func handlerForDelete(repository repositories.PaymentsRepository) func(c *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		err := repository.Delete(id)
		if err != nil {
			handleError(err, context)
		} else {
			context.Status(http.StatusOK)
		}
	}
}

/*
NewGinEngine creates a new Gin engine.
*/
func NewGinEngine(repository repositories.PaymentsRepository,
	serializer serializers.JSONSerializer,
	config configuration.Configuration) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	setupLogging(config)
	engine := gin.Default()
	engine.GET(fmt.Sprintf("%s/:id", config.BaseResource), handlerForFetch(repository))
	engine.POST(config.BaseResource, handlerForCreate(repository, serializer))
	engine.PUT(fmt.Sprintf("%s/:id", config.BaseResource), handlerForUpdate(repository, serializer))
	engine.DELETE(fmt.Sprintf("%s/:id", config.BaseResource), handlerForDelete(repository))
	engine.GET(config.BaseResource, handlerForList(repository))
	return engine
}

/*
Run starts the controller.
*/
func (controller *GinController) Run() {
	ginEngine := NewGinEngine(
		controller.PaymentRepository,
		controller.JSONSerializer,
		controller.Configuration)
	defer controller.PaymentRepository.Close()
	ginEngine.Run()
}
