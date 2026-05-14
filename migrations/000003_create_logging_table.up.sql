CREATE TABLE water_logs(
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  water_intake_quantity INT NOT NULL,
  created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_water_logs_modtime
BEFORE UPDATE ON water_logs 
FOR EACH ROW EXECUTE FUNCTION update_modified_column();
