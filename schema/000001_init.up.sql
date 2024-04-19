CREATE TABLE public.filters (
    id SERIAL PRIMARY KEY,
    filter_name TEXT NOT NULL,
    name TEXT NOT NULL,
    UNIQUE (filter_name, name)
);

CREATE INDEX idx_filters_on_filter_name ON filters (filter_name);

CREATE TABLE public.search_phrases (
    id SERIAL PRIMARY KEY,
    kw TEXT NOT NULL UNIQUE,
    freq INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_search_phrases_kw ON search_phrases(kw);

CREATE TABLE kw (
    id SERIAL PRIMARY KEY,        
    name TEXT NOT NULL,          
    normquery TEXT NOT NULL,     
    cards_qty INTEGER NOT NULL   
);
CREATE INDEX idx_kw_on_normquery ON kw (normquery);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    kw_id INTEGER NOT NULL,
    filter_id INTEGER NOT NULL,
    count INTEGER NOT NULL,    
    FOREIGN KEY (kw_id) REFERENCES kw (id)   
);
CREATE INDEX idx_categories_on_kw_id ON categories (kw_id);
CREATE INDEX idx_categories_on_filter_id ON categories (filter_id);


CREATE TABLE lemmas (
    id SERIAL PRIMARY KEY,
    lemma TEXT NOT NULL UNIQUE
);

CREATE TABLE kw_lemmas (
    kw_id INTEGER NOT NULL,
    lemma_id INTEGER NOT NULL,
    FOREIGN KEY (kw_id) REFERENCES kw (id),
    FOREIGN KEY (lemma_id) REFERENCES lemmas (id),
    PRIMARY KEY (kw_id, lemma_id)
);


CREATE USER test_user WITH ENCRYPTED PASSWORD 'test_user_031501';
GRANT CONNECT ON DATABASE parsing TO test_user;
GRANT ALL PRIVILEGES ON TABLE public.filters TO test_user;
GRANT ALL PRIVILEGES ON TABLE public.kw TO test_user;
GRANT ALL PRIVILEGES ON TABLE public.categories TO test_user;
GRANT ALL PRIVILEGES ON TABLE public.search_phrases TO test_user;
GRANT ALL PRIVILEGES ON TABLE public.kw_lemmas TO test_user;
GRANT ALL PRIVILEGES ON TABLE public.lemmas TO test_user;
GRANT ALL ON SEQUENCE public.filters_id_seq TO test_user;
GRANT ALL ON SEQUENCE public.kw_id_seq TO test_user;
GRANT ALL ON SEQUENCE public.categories_id_seq TO test_user;
GRANT ALL ON SEQUENCE public.search_phrases_id_seq TO test_user;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO test_user;

