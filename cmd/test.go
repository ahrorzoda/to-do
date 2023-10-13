package main

import (
	"github.com/ahrorzoda/to-do/pkg/config"
	"github.com/ahrorzoda/to-do/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
)

type Classroom struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Number   string `gorm:"varchar" json:"number"`
	Busy     bool   `gorm:"type:boolean" json:"busy"`
	UniverID uint   // Внешний ключ для связи с Univer
}

type Univer struct {
	ID      uint        `gorm:"primary_key" json:"id"`
	Address string      `gorm:"varchar" json:"address"`
	Class   []Classroom `gorm:"foreignKey:UniverID" json:"class"`
}

func Add(c *gin.Context) {
	var un Univer

	if err := c.BindJSON(&un); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("univer ==> ", un)
	db, err := database.Connection(&config.InitJsonConfigs().Postgres)
	if err != nil {
		log.Println("Error in database connection ==> ", err.Error())
		return
	}

	if err := db.Create(&un).Error; err != nil {
		log.Println("Error in create ==> ", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"result": un,
	})
}

func main() {
	rout := gin.Default()
	rout.POST("/add", Add)
	if err := rout.Run(); err != nil {
		log.Println("Error in run project ==> ", err.Error())
		return
	}
}
