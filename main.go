package main

import (
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/mazay/dyndns-route53/internal/ipapi"
	"github.com/mazay/dyndns-route53/internal/route53"
)

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the defaultValue.
func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

// main is the entry point of the application. It initializes logging, retrieves
// necessary environment variables such as log level, AWS region, credentials,
// zone ID, and fully qualified domain name (FQDN). It then fetches the public
// IP address using the ipapi package, initializes a Route53 client, and updates
// the DNS record in Route53 with the new IP address. If any required environment
// variables are missing or if there are errors during execution, the application
// will log the error and terminate.
func main() {
	logLevel := getEnv("LOG_LEVEL", "info")
	logger := initLogger(logLevel)
	defer logger.Sync() //nolint:golint,errcheck

	dryRunStr := getEnv("DRY_RUN", "false")
	dryRun, err := strconv.ParseBool(dryRunStr)
	if err != nil {
		logger.Fatal("Invalid DRY_RUN value: " + err.Error())
	}
	logger.Info("DRY_RUN: " + strconv.FormatBool(dryRun))

	// get AWS region from env variable or set default
	region := getEnv("AWS_REGION", "us-east-1")
	logger.Info("AWS region: " + region)

	// get TTL and convert to int64
	ttl, err := strconv.ParseInt(getEnv("TTL", "60"), 10, 64)
	if err != nil {
		logger.Fatal("Invalid TTL value: " + err.Error())
	}
	logger.Info("TTL: " + strconv.FormatInt(ttl, 10))

	// get AWS credentials from env variables, they are optional as we can use other auth methods
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	zoneId, ok := os.LookupEnv("AWS_ZONE_ID")
	if !ok {
		logger.Fatal("AWS_ZONE_ID is not provided")
	}
	logger.Info("zone ID: " + zoneId)

	fdqn, ok := os.LookupEnv("FQDN")
	if !ok {
		logger.Fatal("FQDN is not provided")
	}
	logger.Info("FQDN to be updated: " + fdqn)

	logger.Info("fetching public IP address")
	ip, err := ipapi.GetIp()
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("current IP address: " + ip)

	oldIp, err := net.LookupIP(fdqn)
	if err != nil {
		logger.Info(err.Error())
	}

	if len(oldIp) == 0 {
		logger.Info("no old IP address found")
	}

	for _, _ip := range oldIp {
		if strings.Contains(_ip.String(), ip) {
			logger.Info("old IP address " + _ip.String() + " is the same as the new IP address " + ip)
			return
		}
	}

	r53 := route53.Route53{
		Region:          region,
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
	}

	logger.Info("initializing route53 client")
	r, err := r53.New()
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("updating route53 record with new IP address: " + ip)
	err = r.UpdateRRecord(zoneId, fdqn, ip, ttl, dryRun)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("done")
}
