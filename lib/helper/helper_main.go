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

func GetMaximumPadSize() int {
	// Lookup if the maximum pad size variable exists.
	maxPadSize, exists := os.LookupEnv("MAXIMUM_PAD_SIZE")

	// Check if this environment variable has bee nset
	if !exists {
		// Set the variable ourselves to the default string value
		maxPadSize = "524288"
	}

	// Try and convert the string into an integer
	rez, err := strconv.Atoi(maxPadSize)
	// Check if the conversion has failed
	if err != nil {
		// Simply return the default
		return 524288
	}

	// Return the resulting value
	return rez
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
