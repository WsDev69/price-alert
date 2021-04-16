# Price Alert

Includes 3 microservices:

- api-service
- notification-service
- quote-service

## Getting start

### Prepare environment

#### Run docker-compose

Includes:

- Running DB (postgres)

```shell
docker-compose -f ./infrastructure/docker-compose.yml up -d
```

- Running Redis
- Kafka Server
- SMTP Server

#### Make migrations for DB schema

```shell
cd api-service
```

If you don't have `golang-migrate` on your environment

##### MacOS

```shell
make migrate-macos-install
```

##### Linux

```shell
make migrate-linux-install
```

#### Run migration

```shell
make migrate-up
```

### Run applications

For each microservice you can use this commands to run application

- Copy config sample

```shell
cp config_sample.json config.json
```

- Edit `config.json` if need then

- Run one of the command to start application

```shell
make run
```

or:

```shell
make build
./bin/api
```

or:

```shell
go run src/main.go
```

### Examples

#### Create User
*You should create a user for yourself to use application*

```shell
wget --no-check-certificate --quiet \
  --method POST \
  --timeout=0 \
  --header 'Content-Type: application/json' \
  --body-data '{
    "email": "testSave@gmail.com"
}' \
   '127.0.0.1:8080/api/v1/user'
```

##### Request body
```json
{
    "email": "testSave@gmail.com"
}
```
#### Response 
```json
{
    "id": "1f8680c9-5678-43d6-9354-256ce391bd3d",
    "email": "testSave@gmail.com"
}
```

#### Create Price Alert

```shell
wget --no-check-certificate --quiet \
  --method POST \
  --timeout=0 \
  --header 'Content-Type: application/json' \
  --body-data '{
    "from_symbol":"BTC",
	"to_symbol":"USD",
	"price":63742.2

}' \
   '127.0.0.1:8080/api/v1/price/alert?user_id=1f8680c9-5678-43d6-9354-256ce391bd3d'
```



##### Request body
```json
{
    "from_symbol":"BTC",
    "to_symbol":"USD",
    "price":63742.2
}
```
##### Response
```json
{
    "id": "f4cc691d-184f-4d9d-b5ee-56650e92cc80",
    "from_symbol": "BTC",
    "to_symbol": "USD",
    "price": 63742.2,
    "user_id": "1f8680c9-5678-43d6-9354-256ce391bd3d"
}
```



 


