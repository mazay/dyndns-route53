#!/usr/bin/env sh

cd /opt/dyn_route53

while true; do
    ansible-playbook -i "localhost," dyn_route53.yaml -e "config_file=${CONFIG_FILE}"
    sleep 900
done
