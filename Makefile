BINARY=miniwallet

build:
	go build -o ${BINARY}

docker:
	docker build -t ${BINARY} .

run:
	docker run -it -d -p 80:80 ${BINARY}

mock:
	mockgen -source ./usecase/interfaces.go -destination ./usecase/mocks/interfaces_mock.go