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

CREATE TABLE Employment (
    ID_Employment INT PRIMARY KEY,
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
