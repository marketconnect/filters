CREATE TEMP TABLE temp_kw_updates (
    kw TEXT,
    freq INTEGER
);

COPY temp_kw_updates(kw, freq) FROM '/tmp/requests.csv' DELIMITER ',' CSV HEADER;

UPDATE kw                                                                        
SET freq = temp.freq
FROM temp_kw_updates temp
WHERE kw.kw = temp.kw;

DROP TABLE temp_kw_updates;
