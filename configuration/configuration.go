package configuration

import (
	"github.com/magiconair/properties"
	"github.com/marciogualtieri/paymentsapi/utils"
	"path"
)

/*
Configuration holds the properties in ./resources/configuration.properties.
*/
type Configuration struct {
	RequestsLogFile    string `properties:"requests.log.file,default=requests.log"`
	ErrorsLogFile      string `properties:"errors.log.file,default=errors.log"`
	RepositoryLogFile  string `properties:"repository.log.file,default=repository.log"`
	BaseResource       string `properties:"base.resource,default=/paymentsapi"`
	SqliteDatabaseFile string `properties:"sqlite.database.file,default=sqlite.db"`
	SqliteLogMode      bool   `properties:"sqlite.log.mode,default=false"`
}

func getConfiguration(configurationFile string) Configuration {
	var config Configuration
	props := properties.MustLoadFile(path.Join(utils.GetProjectPath(),
		configurationFile), properties.UTF8)
	if err := props.Decode(&config); err != nil {
		panic(err)
	}
	return config
}

/*
GetConfiguration returns the configuration parameters defined in
  "resources/configuration.properties" for production.
*/
func GetConfiguration() Configuration {
	return getConfiguration("resources/configuration.properties")
}

/*
GetTestConfiguration returns the configuration parameters defined in
  "resources/configuration.properties" for tests.
*/
func GetTestConfiguration() Configuration {
	return getConfiguration("testing/resources/configuration.properties")
}
