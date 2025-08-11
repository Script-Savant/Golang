/*
bootstraps config, database, router and server
- load env variables into typed con
- initialize GORM DB connection - mysql
- run migrationsand seed sample products
- Create Gin router and register routes
- start http server on configured port
*/
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"backend/config"
)

// main - initializes everything and starts the HTTP server
func main() {
	/*
		1. Load Configuration from env
		2. Open DB connection via internal/db
		3. Run migrations and seed sample data
		4. Initialize Gin router and register routes
		5. Start server on configred port
	*/

	cfg, err := configLoad()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	db, err := openDatabase(cfg)
	if err != nil {
		log.Fatalf("db open error: %v", err)
	}

	if err := migrateAndSeed(db); err != nil {
		log.Fatalf("migrate/seed error: %v", err)
	}

	r := setupRouter(db, cfg)
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func configLoad() (*Config, error) {
	/*
		1. Call config.LoadConfigFromEnv
		2. Return Config or an error
	*/
	return config.LoadConfigFromEnv()
}

func openDatabase(cfg *Config) (*gorm.DB, error) {
	/*
		1. Call db.OpenDatabase(cfg)
		2. return *gorm.DB or error
	*/
	return OpenDatabase(cfg)
}

func migrateAndSeed(db *gorm.DB) error {
	/*
		1. Call internal/db.Migrate(db)
		2. Call internal/db.SeedSampleProducts(db)
		3. Return any errors
	*/
	return MigrateAndSeed(db)
}

func setupRouter(db *gorm.DB, cfg *Config) *gin.Engine {
	// Steps:
	// 1. Create gin.Default()
	// 2. Call internal/routes.RegisterRoutes
	// 3. Return configured engine
	r := gin.Default()
	RegisterRoutes(r, db, cfg)
	return r
}
