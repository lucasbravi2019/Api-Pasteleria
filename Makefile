run: build
	.\bin\pasteleria

build: 
	go build -o bin/pasteleria.exe cmd/api/main.go

migrate: 
	go build -o bin/migrate.exe cmd/migration/main.go
	.\bin\migrate

drop:
	go build -o bin/drop cmd/drop/main.go
	copy cmd\drop\tables_drop.xml bin\tables_drop.xml
	./bin/drop