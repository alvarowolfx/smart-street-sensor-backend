all: run

run-docker:
	docker stop smart-street-sensor-backend || :
	docker rm smart-street-sensor-backend || :
	docker run -d --name=smart-street-sensor-backend -p 8080:10000 -p 80:10001 gcr.io/smart-street-sensor/backend

run-grpc-test:
	go test ./server -run=grpc.* -bench=. -benchmem -benchtime=5s

run-http-test:
	go test ./server -run=http.* -bench=Http.* -benchmem -benchtime=5s

build:
	GOOS=linux go build -o app
	docker build -t gcr.io/smart-street-sensor/backend .
	rm -f app

push:
	gcloud docker -- push gcr.io/smart-street-sensor/backend

fmt: 
	gom fmt .