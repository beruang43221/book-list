package router

import (
	"log"
	"os"

	"github.com/beruang43221/book-list/internal/app/controller"
	"github.com/beruang43221/book-list/internal/app/database"
	"github.com/beruang43221/book-list/internal/app/repository/bookrepo/bookpg"
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

	bookRepo := bookpg.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo, categoryRepo)
	bookController := controller.NewBookController(bookService)

	categoryRouter := router.Group("/categories")
	{
		categoryRouter.POST("/", categoryController.CreateCategory)
		categoryRouter.GET("/", categoryController.GetAllCategories)
		categoryRouter.GET("/:id", categoryController.GetCategoriesbyID)
		categoryRouter.PUT("/:id", categoryController.UpdateCategory)
		categoryRouter.DELETE("/:id", categoryController.DeleteCategory)
	}

	bookRouter := router.Group("/books")
	{
		bookRouter.POST("/", bookController.CreateBook)
		bookRouter.GET("/", bookController.GetAllBooks)
		bookRouter.PUT("/:id", bookController.UpdateBook) // buat jika category gk ada itu error
		bookRouter.DELETE("/:id", bookController.DeleteBook)
		bookRouter.GET("/category/:category_id", bookController.GetBooksByCategories) // ganti filter dan benerin error
		bookRouter.GET("", bookController.FilterBooksBySearchText)
		bookRouter.GET("/filter", bookController.GetBooksByDate) // penanganann error msh acakadul
	}

	var PORT = os.Getenv("PORT")
	router.Run(":" + PORT)
}
