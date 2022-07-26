CREATE TABLE schools
(
    id SERIAL PRIMARY KEY ,
    name varchar(1000) not null
);
CREATE TABLE classes(
    id SERIAL PRIMARY KEY ,
    number integer not null
);

CREATE TABLE webinars
(
    id SERIAL PRIMARY KEY ,
    name varchar(255) NOT NULL,
    date TIMESTAMP NOT NULL,
    class_id INTEGER REFERENCES classes(id)
);
CREATE TABLE videos
(
    id SERIAL PRIMARY KEY ,
    name varchar(8000) NOT NULL
);
CREATE TABLE tests
(
    id SERIAL PRIMARY KEY ,
    name varchar(255) NOT NULL ,
    webinar_id INTEGER REFERENCES webinars(id)
);
CREATE TABLE online
(
    id SERIAL PRIMARY KEY ,
    webinar_id INTEGER REFERENCES webinars(id)
);

CREATE TABLE question_categories
(
    id SERIAL PRIMARY KEY,
    name varchar(45)
);
CREATE TABLE question_values
(
    id SERIAL PRIMARY KEY ,
    value text NOT NULL
);
CREATE TABLE questions
(
    id SERIAL PRIMARY KEY ,
    number integer NOT NULL ,
    category INTEGER REFERENCES question_categories(id),
    value INTEGER REFERENCES question_values(id),
    test_id INTEGER REFERENCES tests(id)
);
CREATE TABLE ranks
(
    id SERIAL PRIMARY KEY,
    name varchar(45)
);

CREATE TABLE users
(
    id SERIAL PRIMARY KEY ,
    name varchar(45) NOT NULL ,
    surname varchar(45) NOT NULL ,
    middle_name varchar(45),
    login varchar(255) NOT NULL ,
    email varchar(255) NOT NULL ,
    password varchar(255) NOT NULL ,
    level integer NOT NULL ,
    status text NOT NULL default 'Активный',
    school_id INTEGER REFERENCES schools(id) ,
    class_id INTEGER REFERENCES classes(id),
    rank_id INTEGER REFERENCES ranks(id)
);
CREATE TABLE webinar_users
(
    id SERIAL PRIMARY KEY ,
    user_id INTEGER REFERENCES users(id),
    webinar_id INTEGER REFERENCES webinars(id)
);

CREATE TABLE admin_user_categories
(
    id SERIAL PRIMARY KEY ,
    name varchar(45)
);

CREATE TABLE admin_users
(
    id SERIAL PRIMARY KEY ,
    name varchar(45) NOT NULL ,
    surname varchar(45) NOT NULL ,
    login varchar(255) NOT NULL ,
    email varchar(255) ,
    password varchar(255) NOT NULL,
    status varchar(1) DEFAULT 'N'
);

INSERT INTO classes VALUES (1,5),(2,6),(3,7),(4,8),(5,9),(6,10),(7,11);
INSERT INTO admin_user_categories VALUES (1,'Ведущий'),(2,'Судья'),(3,'root'),(4,'Creater');
