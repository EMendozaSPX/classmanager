<<<<<<< HEAD
CREATE SCHEMA classmanager;

CREATE TABLE classmanager.class_student (
    id integer NOT NULL,
    class_id integer NOT NULL,
    student_id integer NOT NULL
);

CREATE TABLE classmanager.classes (
    id integer NOT NULL,
    class_id character varying(20) NOT NULL,
    teacher_id integer NOT NULL
);

CREATE TABLE classmanager.events (
    id integer NOT NULL,
    year_id integer NOT NULL,
    class_id integer NOT NULL,
    name character varying(100) NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL
);


ALTER TABLE classmanager.events OWNER TO emendoza;

--
-- TOC entry 207 (class 1259 OID 41148)
-- Name: events_id_seq; Type: SEQUENCE; Schema: classmanager; Owner: emendoza
--

CREATE SEQUENCE classmanager.events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE classmanager.events_id_seq OWNER TO emendoza;

--
-- TOC entry 2914 (class 0 OID 0)
-- Dependencies: 207
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: classmanager; Owner: emendoza
--

ALTER SEQUENCE classmanager.events_id_seq OWNED BY classmanager.events.id;


--
-- TOC entry 210 (class 1259 OID 41168)
-- Name: periods; Type: TABLE; Schema: classmanager; Owner: emendoza
--

CREATE TABLE classmanager.periods (
    id integer NOT NULL,
    year_id integer NOT NULL,
    name character varying(100) NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL
);


ALTER TABLE classmanager.periods OWNER TO emendoza;

--
-- TOC entry 209 (class 1259 OID 41166)
-- Name: periods_id_seq; Type: SEQUENCE; Schema: classmanager; Owner: emendoza
--

CREATE SEQUENCE classmanager.periods_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE classmanager.periods_id_seq OWNER TO emendoza;

--
-- TOC entry 2915 (class 0 OID 0)
-- Dependencies: 209
-- Name: periods_id_seq; Type: SEQUENCE OWNED BY; Schema: classmanager; Owner: emendoza
--

ALTER SEQUENCE classmanager.periods_id_seq OWNED BY classmanager.periods.id;


--
-- TOC entry 206 (class 1259 OID 41137)
-- Name: public_holidays; Type: TABLE; Schema: classmanager; Owner: emendoza
--

CREATE TABLE classmanager.public_holidays (
    id integer NOT NULL,
    year_id integer NOT NULL,
    name character varying(100) NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL
);


ALTER TABLE classmanager.public_holidays OWNER TO emendoza;

--
-- TOC entry 205 (class 1259 OID 41135)
-- Name: public_holidays_id_seq; Type: SEQUENCE; Schema: classmanager; Owner: emendoza
--

CREATE SEQUENCE classmanager.public_holidays_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE classmanager.public_holidays_id_seq OWNER TO emendoza;

--
-- TOC entry 2916 (class 0 OID 0)
-- Dependencies: 205
-- Name: public_holidays_id_seq; Type: SEQUENCE OWNED BY; Schema: classmanager; Owner: emendoza
--

ALTER SEQUENCE classmanager.public_holidays_id_seq OWNED BY classmanager.public_holidays.id;


--
-- TOC entry 212 (class 1259 OID 41182)
-- Name: terms; Type: TABLE; Schema: classmanager; Owner: emendoza
--

CREATE TABLE classmanager.terms (
    id integer NOT NULL,
    year_id integer NOT NULL,
    name character varying(100) NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL
);


ALTER TABLE classmanager.terms OWNER TO emendoza;

--
-- TOC entry 211 (class 1259 OID 41180)
-- Name: terms_id_seq; Type: SEQUENCE; Schema: classmanager; Owner: emendoza
--

CREATE SEQUENCE classmanager.terms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE classmanager.terms_id_seq OWNER TO emendoza;

--
-- TOC entry 2917 (class 0 OID 0)
-- Dependencies: 211
-- Name: terms_id_seq; Type: SEQUENCE OWNED BY; Schema: classmanager; Owner: emendoza
--

ALTER SEQUENCE classmanager.terms_id_seq OWNED BY classmanager.terms.id;


--
-- TOC entry 198 (class 1259 OID 41003)
-- Name: users; Type: TABLE; Schema: classmanager; Owner: emendoza
--

CREATE TABLE classmanager.users (
    id integer NOT NULL,
    role character varying(100) NOT NULL,
    username character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL
);


ALTER TABLE classmanager.users OWNER TO emendoza;

--
-- TOC entry 197 (class 1259 OID 41001)
-- Name: users_id_seq; Type: SEQUENCE; Schema: classmanager; Owner: emendoza
--

CREATE SEQUENCE classmanager.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE classmanager.users_id_seq OWNER TO emendoza;

--
-- TOC entry 2918 (class 0 OID 0)
-- Dependencies: 197
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: classmanager; Owner: emendoza
--

ALTER SEQUENCE classmanager.users_id_seq OWNED BY classmanager.users.id;


--
-- TOC entry 200 (class 1259 OID 41037)
-- Name: year; Type: TABLE; Schema: classmanager; Owner: emendoza
--

CREATE TABLE classmanager.year (
    id integer NOT NULL,
    year integer NOT NULL,
    year_group integer NOT NULL
=======
CREATE TABLE users (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    role VARCHAR(100) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE classes (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id VARCHAR(20) UNIQUE NOT NULL,
    teacher_id INTEGER NOT NULL REFERENCES users(id)
);

CREATE TABLE class_student (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES classes(id),
    student_id INTEGER NOT NULL REFERENCES users(id)
);

CREATE TABLE public_holidays (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL REFERENCES year(id),
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    year_id INTEGER NOT NULL REFERENCES year(id),
    class_id INTEGER NOT NULL REFERENCES classes(id),
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE TABLE timetable (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES classes(id),
    period_name VARCHAR(50) NOT NULL,
    week_day INTEGER NOT NULL
>>>>>>> develop
);


ALTER TABLE classmanager.year OWNER TO emendoza;

--
-- TOC entry 199 (class 1259 OID 41035)
-- Name: year_id_seq; Type: SEQUENCE; Schema: classmanager; Owner: emendoza
--

CREATE SEQUENCE classmanager.year_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE classmanager.year_id_seq OWNER TO emendoza;

--
-- TOC entry 2919 (class 0 OID 0)
-- Dependencies: 199
-- Name: year_id_seq; Type: SEQUENCE OWNED BY; Schema: classmanager; Owner: emendoza
--

ALTER SEQUENCE classmanager.year_id_seq OWNED BY classmanager.year.id;


--
-- TOC entry 2731 (class 2604 OID 41109)
-- Name: class_student id; Type: DEFAULT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.class_student ALTER COLUMN id SET DEFAULT nextval('classmanager.class_student_id_seq'::regclass);


--
-- TOC entry 2730 (class 2604 OID 41094)
-- Name: classes id; Type: DEFAULT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.classes ALTER COLUMN id SET DEFAULT nextval('classmanager.classes_id_seq'::regclass);


--
-- TOC entry 2733 (class 2604 OID 41153)
-- Name: events id; Type: DEFAULT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.events ALTER COLUMN id SET DEFAULT nextval('classmanager.events_id_seq'::regclass);


--
-- TOC entry 2734 (class 2604 OID 41171)
-- Name: periods id; Type: DEFAULT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.periods ALTER COLUMN id SET DEFAULT nextval('classmanager.periods_id_seq'::regclass);


--
-- TOC entry 2732 (class 2604 OID 41140)
-- Name: public_holidays id; Type: DEFAULT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.public_holidays ALTER COLUMN id SET DEFAULT nextval('classmanager.public_holidays_id_seq'::regclass);


--
-- TOC entry 2735 (class 2604 OID 41185)
-- Name: terms id; Type: DEFAULT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.terms ALTER COLUMN id SET DEFAULT nextval('classmanager.terms_id_seq'::regclass);


--
-- TOC entry 2728 (class 2604 OID 41006)
-- Name: users id; Type: DEFAULT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.users ALTER COLUMN id SET DEFAULT nextval('classmanager.users_id_seq'::regclass);


--
-- TOC entry 2729 (class 2604 OID 41040)
-- Name: year id; Type: DEFAULT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.year ALTER COLUMN id SET DEFAULT nextval('classmanager.year_id_seq'::regclass);


--
-- TOC entry 2898 (class 0 OID 41106)
-- Dependencies: 204
-- Data for Name: class_student; Type: TABLE DATA; Schema: classmanager; Owner: emendoza
--

COPY classmanager.class_student (id, class_id, student_id) FROM stdin;
\.


--
-- TOC entry 2896 (class 0 OID 41091)
-- Dependencies: 202
-- Data for Name: classes; Type: TABLE DATA; Schema: classmanager; Owner: emendoza
--

COPY classmanager.classes (id, class_id, teacher_id) FROM stdin;
1	12SDD	3
4	12CHEM2	3
5	12IPT	3
7	11SDD	3
\.


--
-- TOC entry 2902 (class 0 OID 41150)
-- Dependencies: 208
-- Data for Name: events; Type: TABLE DATA; Schema: classmanager; Owner: emendoza
--

COPY classmanager.events (id, year_id, class_id, name, start_time, end_time) FROM stdin;
\.


--
-- TOC entry 2904 (class 0 OID 41168)
-- Dependencies: 210
-- Data for Name: periods; Type: TABLE DATA; Schema: classmanager; Owner: emendoza
--

COPY classmanager.periods (id, year_id, name, start_time, end_time) FROM stdin;
\.


--
-- TOC entry 2900 (class 0 OID 41137)
-- Dependencies: 206
-- Data for Name: public_holidays; Type: TABLE DATA; Schema: classmanager; Owner: emendoza
--

COPY classmanager.public_holidays (id, year_id, name, start_time, end_time) FROM stdin;
\.


--
-- TOC entry 2906 (class 0 OID 41182)
-- Dependencies: 212
-- Data for Name: terms; Type: TABLE DATA; Schema: classmanager; Owner: emendoza
--

COPY classmanager.terms (id, year_id, name, start_time, end_time) FROM stdin;
\.


--
-- TOC entry 2892 (class 0 OID 41003)
-- Dependencies: 198
-- Data for Name: users; Type: TABLE DATA; Schema: classmanager; Owner: emendoza
--

COPY classmanager.users (id, role, username, email, password) FROM stdin;
1	admin	admin	admin@stpiusx.nsw.edu.au	$2a$10$KJ70NDRaXiA/ggoN98WP/eaAMeSCoDuStaZavYLurdXydbi6eqEBq
2	student	mend9	mend9@stpiusx.nsw.edu.au	$2a$10$4SCgAd5qovsHnbJbD6Tfa.arzwXm3n6QODzrWmfIFYALkPCYnYXka
3	teacher	jlai	jlai@stpiusx.nsw.edu.au	$2a$10$7/0nHSi6QPlRzYWLwVOE5.gdHLNuBlSJqmqOMbaoGb8xW6TXGpT/a
4	teacher	cab	jlai@stpiusx.sw.edu.au	$2a$10$W3Bir5rcwkPBAfgmqljQ7.GWJoVKviVC8uKEngoRC/058BoefhzHS
\.


--
-- TOC entry 2894 (class 0 OID 41037)
-- Dependencies: 200
-- Data for Name: year; Type: TABLE DATA; Schema: classmanager; Owner: emendoza
--

COPY classmanager.year (id, year, year_group) FROM stdin;
\.


--
-- TOC entry 2920 (class 0 OID 0)
-- Dependencies: 203
-- Name: class_student_id_seq; Type: SEQUENCE SET; Schema: classmanager; Owner: emendoza
--

SELECT pg_catalog.setval('classmanager.class_student_id_seq', 1, false);


--
-- TOC entry 2921 (class 0 OID 0)
-- Dependencies: 201
-- Name: classes_id_seq; Type: SEQUENCE SET; Schema: classmanager; Owner: emendoza
--

SELECT pg_catalog.setval('classmanager.classes_id_seq', 7, true);


--
-- TOC entry 2922 (class 0 OID 0)
-- Dependencies: 207
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: classmanager; Owner: emendoza
--

SELECT pg_catalog.setval('classmanager.events_id_seq', 1, false);


--
-- TOC entry 2923 (class 0 OID 0)
-- Dependencies: 209
-- Name: periods_id_seq; Type: SEQUENCE SET; Schema: classmanager; Owner: emendoza
--

SELECT pg_catalog.setval('classmanager.periods_id_seq', 1, false);


--
-- TOC entry 2924 (class 0 OID 0)
-- Dependencies: 205
-- Name: public_holidays_id_seq; Type: SEQUENCE SET; Schema: classmanager; Owner: emendoza
--

SELECT pg_catalog.setval('classmanager.public_holidays_id_seq', 1, false);


--
-- TOC entry 2925 (class 0 OID 0)
-- Dependencies: 211
-- Name: terms_id_seq; Type: SEQUENCE SET; Schema: classmanager; Owner: emendoza
--

SELECT pg_catalog.setval('classmanager.terms_id_seq', 1, false);


--
-- TOC entry 2926 (class 0 OID 0)
-- Dependencies: 197
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: classmanager; Owner: emendoza
--

SELECT pg_catalog.setval('classmanager.users_id_seq', 4, true);


--
-- TOC entry 2927 (class 0 OID 0)
-- Dependencies: 199
-- Name: year_id_seq; Type: SEQUENCE SET; Schema: classmanager; Owner: emendoza
--

SELECT pg_catalog.setval('classmanager.year_id_seq', 1, false);


--
-- TOC entry 2753 (class 2606 OID 41111)
-- Name: class_student class_student_pkey; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.class_student
    ADD CONSTRAINT class_student_pkey PRIMARY KEY (id);


--
-- TOC entry 2749 (class 2606 OID 41098)
-- Name: classes classes_class_id_key; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.classes
    ADD CONSTRAINT classes_class_id_key UNIQUE (class_id);


--
-- TOC entry 2751 (class 2606 OID 41096)
-- Name: classes classes_pkey; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.classes
    ADD CONSTRAINT classes_pkey PRIMARY KEY (id);


--
-- TOC entry 2757 (class 2606 OID 41155)
-- Name: events events_pkey; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- TOC entry 2759 (class 2606 OID 41173)
-- Name: periods periods_pkey; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.periods
    ADD CONSTRAINT periods_pkey PRIMARY KEY (id);


--
-- TOC entry 2755 (class 2606 OID 41142)
-- Name: public_holidays public_holidays_pkey; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.public_holidays
    ADD CONSTRAINT public_holidays_pkey PRIMARY KEY (id);


--
-- TOC entry 2761 (class 2606 OID 41187)
-- Name: terms terms_pkey; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.terms
    ADD CONSTRAINT terms_pkey PRIMARY KEY (id);


--
-- TOC entry 2737 (class 2606 OID 41012)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 2739 (class 2606 OID 41014)
-- Name: users users_password_key; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.users
    ADD CONSTRAINT users_password_key UNIQUE (password);


--
-- TOC entry 2741 (class 2606 OID 41008)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2743 (class 2606 OID 41010)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 2745 (class 2606 OID 41042)
-- Name: year year_pkey; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.year
    ADD CONSTRAINT year_pkey PRIMARY KEY (id);


--
-- TOC entry 2747 (class 2606 OID 41044)
-- Name: year year_year_key; Type: CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.year
    ADD CONSTRAINT year_year_key UNIQUE (year);


--
-- TOC entry 2763 (class 2606 OID 41112)
-- Name: class_student class_student_class_id_fkey; Type: FK CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.class_student
    ADD CONSTRAINT class_student_class_id_fkey FOREIGN KEY (class_id) REFERENCES classmanager.classes(id);


--
-- TOC entry 2764 (class 2606 OID 41117)
-- Name: class_student class_student_student_id_fkey; Type: FK CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.class_student
    ADD CONSTRAINT class_student_student_id_fkey FOREIGN KEY (student_id) REFERENCES classmanager.users(id);


--
-- TOC entry 2762 (class 2606 OID 41099)
-- Name: classes classes_teacher_id_fkey; Type: FK CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.classes
    ADD CONSTRAINT classes_teacher_id_fkey FOREIGN KEY (teacher_id) REFERENCES classmanager.users(id);


--
-- TOC entry 2767 (class 2606 OID 41161)
-- Name: events events_class_id_fkey; Type: FK CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.events
    ADD CONSTRAINT events_class_id_fkey FOREIGN KEY (class_id) REFERENCES classmanager.classes(id);


--
-- TOC entry 2766 (class 2606 OID 41156)
-- Name: events events_year_id_fkey; Type: FK CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.events
    ADD CONSTRAINT events_year_id_fkey FOREIGN KEY (year_id) REFERENCES classmanager.year(id);


--
-- TOC entry 2768 (class 2606 OID 41174)
-- Name: periods periods_year_id_fkey; Type: FK CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.periods
    ADD CONSTRAINT periods_year_id_fkey FOREIGN KEY (year_id) REFERENCES classmanager.year(id);


--
-- TOC entry 2765 (class 2606 OID 41143)
-- Name: public_holidays public_holidays_year_id_fkey; Type: FK CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.public_holidays
    ADD CONSTRAINT public_holidays_year_id_fkey FOREIGN KEY (year_id) REFERENCES classmanager.year(id);


--
-- TOC entry 2769 (class 2606 OID 41188)
-- Name: terms terms_year_id_fkey; Type: FK CONSTRAINT; Schema: classmanager; Owner: emendoza
--

ALTER TABLE ONLY classmanager.terms
    ADD CONSTRAINT terms_year_id_fkey FOREIGN KEY (year_id) REFERENCES classmanager.year(id);


-- Completed on 2019-06-03 18:45:43

--
-- PostgreSQL database dump complete
--

