CREATE TABLE department(
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE employee(
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    department INTEGER REFERENCES department(id),
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL
);