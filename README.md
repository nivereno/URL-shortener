<h1>URL-shortener</h1>

<h2>Как запустить:</h2>

Билдить коммандой 
```
docker-compose build
```

Потом для режима сохранения в память
```
docker-compose run -e storage=memory -p 8080:8080 url-shortener 
```
А для режима сохранения в postgres датабазу
```
docker-compose run -e storage=postgres -p 8080:8080 url-shortener
```

<h2>Примеры запросов через curl:</h2>
Иногда curl в винде не работает правильно, если запрос делается из powershell иногда нужно сделать 
```
remove-item alisa:\curl
```

POST:
```
curl -d "url=http://go.dev/doc/tutorial/getting-started" http://localhost:8080
```

Вернет сокращенный url и сохранит пару

GET:
```
curl http://localhost:8080/....... (Вместо точек сокращенный url полученый от запроса POST)
```

Вернет полный url который сохраняли
