-- 1. Развернуть сервер PostgreSQL в Docker
/*
docker run \
    --rm -it \
    -p 5432:5432 \
    --name postgres \
    -e POSTGRES_PASSWORD=passwd \
    -e PGDATA=/var/lib/postgresql/data \
    -v $(pwd)/postgres/data/:/var/lib/postgresql/data \
    -v $(pwd)/init:/docker-entrypoint-initdb.d \
    postgres:13.4
 */

-- 2. Создать пользователя и базу данных

create database gb;
create user sanya with encrypted password 'passwd';
grant all privileges on database gb to sanya;

\c gb
SET ROLE sanya;
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