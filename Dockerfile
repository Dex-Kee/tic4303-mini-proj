FROM golang:1.19 as build

# set working directory
WORKDIR /app

# copy source code
COPY . .

# build binary
RUN GOOS=linux CGO_ENABLED=0 go build -o tic4303 ./main.go ./wire_gen.go

FROM alpine:latest

# create USER
ARG DOCKER_USER=app
RUN addgroup -S $DOCKER_USER && adduser -S $DOCKER_USER -G $DOCKER_USER
USER $DOCKER_USER

# set working directory
WORKDIR /app

# copy binary
COPY --from=build /app/tic4303 /app/tic4303

# copy copy template & static file
COPY ./templates /app/templates
COPY ./static /app/static
COPY ./conf /app/conf

# expose port
EXPOSE 8080

# run binary
ENTRYPOINT ["./tic4303", "-c", "/app/conf/app.ini"]