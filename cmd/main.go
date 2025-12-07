package main

import (
	_ "github.com/AhmadKusumahDEV/Warehouse-Management-System/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Warehouse Management API
// @version         1.0
// @description     API untuk sistem manajemen warehouse (WMS).
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@anda.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	// 4. Buat router gin.Default()
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(201, "jawa jawa jawa")
	})
	// 5. Tambahkan route untuk Swagger UI
	// Ini adalah route yang Anda minta
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// (Perhatikan: kita TIDAK perlu mendaftarkan route API
	//  seperti /api/v1/employees di sini.
	//  swag init membaca anotasi dari SEMUA file .go Anda,
	//  bukan hanya dari main.go)

	// 6. Jalankan server
	r.Run(":8080")
}
