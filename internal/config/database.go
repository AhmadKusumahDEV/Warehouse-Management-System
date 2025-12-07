package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DBConfig holds database configuration parameters
type DBConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	SSLMode         string
}

// DefaultDBConfig returns a DBConfig with default values
func DefaultDBConfig() DBConfig {
	return DBConfig{
		MaxIdleConns:    5,
		MaxOpenConns:    20,
		ConnMaxLifetime: 60 * time.Minute,
		ConnMaxIdleTime: 10 * time.Minute,
		SSLMode:         "disable", // Use "require" or "verify-full" in production
	}
}

// NewDB creates a new database connection with enhanced configuration
func NewDB() *sql.DB {
	// Load environment variables
	config, err := loadDBConfig()

	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Set connection pool settings
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}

// loadDBConfig loads database configuration from environment variables
func loadDBConfig() (DBConfig, error) {
	config := DefaultDBConfig()

	// Try to load .env file, but don't fail if it doesn't exist
	// This allows environment variables to be set directly (e.g., in production)
	err := godotenv.Load("configs/local.env")
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
	}

	// Get required environment variables
	config.Host = getEnv("HOST", "localhost")
	config.Port = getEnv("PORT", "5432")
	config.User = getEnv("USER", "postgres")
	config.Password = getEnv("PASSWORD", "postgres")
	config.DBName = getEnv("NAME", "postgres")

	fmt.Println(config.Password)
	// Get optional environment variables with defaults
	if maxIdle := getEnv("DB_MAX_IDLE_CONNS", ""); maxIdle != "" {
		if val, err := strconv.Atoi(maxIdle); err == nil {
			config.MaxIdleConns = val
		}
	}

	if maxOpen := getEnv("DB_MAX_OPEN_CONNS", ""); maxOpen != "" {
		if val, err := strconv.Atoi(maxOpen); err == nil {
			config.MaxOpenConns = val
		}
	}

	if sslMode := getEnv("DB_SSL_MODE", ""); sslMode != "" {
		config.SSLMode = sslMode
	}

	// Validate required configuration
	if config.User == "" || config.Password == "" || config.DBName == "" {
		return config, fmt.Errorf("database user, password, and name are required")
	}
	return config, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
