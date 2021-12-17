ARG APP_IMAGE_NAME=${APP_IMAGE_NAME}
ARG APP_IMAGE_TAG=${APP_IMAGE_TAG}
ARG ALPINE_IMAGE_NAME=${ALPINE_IMAGE_NAME}
ARG PLATFORM=${PLATFORM}

FROM --platform=${PLATFORM} ${APP_IMAGE_NAME}:${APP_IMAGE_TAG} as builder

WORKDIR /go/app

COPY . .

RUN apk update && \
	apk upgrade && \
	apk add --no-cache libc-dev gcc git openssh openssl bash

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o app

FROM --platform=${PLATFORM} ${ALPINE_IMAGE_NAME}

COPY --from=builder /go/app ./

CMD ["./app"]