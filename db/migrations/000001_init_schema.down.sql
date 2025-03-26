DROP TRIGGER IF EXISTS trigger_update_tasks ON tasks;
DROP TRIGGER IF EXISTS trigger_update_users ON users;
DROP FUNCTION IF EXISTS update_timestamp;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS boards;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS task_status;