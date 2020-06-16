package seed

import (
	"log"

	"github.com/Progete-Dev/go-whatsapp-rest/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Email:    "exatasmente@email.com",
		Password: "secret",
	},
	models.User{
		Email:    "admin@root.com",
		Password: "root",
	},
}
var userTokens = []models.UserToken{
	models.UserToken{
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjAsInVzZXJfaWQiOjJ9.sGDXEoKbCzWMHa9m-DPuC_BvUg8JgqqnQkVv2AQOzHI",
		UserId:     2,
		WebhookUrl: "http://psi-laravel.herokuapp.com/botman",
	},
}

func Load(db *gorm.DB) {

	// err := db.Debug().DropTableIfExists(&models.User{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	// err = db.Debug().AutoMigrate(&models.User{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot migrate table: %v", err)
	// }
	// err = db.Debug().DropTableIfExists(&models.WpSession{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	// err = db.Debug().AutoMigrate(&models.WpSession{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot migrate table: %v", err)
	// }
	// err = db.Debug().DropTableIfExists(&models.UserToken{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	// err = db.Debug().AutoMigrate(&models.UserToken{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot migrate table: %v", err)
	// }

	// for i, _ := range users {
	// 	err := db.Debug().Model(&models.User{}).Create(&users[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed users table: %v", err)
	// 	}

	// }
	for i, _ := range userTokens {
		err := db.Debug().Model(&models.UserToken{}).Create(&userTokens[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}
