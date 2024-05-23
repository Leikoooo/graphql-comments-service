package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	DefaultPageSize int
	MaxPageSize     int
	MaxReplyDepth   int
	MaxReplies      int
	StorageType     string
	Port            string
	DB_NAME         string
	DB_PORT         string
	DB_PASSWORD     string
	DB_HOST         string
	DB_USER         string
}

func SetConfig() AppConfig {
	err := godotenv.Load()

	defaultPageSize, err := strconv.Atoi(os.Getenv("DEFAULT_PAGE_SIZE"))
	if err != nil {
		log.Fatalf("Invalid value for DEFAULT_PAGE_SIZE: %v", err)
	}

	maxPageSize, err := strconv.Atoi(os.Getenv("MAX_PAGE_SIZE"))
	if err != nil {
		log.Fatalf("Invalid value for MAX_PAGE_SIZE: %v", err)
	}

	MAX_REPLY_DEPTH, err := strconv.Atoi(os.Getenv("MAX_REPLY_DEPTH"))
	if err != nil {
		log.Fatalf("Invalid value for MAX_REPLY_DEPTH: %v", err)
	}

	MAX_REPLIES, err := strconv.Atoi(os.Getenv("MAX_REPLIES"))
	if err != nil {
		log.Fatalf("Invalid value for MAX_REPLIES: %v", err)
	}

	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "" {
		log.Fatalf("STORAGE_TYPE environment variable is required")
	}

	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME == "" {
		log.Fatalf("DB_NAME environment variable is required")
	}
	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT == "" {
		log.Fatalf("DB_PORT environment variable is required")
	}
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		log.Fatalf("DB_PASSWORD environment variable is required")
	}
	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		log.Fatalf("DB_HOST environment variable is required")
	}
	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		log.Fatalf("DB_USER environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appConfig := AppConfig{
		DefaultPageSize: defaultPageSize,
		MaxPageSize:     maxPageSize,
		MaxReplies:      MAX_REPLIES,
		MaxReplyDepth:   MAX_REPLY_DEPTH,
		StorageType:     storageType,
		Port:            port,
		DB_HOST:         DB_HOST,
		DB_PASSWORD:     DB_PASSWORD,
		DB_PORT:         DB_PORT,
		DB_NAME:         DB_NAME,
		DB_USER:         DB_USER,
	}

	return appConfig
}
