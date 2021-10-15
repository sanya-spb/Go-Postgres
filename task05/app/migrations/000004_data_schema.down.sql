-- SET ROLE sanya;

-- справочник предоставляемых услуг
TRUNCATE TABLE services CASCADE;

-- персонал
TRUNCATE TABLE personal CASCADE;

-- список объектов
TRUNCATE TABLE saunas CASCADE;

-- специализация персонала (8-8)
TRUNCATE TABLE personal_service CASCADE;

-- возможные услуги на объекте
TRUNCATE TABLE sauna_service CASCADE;

-- список заказов
TRUNCATE TABLE orders CASCADE;

-- услуги по заказу
TRUNCATE TABLE order_services CASCADE;
