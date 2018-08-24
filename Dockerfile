# STEP 1 build executable binary

FROM golang:onbuild as builder
COPY . $GOPATH/src/github.com/WianVos/whisper/
WORKDIR $GOPATH/src/github.com/WianVos/whisper/

#get dependancies
#you can also use dep
RUN go get -d -v

#build the binary
RUN go build -o /go/bin/whisper

# STEP 2 build a small image

# start from scratch
FROM scratch

# Copy our static executable
COPY --from=builder /go/bin/whisper /go/bin/whisper
EXPOSE 8000
ENTRYPOINT ["/go/bin/whisper"]