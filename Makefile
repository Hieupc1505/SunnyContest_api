# name app
APP_NAME=server
DOCKER_DB_NAME=sunny
DB_NAME=SunnyContest
DB_URL=postgresql://root:secret@localhost:5432/SunnyContest?sslmode=disable

.PHONY: postgres createdb dropdb migrateup migratedown

run:
	go run ./cmd/${APP_NAME}/

#Định nghĩa target cho Postgress
postgres:
	docker run --name ${DOCKER_DB_NAME} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

# Định nghĩa target để tạo database
createdb:
	docker exec -it ${DOCKER_DB_NAME} createdb --username=root --owner=root ${DB_NAME}

# Định nghĩa target để xóa database
dropdb:
	docker exec -it ${DOCKER_DB_NAME} dropdb ${DB_NAME}

#migrate create -ext sql -dir db/migration -seq init_schema
#Định nghĩa target để thực hiện tạo table db/migration
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

# Định nghĩa target để thực hiện xóa db/migration
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

# Định nghĩa sqlc gen
sqlc:
	sqlc generate