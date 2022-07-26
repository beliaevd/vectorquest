CREATE TABLE webinar_templates
    (
    id SERIAL PRIMARY KEY ,
    name text,
    description text not  null,
    maxPeople int
    );