--1. create user table 
CREATE TABLE users(
  id UUID PRIMARY KEY,
  email TEXT NOT NULL,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz DEFAULT CURRENT_TIMESTAMP
);
--2. Trigger to auto update update_at timestamp
CREATE
OR REPLACE FUNCTION update_modified_column() RETURNS TRIGGER AS $$BEGIN NEW.updated_at = now();
RETURN NEW;

END;
$$LANGUAGE 'plpgsql';
-- 3. Create the Trigger
CREATE TRIGGER update_user_modtime BEFORE UPDATE 
ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
