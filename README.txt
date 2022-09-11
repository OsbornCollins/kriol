migrate create -seq -ext=.sql -dir=./migrations kriolmig

migrate -path ./migrations -database "postgresql://ocollins:Passw0rd@localhost:5432/kriol?sslmode=disable" -verbose down
migrate -path ./migrations -database "postgresql://ocollins:Passw0rd@localhost:5432/kriol?sslmode=disable" -verbose up

curl -i "http://localhost:4000/v1/healthcheck"
curl -X POST  "http://localhost:4000/v1/entries"
curl -i "http://localhost:4000/v1/entries/1"

curl -i "https://api.openweathermap.org/data/2.5/weather?lat=17.18&lon=88.49&appid=b1f25fb3764ecd95fda4c047436f18f8"
