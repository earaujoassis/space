CREATE TABLE public.emails (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    uuid text NOT NULL,
    user_id bigint NOT NULL,
    address text NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT FALSE
);


CREATE SEQUENCE public.emails_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.emails_id_seq OWNED BY public.emails.id;


ALTER TABLE ONLY public.emails ALTER COLUMN id SET DEFAULT nextval('public.emails_id_seq'::regclass);


ALTER TABLE ONLY public.emails
    ADD CONSTRAINT emails_address_key UNIQUE (address);


ALTER TABLE ONLY public.emails
    ADD CONSTRAINT emails_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.emails
    ADD CONSTRAINT emails_uuid_key UNIQUE (uuid);


ALTER TABLE ONLY public.emails
    ADD CONSTRAINT fk_emails_user FOREIGN KEY (user_id) REFERENCES public.users(id);


CREATE INDEX idx_emails_address ON public.emails USING btree (address);


CREATE INDEX idx_emails_uuid ON public.emails USING btree (uuid);
