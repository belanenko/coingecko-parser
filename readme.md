# Парсер coingecko.com

Функционал апи:
1. Получить историю валюты за прошедший год 
POST `/getHistory` 
json body: {"name":"bitcoin"}

env:
1. BINDADDRES - Адрес на котором поднимается сервис `":8080"`
2. DATABASE_URL - Строка подключения к postgress `"postgres://postgres:password@localhost:5432/history?sslmode=disable"` 
3. WALLETS - список валют для сервиса `"tether,xrp,bitcoin,terra-luna,shiba-inu"`
4. POSTGRES_DB - имя бд
