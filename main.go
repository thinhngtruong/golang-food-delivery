package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn-api/component"
	"learn-api/middleware"

	//"learn-api/middleware"
	"learn-api/modules/restaurant/restauranttransport/ginrestaurant"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.DeleteRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run()
}

//CREATE TABLE `restaurants` (
//	`id` int(11) NOT NULL AUTO_INCREMENT,
//	`owner_id` int(11) NOT NULL,
//	`name` varchar(50) NOT NULL,
//	`addr` varchar(255) NOT NULL,
//	`city_id` int(11) DEFAULT NULL,
//	`lat` double DEFAULT NULL,
//	`lng` double DEFAULT NULL,
//	`cover` json NOT NULL,
//	`logo` json NOT NULL,
//	`shipping_fee_per_km` double DEFAULT '0',
//	`status` int(11) NOT NULL DEFAULT '1',
//	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//	PRIMARY KEY (`id`),
//	KEY `owner_id` (`owner_id`) USING BTREE,
//	KEY `city_id` (`city_id`) USING BTREE,
//	KEY `status` (`status`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;
