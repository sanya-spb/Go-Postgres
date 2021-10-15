# Практические задания
Код присылать как zip-архив директории проекта, в которой вы вели работу (т. е. иерархия директорий должна быть сохранена).
1. Реализовать приложение, которое реализует основной use-case вашей системы, т.е. поддерживает выполнение типовых запросов  (из файла queries.sql урока 3, достаточно покрыть один-два запроса). Необходимо реализовать только Storage Layer вашего приложения, т.е. только часть взаимодействия с базой данных.
1. Реализовать интеграционное тестирование функциональности по выборке данных из  базы.
1. Реализовать автоматизацию миграции структуры базы данных (файл schema.sql из предыдущих уроков). В файле README.md в корне проекта описать, как запускать миграцию структуры базы данных.

# Решение

1. запуск БД:
```
./docker_run_pg_05.sh
```
2. миграции:
```
$ migrate -database "postgresql://postgres:passwd@localhost:5432/sauna?sslmode=disable" -path ./app/migrations up                                                 1 ⨯
1/u init_schema (333.270606ms)
2/u skell_schema (1.335919474s)
3/u views_schema (1.53220948s)
4/u data_schema (2.124087911s)
```
3. интеграционные тесты:
```
$ go test ./... -tags=integration -v -count=1 
?       github.com/sanya-spb/Go-Postgres/api/handler    [no test files]
?       github.com/sanya-spb/Go-Postgres/api/router     [no test files]
?       github.com/sanya-spb/Go-Postgres/api/server     [no test files]
?       github.com/sanya-spb/Go-Postgres/app/repos/persons      [no test files]
?       github.com/sanya-spb/Go-Postgres/app/starter    [no test files]
?       github.com/sanya-spb/Go-Postgres/cmd/task05     [no test files]
?       github.com/sanya-spb/Go-Postgres/db/memory/persons/store        [no test files]
=== RUN   TestPersons_GetPerson
=== RUN   TestPersons_GetPerson/name
--- PASS: TestPersons_GetPerson (0.00s)
    --- PASS: TestPersons_GetPerson/name (0.00s)
PASS
ok      github.com/sanya-spb/Go-Postgres/db/postgres/persons/store      0.007s
```
4. запуск приложения:
```
$ go run cmd/task05/main.go
2021/10/15 03:44:16 Let's Go
2021/10/15 03:44:16 listen at: :8080
```
5. демонстрация (пароль: passwd):
```
$ curl http://localhost:8080/p/deep/fake                                                                          
{"status":"Not Found","error":"error when reading: read link error: failed to fetch the personal service: no rows in result set"}

$ psql -U sanya -d sauna -h localhost -c 'select * from personal order by random() limit 1;
Пароль пользователя sanya: 
 id | fname  |  lname   |       phone       |         email         
----+--------+----------+-------------------+-----------------------
 11 | Панфил | Матвеева | +7 (422) 353-5416 | valerija_72@gmail.com
(1 строка)

$ curl http://localhost:8080/p/Панфил/Матвеева  
{"id":11,"fname":"Панфил","lname":"Матвеева","phone":"+7 (422) 353-5416","email":"valerija_72@gmail.com"}
```
6. миграции down:
```
$ migrate -database "postgresql://postgres:passwd@localhost:5432/sauna?sslmode=disable" -path ./app/migrations down
Are you sure you want to apply all down migrations? [y/N]
y
Applying all down migrations
4/d data_schema (2.624246539s)
3/d views_schema (2.759522495s)
2/d skell_schema (2.920562378s)
1/d init_schema (3.072944953s)
```