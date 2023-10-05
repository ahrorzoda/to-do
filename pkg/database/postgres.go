package database

import (
	"fmt"
	"github.com/ahrorzoda/to-do/pkg/config"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

/*type Database interface {
	Connection(config *config.Postgres) (*gorm.DB, error)
	Close() error
}

type Postgres struct {
	DB *gorm.DB
}*/

func /*(db *Postgres)*/ Connection(config *config.Postgres) (*gorm.DB, error) {
	err := godotenv.Load("../password.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		fmt.Println("Ошибка: Пароль не указан")
	}
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.Host, config.Port, config.DBName, config.User, password, config.SSLMode)
	log.Println("dsn ==> ", connectionString)
	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println("Error in DSN ==> ", err.Error())
		return nil, err
	}

	/*if !conn.DisableAutomaticPing {
		log.Printf("%s in Ping connect", err.Error())
	}*/

	//db.DB = conn
	log.Println("Connected to Postgres!")
	return conn, nil
}

/*func (d *Postgres) Close() error {
	db, err := d.DB.DB()
	if err != nil {
		return err
	}
	log.Println("Postgres is disconnected!")
	return db.Close()
}
*/
