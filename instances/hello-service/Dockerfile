FROM centos:7

RUN yum update -y

RUN mkdir -p /app/hello-service
COPY main /app/hello-service/main
COPY main.go /app/hello-service/main.go
COPY Dockerfile /app/hello-service/Dockerfile

EXPOSE 8080
CMD ["/app/hello-service/main"]
