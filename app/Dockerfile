# Dockerfile extending the generic Go image with application files for a
# single application.
FROM gcr.io/google-appengine/golang

RUN apt-get update
RUN apt-get install -y apt-utils build-essential git wget curl libssl-dev libzmq3-dev

RUN go get -u -d github.com/GoogleCloudPlatform/golang-samples/appengine_flexible/datastore
RUN go get -u github.com/daludaluking/ons-sawtooth-sdk/ons_pb2
RUN go get -u github.com/daludaluking/ons-sawtooth-sdk/protobuf/batch_pb2
RUN go get -u github.com/daludaluking/ons-sawtooth-sdk/protobuf/transaction_pb2
RUN go get -u github.com/daludaluking/ons-sawtooth-sdk/signing
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/gorilla/handlers
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket

COPY . /go/src/app
RUN go-wrapper install -tags appenginevm
