CREATE TABLE tasks(
                      id SERIAL PRIMARY KEY,
                      status TEXT,
                      name TEXT,
                      description TEXT,
                      performer TEXT,
                      deadline TIMESTAMP
)