ALTER TABLE users
DROP COLUMN role;

ALTER TABLE boards
DROP COLUMN is_block;

ALTER TABLE boards
DROP COLUMN board_admin_id;

ALTER TABLE tasks
DROP COLUMN is_block;

DROP TYPE user_role;