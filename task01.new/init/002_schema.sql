-- 6. Придумать проект, над которым вы будете работать в последующих уроках.

/*
Система бронирования бань и саун.
*/

-- 7. Кратко (не более 10 предложений) описать суть проекта
/*
Суть проекта:
Планирование расписания групп обслуживающего персонала (массажисты, парильщики) под мероприятия (выезд к клиенту, либо работа в общественной бане/сауне)

Имеется следующие сущности:
	Персонал:
		- имя, фамилия
		- специализация (массажист, парильщик)
		- цена работы (руб/час)
	Объект:
		- название
		- адрес
		- доступные услуги
		- вместимость
		- цена аренды (руб/час)

Необходимо под определенный заказ сформировать персонал и подобрать соответствующий объект
	Заказ:
		- номер
		- дата-время начала
		- продолжительность
		- кол-во человек
		- услуги
		- стоимость услуг
		
(пока на этом остановим ТЗ..)
*/

-- 8. Разработать структуру базы данных, которая будет фундаментом для выбранного проекта

create database sauna;
grant all privileges on database sauna to sanya;

\c sauna
SET ROLE sanya;

-- справочник предоставляемых услуг
CREATE TABLE services (
	id serial not null,
	service text not null,
	descr text not null,
	CONSTRAINT services_pkey PRIMARY KEY (id),
	CONSTRAINT services_ukey UNIQUE (service)
);

insert into services (service, descr) values ('парилка', 'удары веником по спине');
insert into services (service, descr) values ('массаж', 'мыльный массаж');
--select * from services;

-- Персонал
CREATE TABLE personal (
	id serial not null,
	fname text not null,
	lname text not null,
	CONSTRAINT personal_pkey PRIMARY KEY (id)
);

-- Специализация персонала (8-8)
CREATE TABLE personal_service (
	id serial not null,
	personal_id bigint not null,
	service_id bigint not null,
	price numeric not null,
	CONSTRAINT personal_service_pkey PRIMARY KEY (id),
	CONSTRAINT personal_service_personal_fkey FOREIGN KEY (personal_id) REFERENCES public.personal(id),
	CONSTRAINT personal_service_service_fkey FOREIGN KEY (service_id) REFERENCES public.services(id)
);

-- Список объектов
CREATE TABLE saunas (
	id serial not null,
	"name" text not null,
	address text not null,
	CONSTRAINT saunas_pkey PRIMARY KEY (id)
);

-- Возможные услуги на объекте
CREATE TABLE sauna_service (
	id serial not null,
	sauna_id bigint not null,
	service_id bigint not null,
	capacity int not null,
	price numeric not null,
	CONSTRAINT sauna_service_pkey PRIMARY KEY (id),
	CONSTRAINT sauna_service_sauna_fkey FOREIGN KEY (sauna_id) REFERENCES public.saunas(id),
	CONSTRAINT sauna_service_service_fkey FOREIGN KEY (service_id) REFERENCES public.services(id)
);

-- Список заказов
CREATE TABLE orders (
	id serial not null,
	date_start timestamp not null,
	peoples_count int not null,
	duration_common int not null,
	summary numeric not null,
	CONSTRAINT orders_pkey PRIMARY KEY (id)
);

-- Услуги по заказу
CREATE TABLE order_services (
	id serial not null,
	orders_id bigint not null,
	service_id bigint not null,
	peoples int not null,
	duration int not null,
	CONSTRAINT order_services_pkey PRIMARY KEY (id),
	CONSTRAINT sauna_service_orders_fkey FOREIGN KEY (orders_id) REFERENCES public.orders(id),
	CONSTRAINT sauna_service_service_fkey FOREIGN KEY (service_id) REFERENCES public.services(id)
);
