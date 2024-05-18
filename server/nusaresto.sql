--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Ubuntu 14.10-1.pgdg22.04+1)
-- Dumped by pg_dump version 14.10 (Ubuntu 14.10-1.pgdg22.04+1)

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
-- Name: users; Type: TABLE; Schema: public; Owner: nusaresto
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text,
    email text NOT NULL,
    password text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    role_id integer
);


ALTER TABLE public.users OWNER TO nichola;

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: nusaresto
--

COPY public.users (id, name, email, password, created_at, updated_at, role_id) FROM stdin;



--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: nusaresto
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);



<<<<<<< HEAD
--
-- Name: users; Type: TABLE; Schema: public; Owner: nusaresto
--

CREATE TABLE public.users (
=======

-- 
-- Name: menu; Type: TABLE; Schema: public; Owner: nusaresto
-- 
CREATE TABLE public.menu (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    update_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
)
>>>>>>> menu

--
-- PostgreSQL database dump complete
--



