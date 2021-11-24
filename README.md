Как запустить:

Билдить коммандой docker-compose build

Потом docker-compose run -e storage=memory -p 8080:8080 url-shortener (Для режима с сохранением в память, для режима postgress: -e storage=postgres)

Примеры запросов через curl:

POST:

curl -d "url=http://go.dev/doc/tutorial/getting-started" http://localhost:8080 

Вернет сокращенный url и сохранит пару

GET:

curl http://localhost:8080/....... (Вместо точек сокращенный url полученый от запроса POST)

Вернет полный url который сохраняли
