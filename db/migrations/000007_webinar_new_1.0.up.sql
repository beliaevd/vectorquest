CREATE TABLE webinar_class_links 
(
    id SERIAL PRIMARY KEY,
    class_id INTEGER REFERENCES classes(id),
    webinar_templates_id INTEGER REFERENCES webinar_templates(id)
);

ALTER TABLE tests DROP COLUMN webinar_id;
ALTER TABLE online DROP COLUMN webinar_id;
DROP TABLE webinars cascade;
ALTER TABLE webinar_users DROP COLUMN online_id;
ALTER TABLE webinar_templates ADD COLUMN test_id INTEGER REFERENCES tests(id);
DROP TABLE online cascade;
CREATE TABLE onlines 
(
    id SERIAL PRIMARY KEY,
    leader INTEGER REFERENCES admin_users(id),
    judge INTEGER REFERENCES admin_users(id),
    webinar_templates_id INTEGER REFERENCES webinar_templates(id)
);
ALTER TABLE webinar_users ADD COLUMN online_id INTEGER REFERENCES onlines(id);
ALTER TABLE videos ADD COLUMN test_id INTEGER REFERENCES tests(id);
