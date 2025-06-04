CREATE TABLE public.services (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    uuid text NOT NULL,
    name text NOT NULL,
    description text,
    canonical_uri text NOT NULL,
    logo_uri text,
    type text NOT NULL
);


CREATE SEQUENCE public.services_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.services_id_seq OWNED BY public.services.id;


ALTER TABLE ONLY public.services ALTER COLUMN id SET DEFAULT nextval('public.services_id_seq'::regclass);


ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_name_key UNIQUE (name);


ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_uuid_key UNIQUE (uuid);


CREATE INDEX idx_services_name ON public.services USING btree (name);


CREATE INDEX idx_services_uuid ON public.services USING btree (uuid);
