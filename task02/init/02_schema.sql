\c sauna
SET ROLE sanya;

/*
	2. Выявить необходимые ограничения (constraints) и добавить их в структуру базы данных
*/

-- примитивные проверки некорректных данных (потаблично)
ALTER TABLE public.order_services ADD CONSTRAINT order_services_check CHECK (peoples>0 and duration>0);
ALTER TABLE public.orders ADD CONSTRAINT orders_check CHECK (date_start>now() and peoples_count>0 and duration_common>0 and summary>=0);
ALTER TABLE public.personal_service ADD CONSTRAINT personal_service_check CHECK (price>=0);
ALTER TABLE public.sauna_service ADD CONSTRAINT sauna_service_check CHECK (capacity>0 and price>=0);

-- защита от дублирования данных
ALTER TABLE public.personal ADD CONSTRAINT personal_ukey UNIQUE (fname,lname);
ALTER TABLE public.saunas ADD CONSTRAINT saunas_ukey UNIQUE ("name");
ALTER TABLE public.services ADD CONSTRAINT services_un UNIQUE (service);
