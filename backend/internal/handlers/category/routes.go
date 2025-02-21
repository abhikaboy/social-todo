package Category

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Router maps endpoints to handlers
*/
func Routes(app *fiber.App, collections map[string]*mongo.Collection) {
	service := newService(collections)
	handler := Handler{service}

	// Add a group for API versioning
	apiV1 := app.Group("/api/v1")

	// Add Sample group under API Version 1
	Categories := apiV1.Group("/Categories")

	Categories.Post("/", handler.CreateCategory)
	Categories.Get("/", handler.GetCategories)
	
	Categories.Delete("/user/:user/:id", handler.DeleteCategory)
	Categories.Patch("/user/:user/:id", handler.UpdatePartialCategory)
	Categories.Get("/user/:id", handler.GetCategoriesByUser)

}
