CREATE TABLE todos (
  todo_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  header TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
)