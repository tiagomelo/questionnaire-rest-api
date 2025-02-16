--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.2 (Debian 17.2-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: answer_recommendations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.answer_recommendations (
    answer_ulid text NOT NULL,
    recommendation_ulid text NOT NULL
);


ALTER TABLE public.answer_recommendations OWNER TO postgres;

--
-- Name: answers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.answers (
    id integer NOT NULL,
    ulid text NOT NULL,
    question_ulid text NOT NULL,
    text text NOT NULL,
    next_question_ulid text,
    previous_question_ulid text,
    CONSTRAINT answers_ulid_check CHECK ((ulid ~ '^[0-9A-HJKMNP-TV-Z]{26}$'::text))
);


ALTER TABLE public.answers OWNER TO postgres;

--
-- Name: answers_flow; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.answers_flow (
    id integer NOT NULL,
    answer_ulid text NOT NULL,
    previous_answer_ulid text,
    next_question_ulid text
);


ALTER TABLE public.answers_flow OWNER TO postgres;

--
-- Name: answers_flow_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.answers_flow_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.answers_flow_id_seq OWNER TO postgres;

--
-- Name: answers_flow_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.answers_flow_id_seq OWNED BY public.answers_flow.id;


--
-- Name: answers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.answers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.answers_id_seq OWNER TO postgres;

--
-- Name: answers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.answers_id_seq OWNED BY public.answers.id;


--
-- Name: exclusions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.exclusions (
    answer_ulid text NOT NULL,
    reason text NOT NULL
);


ALTER TABLE public.exclusions OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id integer NOT NULL,
    ulid text NOT NULL,
    name text NOT NULL,
    identifier text NOT NULL,
    CONSTRAINT products_ulid_check CHECK ((ulid ~ '^[0-9A-HJKMNP-TV-Z]{26}$'::text))
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: questions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.questions (
    id integer NOT NULL,
    ulid text NOT NULL,
    label character varying(10) NOT NULL,
    text text NOT NULL,
    CONSTRAINT questions_ulid_check CHECK ((ulid ~ '^[0-9A-HJKMNP-TV-Z]{26}$'::text))
);


ALTER TABLE public.questions OWNER TO postgres;

--
-- Name: questions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.questions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.questions_id_seq OWNER TO postgres;

--
-- Name: questions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.questions_id_seq OWNED BY public.questions.id;


--
-- Name: recommendations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.recommendations (
    id integer NOT NULL,
    ulid text NOT NULL,
    product_ulid text NOT NULL,
    CONSTRAINT recommendations_ulid_check CHECK ((ulid ~ '^[0-9A-HJKMNP-TV-Z]{26}$'::text))
);


ALTER TABLE public.recommendations OWNER TO postgres;

--
-- Name: recommendations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.recommendations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.recommendations_id_seq OWNER TO postgres;

--
-- Name: recommendations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.recommendations_id_seq OWNED BY public.recommendations.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: answers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers ALTER COLUMN id SET DEFAULT nextval('public.answers_id_seq'::regclass);


--
-- Name: answers_flow id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers_flow ALTER COLUMN id SET DEFAULT nextval('public.answers_flow_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: questions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions ALTER COLUMN id SET DEFAULT nextval('public.questions_id_seq'::regclass);


--
-- Name: recommendations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recommendations ALTER COLUMN id SET DEFAULT nextval('public.recommendations_id_seq'::regclass);


--
-- Name: answer_recommendations answer_recommendations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer_recommendations
    ADD CONSTRAINT answer_recommendations_pkey PRIMARY KEY (answer_ulid, recommendation_ulid);


--
-- Name: answers_flow answers_flow_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers_flow
    ADD CONSTRAINT answers_flow_pkey PRIMARY KEY (id);


--
-- Name: answers_flow answers_flow_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers_flow
    ADD CONSTRAINT answers_flow_unique UNIQUE (answer_ulid, previous_answer_ulid);


--
-- Name: answers answers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers
    ADD CONSTRAINT answers_pkey PRIMARY KEY (id);


--
-- Name: answers answers_ulid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers
    ADD CONSTRAINT answers_ulid_key UNIQUE (ulid);


--
-- Name: exclusions exclusions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exclusions
    ADD CONSTRAINT exclusions_pkey PRIMARY KEY (answer_ulid);


--
-- Name: products products_identifier_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_identifier_key UNIQUE (identifier);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: products products_ulid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_ulid_key UNIQUE (ulid);


--
-- Name: questions questions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_pkey PRIMARY KEY (id);


--
-- Name: questions questions_ulid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_ulid_key UNIQUE (ulid);


--
-- Name: recommendations recommendations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recommendations
    ADD CONSTRAINT recommendations_pkey PRIMARY KEY (id);


--
-- Name: recommendations recommendations_ulid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recommendations
    ADD CONSTRAINT recommendations_ulid_key UNIQUE (ulid);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: answer_recommendations answer_recommendations_answer_ulid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer_recommendations
    ADD CONSTRAINT answer_recommendations_answer_ulid_fkey FOREIGN KEY (answer_ulid) REFERENCES public.answers(ulid);


--
-- Name: answer_recommendations answer_recommendations_recommendation_ulid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer_recommendations
    ADD CONSTRAINT answer_recommendations_recommendation_ulid_fkey FOREIGN KEY (recommendation_ulid) REFERENCES public.recommendations(ulid);


--
-- Name: answers_flow answers_flow_answer_ulid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers_flow
    ADD CONSTRAINT answers_flow_answer_ulid_fkey FOREIGN KEY (answer_ulid) REFERENCES public.answers(ulid) ON DELETE CASCADE;


--
-- Name: answers_flow answers_flow_next_question_ulid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers_flow
    ADD CONSTRAINT answers_flow_next_question_ulid_fkey FOREIGN KEY (next_question_ulid) REFERENCES public.questions(ulid) ON DELETE CASCADE;


--
-- Name: answers_flow answers_flow_previous_answer_ulid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers_flow
    ADD CONSTRAINT answers_flow_previous_answer_ulid_fkey FOREIGN KEY (previous_answer_ulid) REFERENCES public.answers(ulid) ON DELETE CASCADE;


--
-- Name: exclusions exclusions_answer_ulid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exclusions
    ADD CONSTRAINT exclusions_answer_ulid_fkey FOREIGN KEY (answer_ulid) REFERENCES public.answers(ulid) ON DELETE CASCADE;


--
-- Name: answers fk_next_question; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers
    ADD CONSTRAINT fk_next_question FOREIGN KEY (next_question_ulid) REFERENCES public.questions(ulid);


--
-- Name: answers fk_previous_question; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers
    ADD CONSTRAINT fk_previous_question FOREIGN KEY (previous_question_ulid) REFERENCES public.questions(ulid);


--
-- Name: recommendations fk_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recommendations
    ADD CONSTRAINT fk_product FOREIGN KEY (product_ulid) REFERENCES public.products(ulid);


--
-- Name: answers fk_question; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers
    ADD CONSTRAINT fk_question FOREIGN KEY (question_ulid) REFERENCES public.questions(ulid);


--
-- PostgreSQL database dump complete
--

