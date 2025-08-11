/*
Manage configuration - env variables
1. read env
2. expose Config struct: DSN, JWT_SECRET, PORT, BCRYPT_CONST
*/

package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config - holds app conf loaded from env
type Config struct {
	DSN string
	JWT_SECRET string
	Port string
	BCost int
	TokenTTLMin int
}

// LoadConfigFromEnv reads environment variables and returns a Config.
func LoadConfigFromEnv() (*Config, error) {
	/*
	1. read env variables
	2. parse values slike bcrypt cost
	3. validate and return config or error
	*/
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("failed to load env variables")
	}

	dsn := os.Getenv("DSN")
	jwt := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")
	bcosts := os.Getenv("BCRYPT_COST")
	ttl := os.Getenv("TOKEN_TTL_MIN")

	bc := 12 
	if bcosts != "" {
		if v, err := strconv.Atoi(bcosts); err == nil {
			bc = v
		}
	}

	ttlmin := 15
	if ttl != "" {
		if v, err := strconv.Atoi(ttl); err == nil {
			ttlmin = v
		}
	}

	return &Config{
		DSN: dsn,
		JWT_SECRET: jwt,
		Port: port,
		BCost: bc,
		TokenTTLMin: ttlmin,
	}, nil
}