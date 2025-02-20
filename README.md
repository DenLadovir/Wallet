The Wallet project is written in accordance with the terms of reference:

Good afternoon, dear applicant, this assignment is aimed at identifying your
real level in golang development, so treat it like working on
a project. Do it honestly and prove yourself to the maximum, good luck!

Write an application that accepts a REST request like
POST api/v1/wallet
{
valletId: UUID,
operationType: DEPOSIT or WITHDRAW,
amount: 1000
}
after performing the logic to change the account in the database
, it is also possible to get the wallet balance
GET api/v1/wallets/{WALLET_UUID}

stack:
Golang
Postgresql
Docker

Pay special attention to the problems when working in a competitive environment (1000 RPS
per wallet). No request should be processed (50X error)
the application should run in a docker container, the database too, the entire system
should be lifted using docker-compose
It is necessary to cover the application with tests

Upload the completed task to github and provide a link

The environment variables must be read from the config.env file.

Solve all the issues that arise on the assignment on your own, at your
discretion.

Additionally, a method for viewing the list of wallets (GetAllWalletsHandler) has been implemented

Commands for working via the console:
curl http://localhost:8080/api/v1/wallets - view the list of wallets in the database

curl -X POST http://localhost:8080/api/v1/wallets/323e4567-e89b-12d3-a456-426614174002/deposit -H "Content-Type: application/json" -d "{\"amount\": 100.00}" - performing a deposit operation

curl -X POST http://localhost:8080/api/v1/wallets/323e4567-e89b-12d3-a456-426614174002/withdraw -H "Content-Type: application/json" -d "{\"amount\": 10.00}" - execution of withdrawal operation

There were difficulties when trying to implement the tests, and this issue is in progress.

The project is deployed in a container on Docker.

--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Проект "Кошелёк" пишется в соответствии с техническим заданием:

Добрый день, уважаемый соискатель, данное задание нацелено на выявление вашего
реального уровня в разработке на golang, поэтому отнеситесь к нему, как к работе на
проекте. Выполняйте его честно и проявите себя по максимуму, удачи!

Напишите приложение, которое по REST принимает запрос вида
POST api/v1/wallet
{
valletId: UUID,
operationType: DEPOSIT or WITHDRAW,
amount: 1000
}
после выполнять логику по изменению счета в базе данных
также есть возможность получить баланс кошелька
GET api/v1/wallets/{WALLET_UUID}

стек:
Golang
Postgresql
Docker

Обратите особое внимание проблемам при работе в конкурентной среде (1000 RPS по
одному кошельку). Ни один запрос не должен быть не обработан (50Х error)
приложение должно запускаться в докер контейнере, база данных тоже, вся система
должна подниматься с помощью docker-compose
Необходимо покрыть приложение тестами

Решенное задание залить на гитхаб, предоставить ссылку

Переменные окружения должны считываться из файла config.env

Все возникающие вопросы по заданию решать самостоятельно, по своему
усмотрению.

Дополнительно был реализован метод для просмотра списка кошельков (GetAllWalletsHandler)

Команды для работы через консоль:
curl http://localhost:8080/api/v1/wallets - просмотр списка кошельков в базе данных

curl -X POST http://localhost:8080/api/v1/wallets/323e4567-e89b-12d3-a456-426614174002/deposit -H "Content-Type: application/json" -d "{\"amount\": 100.00}" - выполнение операции депозита

curl -X POST http://localhost:8080/api/v1/wallets/323e4567-e89b-12d3-a456-426614174002/withdraw -H "Content-Type: application/json" -d "{\"amount\": 10.00}" - выполнение операции снятия

При попытке реализовать тесты возникли трудности, данный вопрос в процессе.

Проект развёрнут в контейнере на Docker.
