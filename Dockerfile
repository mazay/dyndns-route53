FROM alpine

USER root

RUN apk -v --update add \
        python \
        py-pip \
        ansible \
        && \
    pip install --upgrade boto && \
    apk -v --purge del py-pip && \
    rm /var/cache/apk/*

RUN mkdir -p /opt/dyn_route53

COPY dyn_route53.yaml /opt/dyn_route53
COPY ansible.cfg /opt/dyn_route53

COPY update_route53.sh /opt/dyn_route53
RUN chmod +x /opt/dyn_route53/update_route53.sh

WORKDIR /opt/dyn_route53

CMD ["/opt/dyn_route53/update_route53.sh"]
