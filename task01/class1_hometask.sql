-- 1. Развернуть сервер PostgreSQL в Docker
/*
 * $ docker run --rm -it -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=password -e PGDATA=/var/lib/postgresql/data -v $(pwd)/postgres/data/:/var/lib/postgresql/data postgres:13.4
 */

-- 2. Создать пользователя и базу данных

CREATE ROLE sanya WITH 
	SUPERUSER
	CREATEDB
	CREATEROLE
	INHERIT
	LOGIN
	NOREPLICATION
	NOBYPASSRLS
	CONNECTION LIMIT UNLIMITED;

create database gb;

create user test with encrypted password 'test123';
grant all privileges on database gb to test;

CREATE SCHEMA task01;

-- 3. В базе из пункта 2 создать таблицу: не менее трёх столбцов различных типов

CREATE TABLE task01.tbl_3 (
	id serial not NULL,
	a text NOT NULL,
	n numeric NULL,
	date_at timestamp with time zone NOT NULL DEFAULT now(),
	CONSTRAINT tbl_3_pkey PRIMARY KEY (id),
	CONSTRAINT tbl_3_ukey UNIQUE (a)
);

-- Column comments

COMMENT ON COLUMN task01.tbl_3.id IS 'ID';
COMMENT ON COLUMN task01.tbl_3.a IS 'Param';
COMMENT ON COLUMN task01.tbl_3.n IS 'Num value';
COMMENT ON COLUMN task01.tbl_3.date_at IS 'Created at';

-- 4. В таблицу из пункта 3 вставить не менее трёх строк.

insert into task01.tbl_3 (a, n) values ('pi', 3.1415::numeric);
insert into task01.tbl_3 (a, n) values ('rub', 643::numeric);
insert into task01.tbl_3 (a, n) 
select tt.a, ascii(tt.a) 
from (
	select unnest(string_to_array('a,b,c', ',')) as a
)tt;

-- select * from task01.tbl_3

-- 5. Используя мета-команды psql, вывести список всех сущностей в базе данных из пункта 2
/*
gb=# \l
                                 Список баз данных
    Имя    | Владелец | Кодировка | LC_COLLATE |  LC_CTYPE  |     Права доступа     
-----------+----------+-----------+------------+------------+-----------------------
 gb        | postgres | UTF8      | en_US.utf8 | en_US.utf8 | =Tc/postgres         +
           |          |           |            |            | postgres=CTc/postgres+
           |          |           |            |            | test=CTc/postgres
 postgres  | postgres | UTF8      | en_US.utf8 | en_US.utf8 | 
 template0 | postgres | UTF8      | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |           |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8      | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |           |            |            | postgres=CTc/postgres
(4 строки)

gb=# \d
Отношения не найдены.
gb=# \dt
Отношения не найдены.
gb=# \dn
    Список схем
  Имя   | Владелец 
--------+----------
 public | postgres
 task01 | sanya
(2 строки)

gb=# set search_path to "task01";
SET
gb=# \dt
          Список отношений
 Схема  |  Имя  |   Тип   | Владелец 
--------+-------+---------+----------
 task01 | tbl_3 | таблица | sanya
(1 строка)

gb=# \dT
  Список типов данных
 Схема | Имя | Описание 
-------+-----+----------
(0 строк)

*/

-- 6. Придумать проект, над которым вы будете работать в последующих уроках.
/*
 * см. schema.sql
 */

