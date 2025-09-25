# KEEPER

> A microservice-based backend API for inventory management, designed for Small and Medium-sized Businesses (SMBs).

This project is the result of a significant learning journey into modern backend development, including microservice architecture, containerization, and agile methodologies.

---

###  ABOUT THE PROJECT

**KEEPER** is a backend system designed to solve a common problem for multi-location businesses: decentralized and inefficient inventory management.

This implementation is modeled for **DGAuto**, a fictional multi-brand car dealership with several branches. Currently, each branch manages its own vehicle stock, leading to synchronization issues and a lack of a unified, real-time view of the entire company's inventory.

KEEPER serves as the **central nervous system** for the dealership's inventory, providing a reliable, single source of truth through a robust REST API. The design considers a role-based access control (RBAC) system to ensure that data can only be modified by authorized personnel (e.g., Sales Managers, Assistants), guaranteeing data integrity and security.

---

### BUILT WITH

The technology stack for KEEPER was chosen to ensure efficiency, scalability, and maintainability.

#### Back-End
* **Go (Golang)**: Chosen for its high performance in concurrent environments and strong typing, making it ideal for building efficient and reliable API services.
* **GORM**: The most popular ORM for Go. It was preferred over query builders like `sqlc` to gain hands-on experience with a full-featured ORM.
* **Python**: Used for the secondary microservice responsible for handling notifications, demonstrating a polyglot architecture.
* **Peewee**: A simple and lightweight ORM for Python. It was chosen over SQLAlchemy for its simplicity, given the notification service's limited database interaction, without compromising future scalability.
* **PostgreSQL**: A powerful, open-source relational database. It was selected for its robustness in handling large volumes of data, complex queries, and mixed read/write workloads, making it highly scalable.
* **REST API**: The architectural standard used for communication between services, with **JSON** as the data format.

#### DevOps & Containerization
* **Docker & Docker Compose**: Docker is used to create isolated containers for each component (Go API, Python Notifier, PostgreSQL DB). Docker Compose orchestrates the entire multi-container application, allowing it to be run with a single command.

### FUTURE IMPLEMENTATIONS
This project lays the foundation for a complete Dealership Management System (DMS). While the core features are implemented, I have a clear vision for version 2.0. Here are some of the key features planned for future development:

#### Advanced Inventory Management (Beyond CRUD)
These features cover the entire vehicle lifecycle, not just its presence in the warehouse.

**Vehicle Acquisition Management**: This use case handles how a vehicle enters the inventory, such as through a trade-in from a customer, an auction purchase, or a transfer from another branch.
Maintenance and Reconditioning Tracking: This use case tracks the costs and status of interventions that a used car often requires, such as inspections, repairs, or detailing, before it can be sold.
**Vehicle Document Management**: Associates each vehicle with its digitized documents, like the vehicle registration, service history, and photos for the website.

#### Sales and Financial Flows
These use cases cover the economic and contractual aspects of a sale.

**Financing Application Management**: This use case manages the collection of customer documents, the submission of the application to the financial company, and the tracking of its outcome.
**Quote Calculation and Generation**: Allows a salesperson to create a detailed price quote for a customer, including the vehicle price, optional features, registration costs, and any trade-in value.
**Test Drive Management**: Formalizes the management of test drives by logging customer data, the vehicle's departure and return times, and collecting post-drive feedback.

#### After-Sales and CRM (Customer Relationship Management)
These features help maintain the customer relationship after the sale, a crucial phase for a dealership's business.

**Service Department Appointment Management**: Allows for booking service and repair appointments for customers who have purchased a vehicle, managing the service center's schedule.
**Customer Communication Management**: Tracks all communications with a customer and allows for sending automated reminders.

---

### ACKNOWLEDGMENTS

If you've taken the time to review this project, thank you. This marks my first major portfolio piece and, I hope, the first of many more to come. It represents months of dedicated effort to acquire the skills necessary to start my career in software development.

---

### DESIGN DECISIONS

#### Data Return Strategy

For API responses and database queries, this project returns slices of struct values (e.g., `[]models.Dealership`) instead of slices of pointers (e.g., `[]*models.Dealership`).  
This choice ensures that API consumers cannot accidentally modify the original data, improving safety and predictability.  
While using pointers can be more efficient for very large structs, in this context the data structures are small enough that the safety and clarity benefits outweigh the minor performance cost.

#### Hibryd Data Access Layer

A deliberate hybrid approach was chosen for the project's data access layer to demonstrate proficiency in two different database interaction methodologies.

- Low-Level Control (Standard database/sql Package)

The core CRUD functionalities for primary entities (e.g., Dealership, CarPark) were implemented using Go's standard database/sql package. This approach was taken to showcase a fundamental understanding of low-level database interaction, including full control over SQL queries and manual data mapping.

- High-Level Abstraction (GORM)

For more standard and repetitive CRUD operations on other entities, the GORM library was integrated. This choice demonstrates the ability to leverage a modern ORM to increase development speed, reduce boilerplate code, and improve maintainability, simulating a real-world production environment where productivity is key.

This hybrid strategy showcases the ability to choose the right tool for the job, balancing fine-grained control with high-level abstraction.

#### Advanced Routing with chi
While Go's standard http.ServeMux is functional, the chi router was chosen for this project to provide a more powerful and organized routing layer. This choice allows for the logical grouping of related endpoints (e.g., all /employees routes are defined in a single block), which significantly improves the readability and maintainability of the server.go file. Furthermore, chi provides a simple and effective middleware system, used here for request logging and panic recovery, which are essential features for a production-ready API.

#### Declarative Request Validation
To ensure data integrity and keep handlers clean, request validation is handled by the go-playground/validator library. Instead of cluttering HTTP handlers with repetitive if/else blocks for each field, validation rules are declaratively defined using validate tags directly on the model structs. This co-locates a data structure with its validation rules, making the system easier to understand. Handlers remain focused on their core business logic, simply calling the validator on the decoded struct to enforce all rules in a single, clean step.
