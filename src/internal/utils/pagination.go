package utils

import (
	"OzonTest/config"
	"net/http"
	"strconv"
)

func GetPage(r *http.Request) int {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(r *http.Request, appConfig config.AppConfig) int {
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil || pageSize <= 0 {
		return appConfig.DefaultPageSize
	}
	if pageSize > appConfig.MaxPageSize {
		return appConfig.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	if page <= 0 {
		return 0
	}
	return (page - 1) * pageSize
}
