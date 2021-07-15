FROM golang
WORKDIR /
COPY svc.go svc.go
RUN go build svc.go
RUN chmod +x svc
ENTRYPOINT /svc
EXPOSE 10001
