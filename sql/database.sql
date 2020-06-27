
SELECT count(*) from users;

DROP TABLE IF EXISTS users CASCADE;
DROP TYPE IF EXISTS status;

CREATE TYPE status AS ENUM('active', 'expire');

CREATE TABLE users (
 	  id serial,
      name text,
	  status status DEFAULT 'active',
      created_at timestamp DEFAULT NOW(),
      updated_at timestamp DEFAULT NULL
  ) PARTITION BY LIST(updated_at);

CREATE INDEX created_at_inx ON users (created_at);
CREATE INDEX updated_at_inx ON users (updated_at);

CREATE TABLE active_users PARTITION OF users
	FOR VALUES IN (null);

CREATE TABLE dead_users PARTITION OF users
	DEFAULT;

-- ALTER TABLE ONLY users ADD PRIMARY KEY (id);
CREATE UNIQUE INDEX active_users_id_inx ON active_users (id);
CREATE UNIQUE INDEX dead_users_id_inx ON dead_users (id);


-- SEED DATA
INSERT INTO users(name) VALUES('sedhossein');
INSERT INTO users(name, updated_at) VALUES('shahab', NOW());

SELECT * FROM users;
SELECT * FROM active_users;
SELECT * FROM dead_users;


UPDATE users SET updated_at = null where id = 2;
