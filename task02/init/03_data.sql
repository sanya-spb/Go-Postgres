\c sauna
SET ROLE sanya;

/*
	3. Подготовить набор данных для вашей базы
*/

-- init the extension
SELECT faker.faker('ru_RU');

-- справочник предоставляемых услуг
INSERT INTO services (service, price, descr)
SELECT
    faker.job(),
    round(random()::numeric * 800 + 1, 2),
    faker.text()
FROM generate_series(1,10)
ON CONFLICT ON constraint services_ukey DO NOTHING;

-- персонал
INSERT INTO personal (fname, lname, email, phone)
SELECT
  faker.first_name(),
  faker.last_name(),
  faker.free_email(),
  faker.phone_number()
FROM generate_series(1,50)
ON CONFLICT ON constraint personal_ukey DO NOTHING;

-- список объектов
INSERT INTO saunas ("name", address, phone, contact, email)
SELECT
    faker.large_company(),
    faker.address(),
    faker.phone_number(),
    faker.name(),
    faker.company_email()
FROM generate_series(1,100)
ON CONFLICT ON constraint saunas_ukey DO NOTHING;

-- специализация персонала (8-8)
INSERT INTO personal_service (personal_id, service_id, price)
SELECT
    (select id from personal where rnd=rnd order by random() limit 1),
    (select id from services where rnd=rnd order by random() limit 1),
    round(random()::numeric * 500 + 1, 2)
FROM generate_series(1,100) rnd;

-- возможные услуги на объекте
INSERT INTO sauna_service (sauna_id, service_id, capacity, price)
SELECT
    (select id from saunas where rnd=rnd order by random() limit 1),
    (select id from services where rnd=rnd order by random() limit 1),
    round(random() * 20) + 1,
    round(random()::numeric * 1000 + 1, 2)
FROM generate_series(1,100) rnd;

-- список заказов
INSERT INTO orders (date_start, peoples_count, duration_common, summary)
SELECT
    faker.date_time_between('+1d'::text, '+1y', null::text)::timestamp,
    1, --update later
    1, --update later
    1::numeric --update later
FROM generate_series(1,1000);

-- услуги по заказу
INSERT INTO order_services (orders_id, service_id, peoples, duration)
SELECT
    (select id from orders where rnd=rnd order by random() limit 1),
    (select id from services where rnd=rnd order by random() limit 1),
    round(random() * 10) + 1,
    round(random() * 8) + 1
FROM generate_series(1,2000) rnd;

-- корректировка данных в таблице "список заказов"
delete from orders where id in (
    select id from orders
    except
    select orders_id from order_services
);

update orders
set
    peoples_count=tt.peoples_count,
    duration_common=tt.duration_common,
    summary=tt.summary
from (
    select
        orders.id,
        sum(order_services.peoples) as peoples_count,
        sum(order_services.duration) as duration_common,
        sum(services.price*order_services.duration*order_services.peoples) as summary
    from orders
    right join order_services on order_services.orders_id = orders.id
    left join services on order_services.service_id = services.id 
    group by orders.id
)tt
where orders.id = tt.id;
