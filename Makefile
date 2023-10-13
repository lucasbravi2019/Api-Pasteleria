run: build
	.\bin\pasteleria

build: 
	go build -o bin/pasteleria.exe cmd/api/main.go

migrate: 
	go build -o bin/migrate cmd/migration/main.go
	copy cmd\migration\tables.xml bin\tables.xml
	./bin/migrate

drop:
	go build -o bin/drop cmd/drop/main.go
	copy cmd\drop\tables_drop.xml bin\tables_drop.xml
	./bin/drop