
.PHONY: build clean

build: ./build/tomato

./build/tomato: $(shell find . -name "*.go" -type f)
	go build -o ./build/tomato

rund: 
	go run main.go demon

clean:
	rm -f ./build/tomato

migrate: 
	@goose -dir db/migrations sqlite /home/aman/.local/share/tomato/tomato.db up 

migrate-down:
	@goose -dir db/migrations sqlite /home/aman/.local/share/tomato/tomato.db down 

migrate-status:
	@goose -dir db/migrations sqlite /home/aman/.local/share/tomato/tomato.db status 
 
seed: 
	@goose -dir db/seeds sqlite /home/aman/.local/share/tomato/tomato.db up  
 
sql: 
	sqlc generate

rebuild: clean build

