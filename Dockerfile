FROM golang
WORKDIR /
COPY svc.go svc.go
COPY openapi.json openapi.json
RUN go build svc.go
RUN chmod +x svc
ENTRYPOINT /svc
EXPOSE 10001
