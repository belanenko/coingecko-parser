# Парсер coingecko.com

Функционал апи:
1. Получить историю валюты за прошедший год `/getHistory?currency={currencyName}`

env:
1. BINDADDRES - Адрес на котором поднимается сервис `":8080"`
2. DATABASE_URL - Строка подключения к postgress `"postgres://postgres:password@localhost:5432/history?sslmode=disable"` 
3. WALLETS - список валют для сервиса `"tether,xrp,bitcoin,terra-luna,shiba-inu"`


TODO:
1. Провести БОЛЬШОЙ рефакторинг
2. Сделать endpoints для апи:
2.1 Получить цену за период `/getHistory?currency={currencyName}&from={startTime}&to={endTime}`
3. Сделать горячее добавление/изменение списка валют
