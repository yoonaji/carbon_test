package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=localhost user=postgres dbname=transaction sslmode=disable password=0000 port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB 연결 실패: " + err.Error())
	}

	// AutoMigrate 추가
	err = DB.AutoMigrate(&transaction.TransactionModel{})
	if err != nil {
		panic("DB 마이그레이션 실패: " + err.Error())
	}
}
