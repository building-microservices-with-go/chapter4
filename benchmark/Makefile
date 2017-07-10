buildserver:
	CGO_ENABLED=0 GOOS=linux go build -o ./kittenserver ./main.go

runserver: buildserver
	docker-compose -p kittenserver -f ./docker-compose.yml up --build

