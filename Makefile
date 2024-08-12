BINARY_NAME=filter_service

run:
	go run app/cmd/main.go

git:
	git add .
	git commit -a -m "$m"
	git push -u origin main

gen:
	protoc -I=app/proto -I=/home/ivan/go/src/github.com/marketconnect/filters/googleapis --go_out=app/gen/ --go-grpc_out=app/gen/ --grpc-gateway_out=app/gen/ --openapiv2_out=app/gen/docs app/proto/*.proto
	protoc --dart_out=grpc:../rewild/lib/pb/ -Iapp/proto -I=/home/ivan/go/src/github.com/marketconnect/filters/googleapis app/proto/*.proto
	protoc -I=app/proto --go_out=../api_bridge/app/gen/ app/proto/*.proto
	protoc --go-grpc_out=../api_bridge/app/gen/ app/proto/*.proto -I=app/proto
	protoc --dart_out=grpc:../rewild/lib/pb/ -Iapp/proto app/proto/*.proto

migrate_up:
	migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5432/mc_db?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5432/mc_db?sslmode=disable' down


build:
	rm -rf app/gen/*
	mkdir -p app/gen/docs
	$(MAKE) gen
	# go test -coverprofile=coverage.out ./...
	rm -f ${BINARY_NAME}
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINARY_NAME} app/cmd/main.go
	echo "Built ${BINARY_NAME}"

test:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# pg_dump -h localhost -U postgres -Fc mystats_db > mystats_db_23092023.sql
# pg_restore -h localhost -U postgres -d mystats_db mystats_db_23092023.sql
#
# pg_dump -h localhost -U postgres -t "public.api_stock" "mystats_db" | psql -U postgres -d "public.stock" "new_db"
