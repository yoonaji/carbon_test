package migrate

import (
	"fmt"

	"github.com/yoonaji/carbon_test/initializers"
	"github.com/yoonaji/carbon_test/models"
)

func Migrate() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.TransactionModel{}, &models.WebhookTransaction{}, &models.User{})
	fmt.Println("👍 Migration complete")
}
