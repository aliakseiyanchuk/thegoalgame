MULTIPLATFORMS=linux/amd64,linux/arm64,linux/arm/v6,linux/386
BINARY_NAME=the_goal_game
DOCKER_IMAGE=lspwd2/${BINARY_NAME}
VERSION=0.1


compile_container_binaries:
	mkdir -p ./docker/dist/linux/amd64
	mkdir -p ./docker/dist/linux/arm64
	mkdir -p ./docker/dist/linux/arm/v6
	mkdir -p ./docker/dist/linux/386

	find ./docker/dist -name ${BINARY_NAME}* -exec /bin/rm {} \;

	GOOS=linux GOARCH=arm64 		go build -o ./docker/dist/linux/arm64/${BINARY_NAME} 	cmd/main.go
	GOOS=linux GOARCH=arm GOARM=6 	go build -o ./docker/dist/linux/arm/v6/${BINARY_NAME} 	cmd/main.go
	GOOS=linux GOARCH=amd64 		go build -o ./docker/dist/linux/amd64/${BINARY_NAME} 	cmd/main.go
	GOOS=linux GOARCH=386 			go build -o ./docker/dist/linux/386/${BINARY_NAME}		cmd/main.go


load_container: compile_container_binaries
	docker buildx build \
		--platform ${MULTIPLATFORMS}  \
		-t ${DOCKER_IMAGE}:${VERSION} -t ${DOCKER_IMAGE}:latest \
		--load \
		./docker/

push_container: compile_container_binaries
	docker buildx build \
		--platform ${MULTIPLATFORMS}  \
		-t ${DOCKER_IMAGE}:${VERSION} -t ${DOCKER_IMAGE}:latest \
		--push \
		./docker/
