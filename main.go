package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/maropook/gopher-slayer/config"
	"github.com/maropook/gopher-slayer/internal/handler"
	"github.com/maropook/gopher-slayer/internal/repository"
	"github.com/maropook/gopher-slayer/internal/service"
)

func main() {
	// Load environment variables from .env file (ignored if not present)
	godotenv.Load()

	cfg := config.Load()

	// Connect to MySQL with retry (Docker may start the app before DB is ready)
	db := connectDB(cfg)
	defer db.Close()

	// ---- Dependency injection (manual wiring) ----
	// repository layer
	heroRepo := repository.NewHeroRepository(db)
	stageRepo := repository.NewStageRepository(db)

	// service layer
	heroService := service.NewHeroService(heroRepo)
	stageService := service.NewStageService(stageRepo, heroRepo)
	battleService := service.NewBattleService()

	// handler layer
	heroHandler := handler.NewHeroHandler(heroService)
	stageHandler := handler.NewStageHandler(stageService)
	battleHandler := handler.NewBattleHandler(battleService)

	// ---- Echo setup ----
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	// Serve frontend files
	e.Static("/", "frontend")

	// Serve images directory
	e.Static("/images", "images")

	// Serve Swagger spec
	e.File("/docs/swagger.yaml", "docs/swagger.yaml")

	// ---- API routes ----
	api := e.Group("/api")

	// Hero
	api.GET("/hero", heroHandler.GetHero)
	api.PUT("/hero/name", heroHandler.UpdateName)
	api.PUT("/hero/experience", heroHandler.UpdateExperience)
	// [Lv3 workshop task - bug plant location]
	// In the bug version, the line below is commented out → 404 error.
	// Students must add this route registration to fix the game.
	api.PUT("/hero/hp", heroHandler.UpdateHP)

	// Stage
	api.GET("/stages", stageHandler.GetStages)
	api.GET("/stages/:id/enemies", stageHandler.GetEnemies)
	api.POST("/stages/:id/clear", stageHandler.ClearStage)

	// Battle
	api.POST("/battle/attack", battleHandler.Attack)
	api.POST("/battle/enemy-attack", battleHandler.EnemyAttack)

	log.Printf("Server starting on :%s", cfg.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}

// connectDB opens a MySQL connection with retry to handle Docker startup timing.
func connectDB(cfg *config.Config) *sql.DB {
	var db *sql.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = sql.Open("mysql", cfg.DSN())
		if err == nil {
			if pingErr := db.Ping(); pingErr == nil {
				log.Println("Connected to database")
				return db
			}
		}
		log.Printf("Database not ready, retrying... (%d/10)", i)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Failed to connect to database after 10 attempts")
	return nil
}
