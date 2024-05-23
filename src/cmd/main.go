package main

import (
	"OzonTest/config"
	"OzonTest/src/Init"
	"OzonTest/src/internal/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	appConfig := config.SetConfig()

	router := mux.NewRouter()

	// Создание сервисов
	postService, commentService, err := Init.InitServices(appConfig.StorageType, appConfig)
	if err != nil {
		log.Fatalf("Error initializing services: %v", err)
	}

	// Установка роутеров
	err = Init.SetupRoutes(router, postService, commentService, appConfig)
	if err != nil {
		log.Fatalf("Error while setup routes: %v", err)
	}

	// Функция восстановления, чтобы приложение не падало
	router.Use(utils.Recovery)
	log.Printf("App Starting on port %s", appConfig.Port)
	log.Fatal(http.ListenAndServe(":"+appConfig.Port, router))
}
