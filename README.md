
# Go Lego API

This project is a simple Go API that connects to a PostgreSQL database and provides an endpoint to fetch data about Lego sets. The data is retrieved from the `lego_table` in the PostgreSQL database.

## Project Structure

```
go-lego-api/
├── api/
│   └── handler.go          # API entry point (package handler, required for Vercel)
├── sql/
|   └── lego_sets.sql    # SQL file with schema and sample LEGO data
├── .env                 # Environment variables (e.g., PostgreSQL URL)
├── README.md            # Project overview and setup instructions
```

## Setup Instructions

1. Clone the repository:

```bash
git clone https://github.com/jabercrombia/go-lego-api.git
cd go-lego-api
```

2. Create a `.env` file at the root of the project and add your PostgreSQL connection URL:
   
   ```
   POSTGRES_URL=your_postgres_connection_url
   ```

3. Make sure you have Go and PostgreSQL installed.

4. Once the server is running, you can access the API at `http://localhost:8080/legosets`.

## Lego Table Schema (`lego_sets.sql`)

Here is the SQL schema to set up the `lego_table` in PostgreSQL:

```sql
CREATE TABLE lego_table (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  theme VARCHAR(255),
  thumbnailurl VARCHAR(255)
);

## API Endpoints

- `GET /legosets`: Fetches a list of Lego sets from the database.

## License

This project is open-source and available under the [MIT License](LICENSE).
