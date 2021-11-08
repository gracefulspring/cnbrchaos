FROM cnbr/infra-alpine:v1.0.0

LABEL maintainer="hatech"

ENV OPERATOR=/usr/local/bin/chaos-operator
COPY ./chaos-operator ${OPERATOR}

ENTRYPOINT ["/usr/local/bin/chaos-operator"]
