CREATE INDEX personal_service_idx ON public.personal_service USING btree (personal_id, service_id);
CREATE INDEX sauna_service_idx ON public.sauna_service USING btree (sauna_id, service_id);
