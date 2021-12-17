# NOTE: see https://www.robustperception.io/environment-substitution-with-docker
ARG ALPINE_IMAGE_NAME=${ALPINE_IMAGE_NAME}
ARG ALPINE_IMAGE_TAG=${ALPINE_IMAGE_TAG}
ARG PROMETHEUS_IMAGE_NAME=${PROMETHEUS_IMAGE_NAME}
ARG PROMETHEUS_IMAGE_TAG=${PROMETHEUS_IMAGE_TAG}
ARG PLATFORM=${PLATFORM}

FROM --platform=${PLATFORM} ${PROMETHEUS_IMAGE_NAME}:${PROMETHEUS_IMAGE_TAG} as build-stage

FROM --platform=${PLATFORM} ${ALPINE_IMAGE_NAME}:${ALPINE_IMAGE_TAG}

RUN apk add gettext

COPY --from=build-stage /bin/prometheus /bin/prometheus

RUN mkdir -p /prometheus /etc/prometheus \
	&& chown -R nobody:nogroup etc/prometheus /prometheus

COPY ./template/prometheus.yml.template /template/prometheus.yml.template

USER nobody

VOLUME [ "/prometheus" ]

WORKDIR /prometheus