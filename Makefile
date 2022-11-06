.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t asic777/bot:0.1 .

start-container:
	docker run --env-file .env -p 80:80 asic777/bot:0.1