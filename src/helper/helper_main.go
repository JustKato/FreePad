package helper

import (
	"os"
	"strconv"
)

func GetDomainBase() string {
	domainBase, domainExists := os.LookupEnv("DOMAIN_BASE")

	if !domainExists {
		os.Setenv("DOMAIN_BASE", "http://localhost:8080")
		domainBase = "http://localhost:8080"
	}

	return domainBase
}

func GetCacheMapLimit() int {
	cacheMapLimit, domainExists := os.LookupEnv("CACHE_MAP_LIMIT")

	if !domainExists {
		os.Setenv("CACHE_MAP_LIMIT", "25")
		cacheMapLimit = "25"
	}

	rez, err := strconv.Atoi(cacheMapLimit)
	if err != nil {
		return 25
	}

	return rez
}
