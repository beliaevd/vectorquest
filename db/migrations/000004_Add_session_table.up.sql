CREATE TABLE sessions
    (
        id SERIAL PRIMARY KEY ,
        sessionId text not null ,
        sessionStore text not null,
        expire int not null
    );