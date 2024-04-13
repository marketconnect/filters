CREATE TABLE public.filters (
    id SERIAL PRIMARY KEY,
    filter_name TEXT NOT NULL,
    name TEXT NOT NULL,
    UNIQUE (filter_name, name)
);

CREATE INDEX idx_filters_on_filter_name ON filters (filter_name);

CREATE TABLE public.search_phrases (
    id SERIAL PRIMARY KEY,
    phrase TEXT NOT NULL UNIQUE,
    frequency INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_search_phrases_phrase ON search_phrases(phrase);

CREATE USER test_user WITH ENCRYPTED PASSWORD 'filter_user_031501';s
GRANT CONNECT ON DATABASE parsing TO filter_user;
GRANT ALL PRIVILEGES ON TABLE public.filters TO filter_user;
GRANT ALL PRIVILEGES ON TABLE public.search_phrases TO filter_user;
GRANT ALL ON SEQUENCE public.filters_id_seq TO filter_user;
GRANT ALL ON SEQUENCE public.search_phrases_id_seq TO filter_user;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO filter_user;
