CREATE TABLE public.settings (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    user_id bigint NOT NULL,
    realm text NOT NULL,
    category text NOT NULL,
    property text NOT NULL,
    type text NOT NULL,
    value text NOT NULL
);


CREATE SEQUENCE public.settings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.settings_id_seq OWNED BY public.settings.id;


ALTER TABLE ONLY public.settings ALTER COLUMN id SET DEFAULT nextval('public.settings_id_seq'::regclass);


ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_user_realm_category_property_key UNIQUE (user_id, realm, category, property);


ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.settings
    ADD CONSTRAINT fk_settings_user FOREIGN KEY (user_id) REFERENCES public.users(id);


CREATE INDEX idx_settings_address ON public.settings USING btree (user_id, realm, category, property);
