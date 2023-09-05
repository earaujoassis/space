--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;


CREATE TABLE public.clients (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    uuid text NOT NULL,
    name text NOT NULL,
    description text,
    key text NOT NULL,
    secret text NOT NULL,
    scopes text NOT NULL,
    canonical_uri text[] NOT NULL,
    redirect_uri text[] NOT NULL,
    type text NOT NULL
);


CREATE SEQUENCE public.clients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.clients_id_seq OWNED BY public.clients.id;


CREATE TABLE public.languages (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    name text NOT NULL,
    iso_code text NOT NULL
);


CREATE SEQUENCE public.languages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.languages_id_seq OWNED BY public.languages.id;


CREATE TABLE public.sessions (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    uuid text NOT NULL,
    user_id bigint NOT NULL,
    client_id bigint NOT NULL,
    moment bigint NOT NULL,
    expires_in bigint DEFAULT 0 NOT NULL,
    ip text NOT NULL,
    user_agent text NOT NULL,
    invalidated boolean DEFAULT false NOT NULL,
    token text NOT NULL,
    token_type text NOT NULL,
    scopes text NOT NULL
);


CREATE SEQUENCE public.sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sessions_id_seq OWNED BY public.sessions.id;


CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    uuid text NOT NULL,
    public_id text NOT NULL,
    username text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    passphrase text NOT NULL,
    active boolean DEFAULT false NOT NULL,
    admin boolean DEFAULT false NOT NULL,
    client_id bigint NOT NULL,
    language_id bigint NOT NULL,
    timezone_identifier text DEFAULT 'GMT'::text NOT NULL,
    code_secret text NOT NULL,
    recover_secret text NOT NULL
);


CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


ALTER TABLE ONLY public.clients ALTER COLUMN id SET DEFAULT nextval('public.clients_id_seq'::regclass);


ALTER TABLE ONLY public.languages ALTER COLUMN id SET DEFAULT nextval('public.languages_id_seq'::regclass);


ALTER TABLE ONLY public.sessions ALTER COLUMN id SET DEFAULT nextval('public.sessions_id_seq'::regclass);


ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_key_key UNIQUE (key);


ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_name_key UNIQUE (name);


ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_uuid_key UNIQUE (uuid);


ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_iso_code_key UNIQUE (iso_code);


ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_name_key UNIQUE (name);


ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_token_key UNIQUE (token);


ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_uuid_key UNIQUE (uuid);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_public_id_key UNIQUE (public_id);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_uuid_key UNIQUE (uuid);


CREATE INDEX idx_clients_key ON public.clients USING btree (key);


CREATE INDEX idx_clients_name ON public.clients USING btree (name);


CREATE INDEX idx_clients_uuid ON public.clients USING btree (uuid);


CREATE INDEX idx_languages_name ON public.languages USING btree (name);


CREATE INDEX idx_sessions_ip ON public.sessions USING btree (ip);


CREATE INDEX idx_sessions_token ON public.sessions USING btree (token);


CREATE INDEX idx_sessions_token_type ON public.sessions USING btree (token_type);


CREATE INDEX idx_sessions_uuid ON public.sessions USING btree (uuid);


CREATE INDEX idx_users_email ON public.users USING btree (email);


CREATE INDEX idx_users_public_id ON public.users USING btree (public_id);


CREATE INDEX idx_users_username ON public.users USING btree (username);


CREATE INDEX idx_users_uuid ON public.users USING btree (uuid);


ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT fk_sessions_client FOREIGN KEY (client_id) REFERENCES public.clients(id);


ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT fk_sessions_user FOREIGN KEY (user_id) REFERENCES public.users(id);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_client FOREIGN KEY (client_id) REFERENCES public.clients(id);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_language FOREIGN KEY (language_id) REFERENCES public.languages(id);

--
-- PostgreSQL database dump complete
--

