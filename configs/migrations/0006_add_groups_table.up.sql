CREATE TABLE public.groups (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    user_id bigint NOT NULL,
    client_id bigint NOT NULL,
    tags text[] NOT NULL
);


CREATE SEQUENCE public.groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_user_client_key UNIQUE (user_id, client_id);


ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.groups
    ADD CONSTRAINT fk_groups_user FOREIGN KEY (user_id) REFERENCES public.users(id);


ALTER TABLE ONLY public.groups
    ADD CONSTRAINT fk_groups_client FOREIGN KEY (client_id) REFERENCES public.clients(id);


CREATE INDEX idx_groups_user_client ON public.groups USING btree (user_id, client_id);
