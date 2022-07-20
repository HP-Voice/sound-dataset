--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

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

--
-- Name: sound_dataset; Type: DATABASE; Schema: -; Owner: sound_dataset_u
--

CREATE DATABASE sound_dataset WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';


ALTER DATABASE sound_dataset OWNER TO sound_dataset_u;

\connect sound_dataset

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
-- Name: label; Type: TABLE; Schema: public; Owner: sound_dataset_u
--

CREATE TABLE public.label (
    id integer NOT NULL,
    name character varying(256) NOT NULL
);


ALTER TABLE public.label OWNER TO sound_dataset_u;

--
-- Name: label_id_seq; Type: SEQUENCE; Schema: public; Owner: sound_dataset_u
--

CREATE SEQUENCE public.label_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.label_id_seq OWNER TO sound_dataset_u;

--
-- Name: label_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sound_dataset_u
--

ALTER SEQUENCE public.label_id_seq OWNED BY public.label.id;


--
-- Name: label_state; Type: VIEW; Schema: public; Owner: sound_dataset_u
--

CREATE VIEW public.label_state AS
SELECT
    NULL::integer AS id,
    NULL::character varying(256) AS name,
    NULL::bigint AS amount,
    NULL::bigint AS amount_approved;


ALTER TABLE public.label_state OWNER TO sound_dataset_u;

--
-- Name: sample; Type: TABLE; Schema: public; Owner: sound_dataset_u
--

CREATE TABLE public.sample (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    label_id integer NOT NULL,
    verdict integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.sample OWNER TO sound_dataset_u;

--
-- Name: label id; Type: DEFAULT; Schema: public; Owner: sound_dataset_u
--

ALTER TABLE ONLY public.label ALTER COLUMN id SET DEFAULT nextval('public.label_id_seq'::regclass);


--
-- Name: label label_pk; Type: CONSTRAINT; Schema: public; Owner: sound_dataset_u
--

ALTER TABLE ONLY public.label
    ADD CONSTRAINT label_pk PRIMARY KEY (id);


--
-- Name: sample sample_pk; Type: CONSTRAINT; Schema: public; Owner: sound_dataset_u
--

ALTER TABLE ONLY public.sample
    ADD CONSTRAINT sample_pk PRIMARY KEY (id);


--
-- Name: sample_label_id_index; Type: INDEX; Schema: public; Owner: sound_dataset_u
--

CREATE INDEX sample_label_id_index ON public.sample USING btree (label_id);


--
-- Name: sample_verdict_index; Type: INDEX; Schema: public; Owner: sound_dataset_u
--

CREATE INDEX sample_verdict_index ON public.sample USING btree (verdict);


--
-- Name: label_state _RETURN; Type: RULE; Schema: public; Owner: sound_dataset_u
--

CREATE OR REPLACE VIEW public.label_state AS
 SELECT l.id,
    l.name,
    count(s.id) AS amount,
    sum(
        CASE
            WHEN (s.verdict = 1) THEN 1
            ELSE 0
        END) AS amount_approved
   FROM (public.label l
     LEFT JOIN public.sample s ON ((s.label_id = l.id)))
  GROUP BY l.id
  ORDER BY l.id;


--
-- Name: sample sample_label_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: sound_dataset_u
--

ALTER TABLE ONLY public.sample
    ADD CONSTRAINT sample_label_id_fk FOREIGN KEY (label_id) REFERENCES public.label(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

