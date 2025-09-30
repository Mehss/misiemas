package main

import (
	"log"
	"tripatra-dct-service-config/context"
	"tripatra-dct-service-config/database/config"
	"tripatra-dct-service-config/routes"

	projectRepository "tripatra-dct-service-config/database/repository/project"

	projectResolver "tripatra-dct-service-config/resolver/project"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitializeApp() *fiber.App {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	db := config.ConnectDB()

	// Run migrations
	if err := config.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Initialize repositories
	projectRepo := projectRepository.NewProjectRepository(db)

	// Initialize Resolvers
	projectRoutes := projectResolver.NewProjectResolver(projectRepo)

	// Initialize the Fiber app
	app := fiber.New()

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // Include both your local and deployed frontend
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: false, // Change to true if you're using cookies or auth tokens
	}))

	app.Use(context.DatabaseMiddleware(db))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Tap User Management Service</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						display: flex;
						justify-content: center;
						align-items: center;
						height: 100vh;
						margin: 0;
						background: #f4f4f4;
					}
					.container {
						text-align: center;
						background: white;
						padding: 20px;
						border-radius: 10px;
						box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
					}
					h1 {
						color: #333;
					}
					p {
						color: #666;
					}
					.button {
						display: inline-block;
						margin-top: 15px;
						padding: 10px 20px;
						text-decoration: none;
						color: white;
						background: #007BFF;
						border-radius: 5px;
						font-weight: bold;
						transition: background 0.3s;
					}
					.button:hover {
						background: #0056b3;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>🚀 TAP User Management Service API V1</h1>
					<p>Your API is running smoothly! You can now start using it.</p>
					<a href="https://your-api-docs.com" class="button">View API Docs</a>
				</div>
			</body>
			</html>
		`)
	})

	// Register routes API
	// var oauthConfig = &oauth2.Config{
	// 	ClientID:     os.Getenv("AZURE_AD_CLIENT_ID"),
	// 	ClientSecret: os.Getenv("AZURE_AD_CLIENT_SECRET"),
	// 	RedirectURL:  os.Getenv("AZURE_REDIRECT_URI"),
	// 	Scopes:       []string{"openid", "profile", "email"},
	// 	Endpoint:     microsoft.AzureADEndpoint(os.Getenv("AZURE_AD_TENANT_ID")),
	// }

	// routes.AuthorizationAzure(app, oauthConfig)

	// Create API groups with different middlewares
	api := app.Group("/api")
	// apiMicro := app.Group("/micro")

	// Routes accessible by frontend users
	api.Use(context.MiddlewareUser)
	routes.ProjectApi(api, projectRoutes)

	return app
}
