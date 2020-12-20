GOPATH=$(pwd) GOOS=linux GOARCH=amd64 go build -o bin/cubicasa src/cubicasa/main.go
docker build -t cubicasa_web .
docker-compose -f docker-compose.yml down
docker-compose -f docker-compose.yml up -d
sleep 5
docker-compose -f docker-compose.yml up -d
docker ps -a