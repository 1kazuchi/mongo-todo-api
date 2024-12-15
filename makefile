include .env


up:
	@echo "Starting containers..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stoping containers..."
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/api/

start:
	MONGO_DB_USERNAME=${MONGO_DB_USERNAME} MONGO_DB_PASSWORD=${MONGO_DB_PASSWORD} ./${BINARY}
	
restart: build start

format_all_code:
	go fmt ./...

# connection string
# mongodb+srv://apisit936:<db_password>@cluster0.g1zqm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0


