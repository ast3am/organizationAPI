CREATE TABLE organization (
    id serial PRIMARY KEY,
    organization_name VARCHAR(255) NOT NULL,
    legal_type VARCHAR(50),
    legal_address VARCHAR(255),
    inn VARCHAR(12),
    owner_id INTEGER  Unique
);

CREATE TABLE filial (
    filial_id SERIAL PRIMARY KEY,
    filial_name VARCHAR(255) NOT NULL,
    country VARCHAR(255),
    city VARCHAR(255),
    address VARCHAR(255),
    filial_type VARCHAR(255),
    phone_list TEXT,
    email_list TEXT,
    photo_id_list TEXT,
    organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
    director_id INT UNIQUE
);

CREATE TABLE employee (
    id serial PRIMARY KEY,
    organization_id integer REFERENCES organization (id) ON DELETE CASCADE NOT NULL,
    filial_id integer REFERENCES filial (filial_id) ON DELETE CASCADE,
    position varchar(50),
    email varchar(50),
    email_confirmation_flag boolean DEFAULT false
);

CREATE TABLE employee_invite (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES employee(id) ON DELETE CASCADE,
    token VARCHAR(255),
    creation_date TIMESTAMP
);

INSERT INTO organization (organization_name, legal_type, legal_address, inn, owner_id)
VALUES ('name', 'ooo', 'some street 123', '098765432110', 10);

INSERT INTO employee (id, organization_id, position, email_confirmation_flag)
VALUES (10, 1, 'owner', true);

INSERT INTO employee (organization_id, position, email_confirmation_flag)
VALUES (1, 'director', true);

