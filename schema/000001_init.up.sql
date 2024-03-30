CREATE TABLE public.filters (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY,
    kw_id BIGINT,
    filter_name TEXT,
    name TEXT,
    filter_id BIGINT,
    count BIGINT,
    PRIMARY KEY (id)
);

CREATE INDEX idx_filter_name_name ON filters (filter_name, name);


CREATE USER test_user WITH ENCRYPTED PASSWORD 'test_user_031501';
GRANT CONNECT ON DATABASE test_db TO test_user;
GRANT ALL PRIVILEGES ON TABLE public.filters TO test_user;
GRANT ALL ON SEQUENCE public.filters_id_seq TO test_user;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO test_user;
