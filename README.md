# Cubicasa backend Dev Test

**Author: Sarvesh Chitko *(chitkosarvesh@gmail.com)***
Written in Golang using the Gin Framework, this application, when run will create 2 docker containers (web, db), load the initial database schema and start the web application over port 8080. To refer to the list of APIs please use the Postman Collection:
`Cubicasa Backend Dev Test.postman_collection.json`.
## Steps to run
- `chmod +x build.sh`
- `sh build.sh`
## Files
- ### database_schema.sql
    this contains the schema dump of the PostgreSQL database
- ### Dockerfile
    this has the Docker code to build the web application image
- ### docker-compose.yml
    this has the code to start the containers one after the other and work seamlessly as one environment

## Notable Libraries Used
- github.com/spf13/viper
- gorm.io/gorm
- github.com/gin-gionic/gin
