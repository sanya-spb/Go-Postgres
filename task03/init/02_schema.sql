\c sauna
SET ROLE sanya;

/*
    Предложить, на каких полях можно создать индексы для ускорения запросов из п. 1. Создать требуемые индексы (не более трёх)
*/

CREATE OR REPLACE VIEW public.personal_info
AS SELECT p.id,
    p.fname,
    p.lname,
    p.phone,
    p.email,
    s.id AS service_id,
    s.service,
    ps.price
   FROM personal p
     LEFT JOIN personal_service ps ON ps.personal_id = p.id
     LEFT JOIN services s ON s.id = ps.service_id;

CREATE OR REPLACE VIEW public.saunas_info
AS SELECT s.id,
    s.name,
    s.address,
    s.phone,
    s.contact,
    s.email,
    ss.service_id,
    s2.service,
    ss.capacity,
    ss.price
   FROM saunas s
     RIGHT JOIN sauna_service ss ON ss.sauna_id = s.id
     LEFT JOIN services s2 ON s2.id = ss.service_id;

CREATE INDEX personal_service_idx ON public.personal_service USING btree (personal_id, service_id);
CREATE INDEX sauna_service_idx ON public.sauna_service USING btree (sauna_id, service_id);
