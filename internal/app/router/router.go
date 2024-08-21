package router

import (
	"log"
	"os"

	"github.com/beruang43221/book-list/internal/app/controller"
	"github.com/beruang43221/book-list/internal/app/database"
	"github.com/beruang43221/book-list/internal/app/repository/categoryrepo/categorypg"
	"github.com/beruang43221/book-list/internal/app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartServer() {
	db, err = database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	categoryRepo := categorypg.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	categoryRouter := router.Group("/categories")
	{
		categoryRouter.POST("/", categoryController.CreateCategory)
		categoryRouter.GET("/", categoryController.GetAllCategories)
		categoryRouter.GET("/:id", categoryController.GetCategoriesbyID)
		// categoryRouter.PUT("/:id", controller.UpdateCategory)
		// categoryRouter.DELETE("/:id", controller.DeleteCategory)
	}

	// bookRouter := router.Group("/books")
	// {
	// 	bookRouter.POST("/", controller.CreateBook)
	// 	bookRouter.GET("", controller.FilterBooksBySearchText)
	// 	bookRouter.GET("/", controller.GetAllBooks)
	// 	bookRouter.GET("/filter", controller.GetBooksByDate)
	// 	bookRouter.GET("/category/:category_id", controller.GetBooksByCategories)
	// 	bookRouter.PUT("/:id", controller.UpdateBook)
	// 	bookRouter.DELETE("/:id", controller.DeleteBook)
	// }

	var PORT = os.Getenv("PORT")
	router.Run(":" + PORT)
}
