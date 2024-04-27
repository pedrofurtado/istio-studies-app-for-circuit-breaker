FROM golang:1.22.2 as Builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o istio-studies-app-for-circuit-breaker main.go

FROM busybox:1.36.1
WORKDIR /app
COPY --from=Builder /app/istio-studies-app-for-circuit-breaker /app
CMD ["/app/istio-studies-app-for-circuit-breaker"]
EXPOSE 8080
