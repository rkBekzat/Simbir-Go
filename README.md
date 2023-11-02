## Renting transport 

### Requirements

+ Go 
+ Docker
+ migrate - utility

### Used libs:

+  Gin
+ sqlx
+ viper
+ godotenv - to load .env file for access sensitive data 
+ pq - driver for db use with sqlx 
+ logrus - for logging the project and track wrongs

### Инструкция 

Скачиваем образ postgres с помощью этой команды: `docker pull postgres`

Дальше на основе этой образе создаем новый образ для базы данных: `make postgres`

Делаем миграцию, создать таблицы в базе данных: `make createdb`

Перед началом запуска импортируем в модуль нужный библиотеки: `make lib` 

Запуск приложений: `go run cmd/main.go`

Дальше делаем запрос на этот урл: `localhost:8000`

После запуска можно зайти и проверит эндпоинты: `http://localhost:8000/swagger/index.html`.

Также можно проверит в postman(работает). 