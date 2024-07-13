install:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql
	go get -u github.com/golang-migrate/migrate/v4
	go get github.com/dgrijalva/jwt-go
	go get github.com/go-redis/redis/v8
	go get golang.org/x/crypto/bcrypt

migrate-mysql:
	migrate -path ./migrations -database "mysql://root:root@tcp(localhost:13306)/go_base" up

migrate-postgres:
	migrate -path ./migrations -database "postgres://root:root@localhost:15432/go_base?sslmode=disable" up

run:
	go run main.go