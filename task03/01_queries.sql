/*
    Составить 3–5 типовых запросов к данным в созданном проекте БД
*/

-- для начала упростим себе жизнь
create or replace
view public.personal_info as
select
    p.id
    , p.fname
    , p.lname
    , p.phone
    , p.email
    , s.id as service_id
    , s.service
    , ps.price
from
    personal p
left join personal_service ps on ps.personal_id = p.id
left join services s on s.id = ps.service_id;

create or replace
view saunas_info as
select
    s.id
    , s."name"
    , s.address
    , s.phone
    , s.contact
    , s.email
    , ss.service_id
    , s2.service
    , ss.capacity
    , ss.price
from
    saunas s
right join sauna_service ss on
    ss.sauna_id = s.id
left join services s2 on
    s2.id = ss.service_id;

-- карточка сотрудника
select * from personal
where fname='Руслан' and lname='Аксенов';
  
-- поиск объекта
  select *
from
    saunas s
where
    address ~* 'бологое';
select *
from
    saunas s
where
    contact ~* 'агафонов';
select *
from
    saunas s
where
    name ~* 'тк';

-- подбор объекта по параметрам
select *
from
    saunas_info si
where
    service = 'Таксист'
    and capacity >5
order by
    "name" ;

-- информация о навыках конкретного сотрудника
select
    service_id
    , service
    , price
from
    personal_info
where
    fname = 'Руслан'
    and lname = 'Аксенов'
order by
    service;

-- поиск сотрудников с определенными навыками
select
    *
from
    personal_info
where
    service = 'Таксист'
order by
    lname
    , fname;

-- вариант с предварительным ознакомлением со списком навыков
select *
from
    personal_info
where
    service_id in (
        select
            id
        from
            services s
        where
            s.service in (
                'Таксист',
                'Телеграфист'
            )
    )
order by
    lname
    , fname;
