DROP TRIGGER IF EXISTS trigger_update_tasks ON tasks;
DROP TRIGGER IF EXISTS trigger_update_users ON users;
DROP TRIGGER IF EXISTS trigger_update_boards ON boards;

DROP FUNCTION IF EXISTS update_timestamp;

DROP TABLE IF EXISTS tasks CASCADE;
DROP TABLE IF EXISTS boards CASCADE;
DROP TABLE IF EXISTS users CASCADE;

DROP TYPE IF EXISTS task_status;
DROP TYPE IF EXISTS board_visibility;