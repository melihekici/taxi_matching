CREATE TABLE IF NOT EXISTS public.usertable
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    email text COLLATE pg_catalog."default" NOT NULL,
    username text COLLATE pg_catalog."default" NOT NULL,
    pass text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT usertable_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.usertable
    OWNER to postgres;
