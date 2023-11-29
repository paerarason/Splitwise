
CREATE TABLE account (
  ID SERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR UNIQUE,
  balance FLOAT
);

CREATE TABLE groups (
  ID SERIAL PRIMARY KEY,
  admin_id   INT REFERENCES account (ID) ON DELETE SET NULL,
  name VARCHAR NOT NULL,
  description TEXT,
  split_for FLOAT NOT NULL
);

CREATE TABLE account_Group (
  ID SERIAL PRIMARY KEY,
  account_id INT REFERENCES account (ID) ON DELETE SET NULL,
  group_id INT REFERENCES groups (ID) ON DELETE SET NULL
);

CREATE TABLE transaction (
  ID SERIAL PRIMARY KEY,
  Account_Group_id INT REFERENCES account_Group (ID) ON DELETE SET NULL,
  amount FLOAT
);
