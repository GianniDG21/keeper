CREATE TABLE Dealership (
    ID_Dealership INT PRIMARY KEY,
    PostalCode VARCHAR(5),
    City VARCHAR(30),
    Address VARCHAR(100)
    Phone VARCHAR(20)
)

CREATE TABLE Employee (
    ID_Employee INT PRIMARY KEY,
    TIN VARCHAR(16),
    Name VARCHAR(50),
    Surname VARCHAR(50),
    Phone VARCHAR(20),
    Role RoleEnum 
)

CREATE TYPE role AS ENUM ('Assistant', 'Seller', 'Manager', 'AreaManager', 'ADMIN')
CREATE TABLE Employment (
    ID_Employment INT PRIMARY KEY,
    Role role NOT NULL DEFAULT 'INSERT ROLE',
    ID_Employee INT,
    ID_Dealership INT,
    StartDate DATE,
    EndDate DATE NULL,
    FOREIGN KEY (ID_Employee) REFERENCES Employee(ID_Employee),
    FOREIGN KEY (ID_Dealership) REFERENCES Dealership(ID_Dealership)
)

CREATE TYPE condition AS ENUM ('New', 'Used');
CREATE TABLE CarPark (
    VIN VARCHAR(17) PRIMARY KEY,
    ID_Dealership INT,
    Brand VARCHAR(30),
    Model VARCHAR(30),
    Year VARCHAR(4),
    Notes TEXT,
    Condition condition NOT NULL DEFAULT 'New',
    KM VARCHAR(7),
    Incidents TEXT NULL,
    OilChange DATE NULL,
    FOREIGN KEY (ID_Dealership) REFERENCES Dealership(ID_Dealership)
)

CREATE TABLE Appointment (
    ID_Appointment INT PRIMARY KEY,
    ID_Employee INT,
    VIN VARCHAR(17),
    AppointmentDate DATE,
    Reason VARCHAR(256),
    FOREIGN KEY (ID_Employee) REFERENCES Employee(ID_Employee),
    FOREIGN KEY (VIN) REFERENCES CarPark(VIN)
)
-- Definition and specialization for Client --
CREATE TABLE Client (
    ID_Client INT PRIMARY KEY,
    Phone VARCHAR(20),
    Email VARCHAR(50)
)
CREATE TABLE ClientPrivate (
    ID_ClientPrivate INT PRIMARY KEY,
    ID_Client INT REFERENCES Client(ID_Client),
    FiscalCode VARCHAR(16)
);

CREATE TABLE ClientBusiness (
    ID_ClientBusiness INT PRIMARY KEY,
    ID_Client INT REFERENCES Client(ID_Client),
    VATNumber VARCHAR(11),
    CompanyName VARCHAR(100)
);
CREATE TABLE Order (
    ID_Order INT PRIMARY KEY,
    ID_Client INT REFERENCES Client(ID_Client),
    OrderDate DATE
);
-- End of Client Definition and Specialization --

CREATE TYPE status AS ENUM ('Pending', 'Completed', 'Cancelled', 'InProgress');
CREATE TABLE Order (
    ID_Order INT PRIMARY KEY,
    Status status NOT NULL DEFAULT 'InProgress',
    ID_Client INT REFERENCES Client(ID_Client),
    ID_Dealership INT REFERENCES Dealership(ID_Dealership),
    VIN VARCHAR(17) REFERENCES CarPark(VIN),
    ID_Employee INT REFERENCES Employee(ID_Employee),
    OrderDate DATE
    LastUpdate DATE
);