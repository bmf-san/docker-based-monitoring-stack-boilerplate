ARG FLUENTD_IMAGE_NAME=${FLUENTD_IMAGE_NAME}
ARG FLUENTD_IMAGE_TAG=${FLUENTD_IMAGE_TAG}
ARG PLATFORM=${PLATFORM}

FROM --platform=${PLATFORM} ${FLUENTD_IMAGE_NAME}:${FLUENTD_IMAGE_TAG}

USER root

RUN gem install fluent-plugin-elasticsearch \ 
	&& gem install fluent-plugin-rewrite-tag-filter \
	&& gem install fluent-plugin-mysqlslowquery

USER fluent
