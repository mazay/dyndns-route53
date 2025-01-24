# Use case

A tiny app that helps with using AWS Route 53 as dynamic DNS service.

## Usage

Set required variables:
```bash
export AWS_ACCESS_KEY_ID=******
export AWS_SECRET_ACCESS_KEY=************
export ZONE_ID=EXAMPLEZONE
export FQDN=test.example.com
```

Run the DNS updater:
```bash
docker run \
  -e "AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}" \
  -e "AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}" \
  -e "ZONE_ID=${ZONE_ID}" \
  -e "FQDN=${FQDN}" \
  ghcr.io/mazay/dyndns-route53:main
```

Alternatively there are prebuilt binaries for various OS/arch sets.

## How it works

The application will:

1. get your current IP using [the IP Geolocation API](https://ip-api.com/)
1. try to resolve the `FQDN` and compare the result with the current IP from step 1
1. if `FQDN` resolve has failed or IPs missmatch the `FQDN` record will be created/updated
