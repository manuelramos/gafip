## Project Database Configuration

This document explains how to configure the database and the necessary environment variables to use the PostgreSQL implementation in this project.

### Prerequisites

Before proceeding with the database configuration, ensure the following:

- PostgreSQL is installed on your system or available on a remote server.
- You have the necessary credentials (username, password, and database name) to access the  PostgreSQL database.

### Environment Variables

The PostgreSQL implementation in this project relies on environment variables to establish a connection to the database. Set the following environment variables with appropriate values before running the project:

- `DB_TYPE`: The type of database to connect to. Values supported postgres and sqlite (default: sqlite).
- `DB_HOST`: The hostname or IP address of the PostgreSQL server.
- `DB_PORT`: The port number on which the PostgreSQL server is running (default: 5432).
- `DB_USER`: The username to authenticate with the PostgreSQL server.
- `DB_PASSWORD`: The password associated with the specified user.
- `DB_NAME`: The name of the database to connect to.

### Setting Environment Variables

On UNIX-based systems (Linux, macOS), you can set the environment variables temporarily for the current session using the export command:

```bash
export DB_TYPE=postgres
export DB_HOST=<your_postgresql_host>
export DB_PORT=<your_postgresql_port>
export DB_USER=<your_postgresql_username>
export DB_PASSWORD=<your_postgresql_password>
export DB_NAME=<your_postgresql_database_name>
```

On Windows, you can set the environment variables using the set command:

```bash
set DB_TYPE=postgres
set DB_HOST=<your_postgresql_host>
set DB_PORT=<your_postgresql_port>
set DB_USER=<your_postgresql_username>
set DB_PASSWORD=<your_postgresql_password>
set DB_NAME=<your_postgresql_database_name>
```

Alternatively, you can store these environment variables in a file (e.g., .env) and load them into your session using a tool like `direnv` or by sourcing the file manually.

### Connecting to PostgreSQL

With the environment variables properly set, the project will use the PostgreSQL implementation to connect to the database using the specified credentials.
