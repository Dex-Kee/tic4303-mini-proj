# Makefile

# Define variables
IMAGE_NAME = tic4303/app
TAG = latest
GO_BINARY_NAME = tic4303

# Define targets and their recipes

# Build the Go binary
build-binary:
	go build -o $(GO_BINARY_NAME) ./main.go ./wire_gen.go

# Build the Docker image
build-docker:
	docker build -t $(IMAGE_NAME):$(TAG) .

# Push the Docker image to a registry
docker-push:
	docker push $(IMAGE_NAME):$(TAG)

# Remove the local Docker image
docker-clean:
	docker rmi $(IMAGE_NAME):$(TAG)

# By default, build the Go binary and the Docker image
default: build-binary
