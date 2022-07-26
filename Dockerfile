FROM ubuntu:20.04
RUN apt update && \
    apt install -y curl && \
    apt-get install -y wget &&\
    rm -rf /var/lib/apt/lists/*
# Work inside the /tmp directory
WORKDIR /tmp
RUN curl https://storage.googleapis.com/golang/go1.18.linux-amd64.tar.gz -o go.tar.gz && \
    tar -zxf go.tar.gz && \
    rm -rf go.tar.gz && \
    mv go /go
ENV GOPATH /go
ENV PATH $PATH:/go/bin:$GOPATH/bin
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz && \
tar -xzf migrate.linux-amd64.tar.gz && \
mkdir -p /bin && \ 
mv migrate.linux-amd64 $GOPATH/bin/migrate && \
rm migrate.linux-amd64.tar.gz
ENV PATH /bin:$PATH
# If you enable this, then gcc is needed to debug your app
ENV CGO_ENABLED 0

# TODO: Add other dependencies and stuff here