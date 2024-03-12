package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"martpay/app/models"
	"martpay/database/query"
)

var Db *gorm.DB

// Connect - Configure and open the database connection
func Connect() {
	var err error

	g := gen.NewGenerator(gen.Config{
		OutPath: "./database/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	Db, err = gorm.Open(postgres.Open(postgresDsn()), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Error", err)
	}

	fmt.Println("Database Connected")

	g.UseDB(Db)

	g.ApplyBasic(&models.Fee{},
		&models.GeneralLedger{},
		&models.GlTransaction{},
		&models.KycLevel{},
		&models.Loan{},
		&models.Service{},
		&models.ServiceProvider{},
		&models.Terminal{},
		&models.TerminalGroup{},
		&models.Transactions{},
		&models.User{},
		&models.Wallet{},
		&models.WalletTransaction{},
	)

	query.SetDefault(Db)

	// Generate the code
	//g.Execute()
}

func mysqlDsn() string {
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
}

func postgresDsn() string {
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")
	ssl := viper.GetString("DB_SSL")
	timezone := viper.GetString("DB_TIMEZONE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, ssl, timezone)
}

/*func (database Database) runMigration(gormDB *gorm.DB) {
	_ = gormDB.AutoMigrate(&model2.User{})
	_ = gormDB.AutoMigrate(&model2.Wallet{})
	_ = gormDB.AutoMigrate(&model2.WalletTransaction{})
	_ = gormDB.AutoMigrate(&model2.Transactions{})
	_ = gormDB.AutoMigrate(&model2.VirtualAccounts{})
	_ = gormDB.AutoMigrate(&model2.Kyc{})
	_ = gormDB.AutoMigrate(&model2.KycDoc{})
	_ = gormDB.AutoMigrate(&model2.Fee{})
	_ = gormDB.AutoMigrate(&model2.Commission{})
	_ = gormDB.AutoMigrate(&model2.UpgradeRequest{})
	_ = gormDB.AutoMigrate(&model2.Beneficiary{})
	_ = gormDB.AutoMigrate(&model2.SchedulePayment{})
}*/

func runSeeder() {
	//seeder.SeedTable(&seeder.KycSeeder{})
	//seeder.SeedTable(&seeder.FeeSeeder{})
	//seeder.SeedTable(&seeder.CommissionSeeder{})
}
