# KEEPER

> A microservice-based backend API for inventory management, designed for Small and Medium-sized Businesses (SMBs).

This project is the result of a significant learning journey into modern backend development, including microservice architecture, containerization, and agile methodologies.

---

### ABOUT THE PROJECT

**KEEPER** is a backend system designed to solve a common problem for multi-location businesses: decentralized and inefficient inventory management.

This implementation is modeled for **DGAuto**, a fictional multi-brand car dealership with several branches. Currently, each branch manages its own vehicle stock, leading to synchronization issues and a lack of a unified, real-time view of the entire company's inventory.

KEEPER serves as the **central nervous system** for the dealership's inventory, providing a reliable, single source of truth through a robust REST API. The design considers a role-based access control (RBAC) system to ensure that data can only be modified by authorized personnel, guaranteeing data integrity and security.

---

### BUILT WITH

The technology stack for KEEPER was chosen to ensure efficiency, scalability, and maintainability.

#### Backend
* **Go (Golang)**: Chosen for its high performance in concurrent environments and strong typing, making it ideal for building efficient and reliable API services.
* **Python**: Used for the secondary microservice responsible for handling notifications, demonstrating a polyglot architecture.
* **PostgreSQL**: A powerful, open-source relational database.
* **GORM**: The most popular ORM for Go, used for productivity.
* **Chi**: A lightweight and idiomatic router for building structured and maintainable Go APIs.
* **go-playground/validator**: A library for declarative, tag-based struct validation.

#### DevOps & Containerization
* **Docker & Docker Compose**: Used to create a consistent, reproducible development and production environment for all services.

---

### FUTURE IMPLEMENTATIONS

This project lays the foundation for a complete Dealership Management System (DMS). While the core features are implemented, I have a clear vision for version 2.0. Here are some of the key features planned for future development:

#### Advanced Inventory Management (Beyond CRUD)
**Vehicle Acquisition Management**: Handles how a vehicle enters the inventory (e.g., trade-in, auction purchase).
**Maintenance and Reconditioning Tracking**: Tracks costs and status of interventions on used cars before they can be sold.
**Vehicle Document Management**: Associates each vehicle with its digitized documents.

#### Sales and Financial Flows
**Financing Application Management**: Manages the collection and submission of customer documents for financing.
**Quote Calculation and Generation**: Allows a salesperson to create detailed price quotes.
**Test Drive Management**: Formalizes the management of test drives, including logging and feedback.

#### After-Sales and CRM
**Service Department Appointment Management**: Manages the service center's schedule for repairs and maintenance.
**Customer Communication Management**: Tracks all communications and allows for sending automated reminders.

---

### DESIGN DECISIONS

#### Hybrid Data Access Layer
A deliberate hybrid approach was chosen for the data access layer to demonstrate proficiency in two methodologies:
- **Low-Level Control (`database/sql`)**: Core functionalities for key entities were implemented using Go's standard library to showcase a fundamental understanding of SQL interaction and manual data mapping.
- **High-Level Abstraction (GORM)**: For more standard CRUD operations, the GORM library was integrated to increase development speed and reduce boilerplate, simulating a real-world production environment.

#### Advanced Routing with `chi`
The `chi` router was chosen over the standard library's `ServeMux` to provide a more powerful and organized routing layer. Key benefits include logical route grouping and a simple middleware system, used here for request logging and panic recovery.

#### Declarative Request Validation
To ensure data integrity, request validation is handled by the `go-playground/validator` library. Instead of cluttering HTTP handlers with repetitive `if/else` blocks, validation rules are declaratively defined using `validate` tags directly on the model structs.

#### Database Key Strategy (Surrogate vs. Natural Keys)
While some entities have a "natural" unique key (like the `VIN` for a vehicle), a "surrogate key" strategy was adopted for all primary keys (e.g., `id_car_park SERIAL PRIMARY KEY`). This choice ensures stable integer-based foreign key relationships, improves `JOIN` performance, and leads to cleaner RESTful API endpoints (e.g., `/vehicles/123`). Natural keys like `VIN` are maintained with a `UNIQUE` constraint.

#### Security Analysis & Static Tooling (CSRF False Positive)
The project was configured for static analysis with SonarQube. The tool correctly identified that the Python/Flask microservice, by default, does not have CSRF protection. After analysis, this was determined to be a **false positive** and marked as such. The notifier is a backend-to-backend service with no browser-based user sessions, which are the attack vectors for CSRF. This highlights a pragmatic approach to security: leveraging automated tools while also applying contextual analysis to distinguish real risks from theoretical ones.

---

### ACKNOWLEDGMENTS

Thank you for taking the time to review this project. It marks my first major portfolio piece and, I hope, the first of many more to come. It represents months of dedicated effort to acquire the skills necessary to start my career in software development.