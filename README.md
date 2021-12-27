# go_client_secret
Микросервис для формирования client_secret при подключении к ЕСИА через oAuth 2.0.
## GOST
В папку проекта необходимо добавить файлы сертификата и ключа(cert.pem & key.pem) соответственно. Или использовать тестовые ключи, которые выдаются при оформлении доступа на технологический портал.
## REST
Запрос
`` curl http://localhost:5959/api/v1/getClientSecret
Ответ
``{"code":0,"text":"","error":"","data":{"client_secret":"MIIJ8wYJKoZIhvcNAQcCoIIJ5DC....xUrNJ4GonIKPMBYz1I49eWdA","state":"65642552-2b34-e4ce-adab-6a94a91a92a9"}}
