-- ========= DEFINIZIONE DEI TIPI ENUMERATIVI =========
-- Mettiamo tutte le definizioni dei tipi all'inizio per chiarezza.

CREATE TYPE role_enum AS ENUM ('Assistant', 'Seller', 'Manager', 'AreaManager', 'ADMIN');
CREATE TYPE vehicle_condition AS ENUM ('New', 'Used');
CREATE TYPE order_status_enum AS ENUM ('Pending', 'Completed', 'Cancelled', 'InProgress');
CREATE TYPE client_type_enum AS ENUM ('Private', 'Business');


-- ========= CREAZIONE DELLE TABELLE =========

CREATE TABLE "Dealership" (
    ID_Dealership SERIAL PRIMARY KEY,
    PostalCode VARCHAR(5) NOT NULL,
    City VARCHAR(30) NOT NULL,
    Address VARCHAR(100) NOT NULL,
    Phone VARCHAR(20)
);

CREATE TABLE "Employee" (
    ID_Employee SERIAL PRIMARY KEY,
    TIN VARCHAR(16) UNIQUE NOT NULL,
    "Name" VARCHAR(50) NOT NULL,
    "Surname" VARCHAR(50) NOT NULL,
    Phone VARCHAR(20),
    "Role" role_enum NOT NULL
);

CREATE TABLE "Employment" (
    ID_Employment SERIAL PRIMARY KEY,
    ID_Employee INT NOT NULL,
    ID_Dealership INT NOT NULL,
    StartDate DATE NOT NULL,
    EndDate DATE,
    FOREIGN KEY (ID_Employee) REFERENCES "Employee"(ID_Employee),
    FOREIGN KEY (ID_Dealership) REFERENCES "Dealership"(ID_Dealership)
);

CREATE TABLE "CarPark" (
    VIN VARCHAR(17) PRIMARY KEY,
    ID_Dealership INT NOT NULL,
    Brand VARCHAR(30) NOT NULL,
    Model VARCHAR(30) NOT NULL,
    "Year" INT NOT NULL,
    Notes TEXT,
    "Condition" vehicle_condition NOT NULL DEFAULT 'New',
    KM INT,
    Incidents TEXT,
    OilChange DATE,
    FOREIGN KEY (ID_Dealership) REFERENCES "Dealership"(ID_Dealership)
);

-- Tabella Client unificata (approccio "Single Table Inheritance")
CREATE TABLE "Client" (
    ID_Client SERIAL PRIMARY KEY,
    ClientType client_type_enum NOT NULL,
    Phone VARCHAR(20),
    Email VARCHAR(50) UNIQUE,
    -- Campi per clienti privati (possono essere NULL)
    FiscalCode VARCHAR(16) UNIQUE,
    FirstName VARCHAR(50),
    LastName VARCHAR(50),
    -- Campi per clienti business (possono essere NULL)
    VATNumber VARCHAR(11) UNIQUE,
    CompanyName VARCHAR(100)
);

CREATE TABLE "Appointment" (
    ID_Appointment SERIAL PRIMARY KEY,
    ID_Employee INT,
    ID_Client INT,
    VIN VARCHAR(17),
    AppointmentDate TIMESTAMP NOT NULL, -- Usiamo TIMESTAMP per includere data e ora
    Reason VARCHAR(256),
    FOREIGN KEY (ID_Employee) REFERENCES "Employee"(ID_Employee),
    FOREIGN KEY (ID_Client) REFERENCES "Client"(ID_Client),
    FOREIGN KEY (VIN) REFERENCES "CarPark"(VIN)
);

CREATE TABLE "Order" (
    ID_Order SERIAL PRIMARY KEY,
    "Status" order_status_enum NOT NULL DEFAULT 'InProgress',
    ID_Client INT NOT NULL,
    ID_Dealership INT NOT NULL,
    VIN VARCHAR(17) NOT NULL,
    ID_Employee INT NOT NULL,
    "OrderDate" DATE,
    "LastUpdate" DATE,
    FOREIGN KEY (ID_Client) REFERENCES "Client"(ID_Client),
    FOREIGN KEY (ID_Dealership) REFERENCES "Dealership"(ID_Dealership),
    FOREIGN KEY (VIN) REFERENCES "CarPark"(VIN),
    FOREIGN KEY (ID_Employee) REFERENCES "Employee"(ID_Employee)
);