package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Using sqllite at the moment
	"github.com/marciogualtieri/paymentsapi/configuration"
	"github.com/marciogualtieri/paymentsapi/errors"
	"github.com/marciogualtieri/paymentsapi/models"
	"github.com/satori/go.uuid"
	"log"
	"os"
)

func copyAttributesRelatedData(repo *SqlitePaymentRepository, payment *models.Payment) {
	repo.DB.
		Preload("BeneficiaryParty").
		Preload("DebtorParty").
		Preload("SponsorParty").
		Preload("Fx").
		Find(&payment.Attributes)
	repo.DB.Preload("SenderCharges").Find(&payment.Attributes.ChargesInformation)
}

/*
SqlitePaymentRepository is the persistence layer implemented using GORM.
*/
type SqlitePaymentRepository struct {
	DB            *gorm.DB
	Configuration configuration.Configuration
}

func setupLogging(db *gorm.DB, logMode bool, logFile string) {
	db.LogMode(logMode)
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	db.SetLogger(log.New(file, "\r\n", 0))
}

/*
NewSqlitePaymentRepository creates a new sqlite payment repository.
*/
func NewSqlitePaymentRepository(configuration configuration.Configuration) *SqlitePaymentRepository {
	db, _ := gorm.Open("sqlite3", configuration.SqliteDatabaseFile)
	setupLogging(db, configuration.SqliteLogMode,
		configuration.RepositoryLogFile)
	db.AutoMigrate(&models.Payment{},
		&models.BeneficiaryParty{},
		&models.DebtorParty{},
		&models.SponsorParty{},
		&models.Attributes{},
		&models.SenderCharges{},
		&models.ChargesInformation{},
		&models.Fx{})
	return &SqlitePaymentRepository{DB: db, Configuration: configuration}
}

/*
Close closes the database.
*/
func (repo *SqlitePaymentRepository) Close() {
	repo.DB.Close()
}

/*
Create persists a payment to the data store.
*/
func (repo *SqlitePaymentRepository) Create(payment models.Payment) (*string, error) {
	uuid, _ := uuid.NewV4()
	payment.ID = uuid.String()
	err := repo.DB.Save(&payment).Error
	if err != nil {
		return nil, errors.NewErrRepository(err.Error())
	}
	return &payment.ID, nil
}

/*
Read fetchs a payment from the data store.
*/
func (repo *SqlitePaymentRepository) Read(id string) (*models.Payment, error) {
	var payment models.Payment
	err := repo.DB.Preload("Attributes").Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, errors.NewErrRepository(err.Error())
	}
	copyAttributesRelatedData(repo, &payment)
	return &payment, nil
}

/*
Update updates a payment in the data store.
*/
func (repo *SqlitePaymentRepository) Update(payment models.Payment) error {
	err := repo.DB.Save(&payment).Error
	if err != nil {
		return errors.NewErrRepository(err.Error())
	}
	return nil
}

/*
Delete deletes a payment from the data store.
*/
func (repo *SqlitePaymentRepository) Delete(id string) error {
	var payment models.Payment
	err := repo.DB.Where("id = ?", id).First(&payment).Error
	if err != nil {
		return errors.NewErrRepository(err.Error())
	}
	err = repo.DB.Delete(&payment).Error
	if err != nil {
		return errors.NewErrRepository(err.Error())
	}
	return nil
}

/*
ReadAll lists all payments in the data store.
*/
func (repo *SqlitePaymentRepository) ReadAll() ([]models.Payment, error) {
	var payments []models.Payment
	err := repo.DB.Preload("Attributes").Find(&payments).Error
	if err != nil {
		return nil, errors.NewErrRepository(err.Error())
	}
	for i := range payments {
		copyAttributesRelatedData(repo, &payments[i])
	}
	return payments, nil
}
