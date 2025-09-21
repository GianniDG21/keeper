create type client_type_enum as enum ('private', 'company');

CREATE Table "client" (
    id_client SERIAL PRIMARY KEY,
    "type" client_type_enum NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(50) UNIQUE,
    tin_vat VARCHAR(16) UNIQUE,
    name VARCHAR(50),
    surname VARCHAR(50),
    companyname VARCHAR(100) DEFAULT NULL,
    profession VARCHAR(50) DEFAULT NULL
);
CREATE Table dealership (
    id_dealership SERIAL PRIMARY KEY,
    postalcode VARCHAR(5) NOT NULL,
    city VARCHAR(30) NOT NULL,
    address VARCHAR(100) NOT NULL,
    phone VARCHAR(20)
);
create type role_enum as enum ('Manager', 'Mechanic', 'Salesperson', 'Assistant', 'ADMIN');
CREATE Table employee (
    id_employee SERIAL PRIMARY KEY,
    role role_enum NOT NULL DEFAULT 'Assistant',
    tin VARCHAR(16) UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    phone VARCHAR(20)
);

CREATE Table employment (
    id_employment SERIAL PRIMARY KEY,
    id_employee INT NOT NULL,
    id_dealership INT NOT NULL,
    startdate DATE NOT NULL,
    enddate DATE,
    FOREIGN KEY (id_employee) REFERENCES employee(id_employee),
    FOREIGN KEY (id_dealership) REFERENCES dealership(id_dealership)
);
create type condition_enum as enum ('New', 'Used');
create table car_park (
    vin VARCHAR(17) PRIMARY KEY,
    id_dealership INT NOT NULL,
    brand VARCHAR(30) NOT NULL,
    model VARCHAR(30) NOT NULL,
    condition condition_enum NOT NULL DEFAULT 'New',
    "year" INT NOT NULL CHECK ("year" > 1900 AND "year" <= (EXTRACT(YEAR FROM CURRENT_DATE) + 1)), 
    km int NOT NULL DEFAULT 0
);

create type status_enum as enum ('Pending', 'Completed', 'Cancelled', 'InProgress');
create table "order" (
    id_order SERIAL PRIMARY KEY,
    status status_enum NOT NULL DEFAULT 'Pending',
    id_client INT NOT NULL,
    id_employee INT NOT NULL,
    vin VARCHAR(17) NOT NULL,
    id_dealership INT NOT NULL,
    last_update DATE NOT NULL,
    FOREIGN KEY (id_client) REFERENCES client(id_client),
    FOREIGN KEY (id_employee) REFERENCES employee(id_employee),
    FOREIGN KEY (vin) REFERENCES car_park(vin),
    FOREIGN KEY (id_dealership) REFERENCES dealership(id_dealership)
);

create table appointment (
    id_appointment SERIAL PRIMARY KEY,
    id_client INT NOT NULL,
    id_employee INT NOT NULL,
    id_dealership INT NOT NULL,
    "date" TIMESTAMP NOT NULL,
    reason VARCHAR(100) NOT NULL,
    notes TEXT,
    FOREIGN KEY (id_client) REFERENCES client(id_client),
    FOREIGN KEY (id_employee) REFERENCES employee(id_employee),
    FOREIGN KEY (id_dealership) REFERENCES dealership(id_dealership)
);

