# Use case
Use AWS Route53 as a dynamic DNS service

## Usage
Pull docker image:
```bash
docker pull zmazay/dynamic-dns-route53
```

Or build it yourself:
```bash
docker build . -t dynamic-dns-route53
```

Create a configuration file similar to the following:
```yaml
route53_zones:
  - zone: example1.com
    hostnames:
    - example1.com
    - subdomain0.example1.com
    - subdomain1.example1.com
    - subdomain2.example1.com 
  - zone: example2.com
    hostnames:
    - example1.com
```

Map the directory containing the configuration file to the container with the following argument:
```bash
-v /host/directory:/container/directory
```

Set required variables:
```bash
export AWS_ACCESS_KEY_ID=******
export AWS_SECRET_ACCESS_KEY=************
export CONFIG_FILE=/path/to/your/config_file # should be exact path on the container
```

Run the DNS updater:
```bash
docker run -e "AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}" -e "AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}" -e "CONFIG_FILE=${CONFIG_FILE}" <IMAGE_ID>
```

## How it works
The container runs continuously and every 15 minutes executes certain ansible playbook for updating your Route53 zone(s).
