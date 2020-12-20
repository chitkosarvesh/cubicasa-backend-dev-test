--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0 (Debian 13.0-1.pgdg100+1)
-- Dumped by pg_dump version 13.0 (Debian 13.0-1.pgdg100+1)

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

--
-- Name: hubs; Type: TABLE; Schema: public; Owner: cubicasa
--

CREATE TABLE public.hubs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    location text
);


ALTER TABLE public.hubs OWNER TO cubicasa;

--
-- Name: hubs_id_seq; Type: SEQUENCE; Schema: public; Owner: cubicasa
--

CREATE SEQUENCE public.hubs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.hubs_id_seq OWNER TO cubicasa;

--
-- Name: hubs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cubicasa
--

ALTER SEQUENCE public.hubs_id_seq OWNED BY public.hubs.id;


--
-- Name: teams; Type: TABLE; Schema: public; Owner: cubicasa
--

CREATE TABLE public.teams (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    hub_id bigint
);


ALTER TABLE public.teams OWNER TO cubicasa;

--
-- Name: teams_id_seq; Type: SEQUENCE; Schema: public; Owner: cubicasa
--

CREATE SEQUENCE public.teams_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.teams_id_seq OWNER TO cubicasa;

--
-- Name: teams_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cubicasa
--

ALTER SEQUENCE public.teams_id_seq OWNED BY public.teams.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: cubicasa
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text,
    last_name text,
    contact text,
    email text,
    username text,
    password text,
    team_id bigint
);


ALTER TABLE public.users OWNER TO cubicasa;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: cubicasa
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO cubicasa;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cubicasa
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: hubs id; Type: DEFAULT; Schema: public; Owner: cubicasa
--

ALTER TABLE ONLY public.hubs ALTER COLUMN id SET DEFAULT nextval('public.hubs_id_seq'::regclass);


--
-- Name: teams id; Type: DEFAULT; Schema: public; Owner: cubicasa
--

ALTER TABLE ONLY public.teams ALTER COLUMN id SET DEFAULT nextval('public.teams_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: cubicasa
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: hubs hubs_pkey; Type: CONSTRAINT; Schema: public; Owner: cubicasa
--

ALTER TABLE ONLY public.hubs
    ADD CONSTRAINT hubs_pkey PRIMARY KEY (id);


--
-- Name: teams teams_pkey; Type: CONSTRAINT; Schema: public; Owner: cubicasa
--

ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: cubicasa
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_hubs_deleted_at; Type: INDEX; Schema: public; Owner: cubicasa
--

CREATE INDEX idx_hubs_deleted_at ON public.hubs USING btree (deleted_at);


--
-- Name: idx_teams_deleted_at; Type: INDEX; Schema: public; Owner: cubicasa
--

CREATE INDEX idx_teams_deleted_at ON public.teams USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: cubicasa
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

