run: build
	.\bin\pasteleria

make prod
	.\bin\pasteleria

build: 
	go build -o bin/pasteleria.exe cmd/api/main.go

migrate: 
	go build -o bin/migrate.exe cmd/migration/main.go
	.\bin\migrate

drop:
	go build -o bin/drop.exe cmd/drop/main.go
	.\bin\drop