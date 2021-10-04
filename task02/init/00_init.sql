create user sanya with encrypted password 'passwd';
grant all privileges on database sauna to sanya;

-- create the extension
CREATE SCHEMA faker;
CREATE EXTENSION faker SCHEMA faker CASCADE;
GRANT USAGE ON SCHEMA faker TO sanya;
