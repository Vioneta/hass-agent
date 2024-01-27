# Copyright (c) 2023 Joshua Rich <joshua.rich@gmail.com>
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
FROM golang:1.21

WORKDIR /usr/src/go-hass-agent

# https://developer.fyne.io/started/#prerequisites
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get -y install gcc pkg-config libgl1-mesa-dev xorg-dev && rm -rf /var/lib/apt/lists/* /var/cache/apt/*

# copy the src to the workdir
ADD . .

# install build dependencies
RUN go install github.com/matryer/moq@latest && \
  go install golang.org/x/tools/cmd/stringer@latest && \
  go install golang.org/x/text/cmd/gotext@latest

# build the binary
RUN go generate ./... && \
  go build -v -o /go/bin/go-hass-agent && \
  go clean -cache -modcache && \
  rm -fr /usr/src/go-hass-agent

# remove fyne build dependencies
RUN apt-get -y remove gcc pkg-config libgl1-mesa-dev xorg-dev

# create a user to run the agent
RUN useradd -ms /bin/bash gouser
USER gouser
WORKDIR /home/gouser

ENTRYPOINT ["go-hass-agent"]
CMD ["--terminal"]
