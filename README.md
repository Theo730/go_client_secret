# go_client_secret
Микросервис для формирования client_secret при подключении к ЕСИА через oAuth 2.0.
## GOST
В папку проекта необходимо добавить файлы сертификата и ключа(cert.pem & key.pem) соответственно. Можно использовать тестовые ключи, которые выдаются при оформлении доступа на технологический портал.
## REST
Запрос

  curl http://localhost:5959/api/v1/getClientSecret

Ответ

  {"code":0,"text":"","error":"","data":{"client_secret":"MIIJ8wYJKoZIhvcNAQcCoIIJ5DC....xUrNJ4GonIKPMB9eWdA","state":"65642552-2b34-e4ce-adab-6a94a91a92a9","data":"2021.12.28+11%3A07%3A32+%2B0300"}}

переменные client_secret, state, data  используются для получения токена на портале Единая система идентификации и аутентификации.

## Установка

  git clone https://github.com/Theo730/go_client_secret.git
  cd ./go_client_secret


###Установка и сборка в режиме службы

  go build

  ./go_client_secret -Cert ./cert.pem -Key ./key.pem

###Установка и сборка в режиме Docker

  docker build -t go_client_secret -f Dockerfile .
  
  docker run -p 5959:5959 go_client_secret
