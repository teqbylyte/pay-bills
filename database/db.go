package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

// Connect - Initialize database connection
func Connect() {
	var err error
	/*g := gen.NewGenerator(gen.Config{
		OutPath: "./database/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})*/

	Db, err = gorm.Open(mysql.Open(getDsn()), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Error", err)
	}

	fmt.Println("Database Connected")

	//g.UseDB(Db)

	//g.ApplyBasic(model)

	// Generate the code
	//g.Execute()

	//database.runMigration(db)
	//database.runSeeder()
}

// Get database config string from env
func getDsn() string {
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
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
