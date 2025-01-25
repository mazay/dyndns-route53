# Use case

A tiny app that helps with using AWS Route 53 as dynamic DNS service.

## Usage

Set required variables:

```bash
export AWS_ACCESS_KEY_ID=****** # may be skipped if any other auth method is available
export AWS_SECRET_ACCESS_KEY=************ # may be skipped if any other auth method is available
export AWS_ZONE_ID=EXAMPLEZONE
export FQDN=test.example.com
```

Run the DNS updater:

```bash
docker run \
  -e "AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}" \
  -e "AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}" \
  -e "AWS_ZONE_ID=${AWS_ZONE_ID}" \
  -e "FQDN=${FQDN}" \
  ghcr.io/mazay/dyndns-route53:main
```

Alternatively there are prebuilt binaries for various OS/arch sets.

Optional setting/variables:

- `LOG_LEVEL`, valid options - `debug`, `info`, `warn`, `error`, `fatal`, `panic`. Defaults to `info`.
- `AWS_REGION`, AWS region the `AWS_ZONE_ID` exists in. Defaults to `us-east-1`.
- `TTL`, the TTL of the DNS record in seconds. AWS Recommended values: `60` to `172800`. Defaults to `60`.

## How it works

The application will:

1. get your current IP using [the IP Geolocation API](https://ip-api.com/)
1. try to resolve the `FQDN` and compare the result with the current IP from step 1
1. if `FQDN` resolve has failed or IPs missmatch the `FQDN` record will be created/updated
