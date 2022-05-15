package helper

import "os"

func GetDomainBase() string {
	domainBase, domainExists := os.LookupEnv("DOMAIN_BASE")

	if !domainExists {
		os.Setenv("DOMAIN_BASE", "http://localhost:8080")
		domainBase = "http://localhost:8080"
	}

	return domainBase
}
