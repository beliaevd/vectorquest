CREATE TABLE webinars
(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    date TIMESTAMP NOT NULL,
    class_id INTEGER REFERENCES classes(id)
);
CREATE TABLE online 
(
    id SERIAL PRIMARY KEY,
    webinar_id INTEGER REFERENCES webinars(id)
);
DROP TABLE webinar_class_links cascade;
ALTER TABLE tests ADD COLUMN webinar_id INTEGER REFERENCES webinars(id);
DROP TABLE onlines cascade;
ALTER TABLE videos DROP COLUMN test_id;
ALTER TABLE webinar_users DROP COLUMN online_id;
ALTER TABLE webinar_users ADD COLUMN online_id INTEGER REFERENCES online(id);
ALTER TABLE webinar_templates ADD COLUMN calss_id INTEGER REFERENCES classes(id);
ALTER TABLE DROP COLUMN test_id;
