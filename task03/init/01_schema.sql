\c sauna
SET ROLE sanya;

/*
	1. Добавить ограничения foreign keys для всех имеющихся связей между таблицами в БД, созданной в первом занятии
*/


-- справочник предоставляемых услуг
CREATE TABLE services (
	id serial not null,
	service text not null,
	price numeric not null,
	descr text not null,
	CONSTRAINT services_pkey PRIMARY KEY (id),
	CONSTRAINT services_ukey UNIQUE (service),
	CONSTRAINT services_check CHECK (price>=0)
);
COMMENT ON TABLE public.services IS 'справочник предоставляемых услуг';
COMMENT ON COLUMN public.services.service IS 'услуга';
COMMENT ON COLUMN public.services.price IS 'цена (руб/час)';
COMMENT ON COLUMN public.services.descr IS 'описание услуги';


-- персонал
CREATE TABLE personal (
	id serial not null,
	fname text not null,
	lname text not null,
	phone text not null,
	email text not null,
	CONSTRAINT personal_pkey PRIMARY KEY (id),
	CONSTRAINT personal_ukey UNIQUE (fname,lname)
);
COMMENT ON TABLE public.personal IS 'персонал';
COMMENT ON COLUMN public.personal.fname IS 'имя';
COMMENT ON COLUMN public.personal.lname IS 'фамилия';
COMMENT ON COLUMN public.personal.phone IS 'телефон';
COMMENT ON COLUMN public.personal.email IS 'эл. почта';


-- специализация персонала (8-8)
CREATE TABLE personal_service (
	id serial not null,
	personal_id bigint not null,
	service_id bigint not null,
	price numeric not null,
	CONSTRAINT personal_service_pkey PRIMARY KEY (id),
	CONSTRAINT personal_service_personal_fkey FOREIGN KEY (personal_id) REFERENCES public.personal(id),
	CONSTRAINT personal_service_service_fkey FOREIGN KEY (service_id) REFERENCES public.services(id),
	CONSTRAINT personal_service_check CHECK (price>=0)
);
COMMENT ON TABLE public.personal_service IS 'специализация персонала (8-8)';
COMMENT ON COLUMN public.personal_service.price IS 'цена работы (руб/час)';

-- список объектов
CREATE TABLE saunas (
	id serial not null,
	"name" text not null,
	address text not null,
	phone text not null,
	contact text not null,
	email text not null,
	CONSTRAINT saunas_pkey PRIMARY KEY (id),
	CONSTRAINT saunas_ukey UNIQUE ("name")
);
COMMENT ON TABLE public.saunas IS 'список объектов';
COMMENT ON COLUMN public.saunas."name" IS 'название';
COMMENT ON COLUMN public.saunas.address IS 'адрес';
COMMENT ON COLUMN public.saunas.phone IS 'телефон';
COMMENT ON COLUMN public.saunas.contact IS 'конактное лицо/должность';
COMMENT ON COLUMN public.saunas.email IS 'эл. почта';


-- возможные услуги на объекте
CREATE TABLE sauna_service (
	id serial not null,
	sauna_id bigint not null,
	service_id bigint not null,
	capacity int not null,
	price numeric not null,
	CONSTRAINT sauna_service_pkey PRIMARY KEY (id),
	CONSTRAINT sauna_service_sauna_fkey FOREIGN KEY (sauna_id) REFERENCES public.saunas(id),
	CONSTRAINT sauna_service_service_fkey FOREIGN KEY (service_id) REFERENCES public.services(id),
	CONSTRAINT sauna_service_check CHECK (capacity>0 and price>=0)
);
COMMENT ON TABLE public.sauna_service IS 'возможные услуги на объекте';
COMMENT ON COLUMN public.sauna_service.capacity IS 'вместимость (кол-во клиентов)';
COMMENT ON COLUMN public.sauna_service.price IS 'цена аренды (руб/час)';

-- список заказов
CREATE TABLE orders (
	id serial not null,
	date_start timestamp not null,
	peoples_count int not null,
	duration_common int not null,
	summary numeric not null,
	CONSTRAINT orders_pkey PRIMARY KEY (id),
	CONSTRAINT orders_check CHECK (date_start>now() and peoples_count>0 and duration_common>0 and summary>=0)
);
COMMENT ON TABLE public.orders IS 'список заказов';
COMMENT ON COLUMN public.orders.date_start IS 'дата-время начала';
COMMENT ON COLUMN public.orders.peoples_count IS 'общее кол-во человек (клиентов)';
COMMENT ON COLUMN public.orders.duration_common IS 'общая продолжительность';
COMMENT ON COLUMN public.orders.summary IS 'общая стоимость услуг (итого)';

-- услуги по заказу
CREATE TABLE order_services (
	id serial not null,
	orders_id bigint not null,
	service_id bigint not null,
	peoples int not null,
	duration int not null,
	CONSTRAINT order_services_pkey PRIMARY KEY (id),
	CONSTRAINT sauna_service_orders_fkey FOREIGN KEY (orders_id) REFERENCES public.orders(id),
	CONSTRAINT sauna_service_service_fkey FOREIGN KEY (service_id) REFERENCES public.services(id),
	CONSTRAINT order_services_check CHECK (peoples>0 and duration>0)
);
COMMENT ON TABLE public.order_services IS 'услуги по заказу';
COMMENT ON COLUMN public.order_services.peoples IS 'кол-во человек (клиентов)';
COMMENT ON COLUMN public.order_services.duration IS 'продолжительность';