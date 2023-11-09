FROM golang:1.19 as build

# set working directory
WORKDIR /repo

# copy source code
COPY . .

# build binary
RUN go build -o /app/bin/program ./main.go

FROM alpine:latest

# set working directory
WORKDIR /app

# copy binary
COPY --from=build /repo/bin/program /app/program

# copy copy template & static file
COPY ./template /app/template
COPY ./static /app/static
COPY ./conf /app/conf

# expose port
EXPOSE 18080

# run binary
CMD ["/app/program", "-c", "/app/conf/app.ini"]


