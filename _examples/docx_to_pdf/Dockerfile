FROM golang:1.12.5-alpine3.9

ENV JAVA_HOME=/usr/lib/jvm/java-1.8-openjdk

# Input file
ENV INPUT_FILE_NAME mydoc.docx

# env build packages
ARG BUILD_PACKAGES="git wget make"

# Set the Current Working Directory inside the container
WORKDIR /usr/app

# Add the source code
ENV SRC_DIR=/usr/app/

# create volume data
RUN mkdir data

COPY . $SRC_DIR

RUN apk update && apk add --no-cache --update $BUILD_PACKAGES \
    dumb-init \
    openjdk8 \
    libreoffice libreoffice-writer libreoffice-base libreoffice-common\
    libreoffice-lang-de libreoffice-lang-uk \
    mesa-gl \
    ttf-freefont ttf-opensans ttf-ubuntu-font-family ttf-inconsolata ttf-liberation ttf-dejavu \
    && go get github.com/telkomdev/go-libreoffice \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o app \
    && apk del $BUILD_PACKAGES

VOLUME [ "/data" ]

ENTRYPOINT ["sh", "-c", "./app /usr/app/data/${INPUT_FILE_NAME}"]