package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn-api/component"
	"learn-api/component/uploadprovider"
	"learn-api/middleware"
	"learn-api/modules/upload/uploadtransport/ginupload"

	//"learn-api/middleware"
	"learn-api/modules/restaurant/restauranttransport/ginrestaurant"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	r.POST("/upload", ginupload.Upload(appCtx))

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
