/* insert Group values for Testing */
INSERT INTO account (username, password, email, balance)
VALUES ('user1', '$2a$14$VlvQ/zqvXL.vbhJ6P1iDbeYCXKQpyt9D.ao8gjOqClt16CnXBx28a', 'user1@example.com', 1000);

INSERT INTO account (username, password, email, balance)
VALUES ('user2', '$2a$14$a0dsFk203tJzlkzW5kaUwejMG5roMVhhXcexZ9265b9AnADknJs4G', 'user2@example.com', 500);

INSERT INTO account (username, password, email, balance)
VALUES ('user3', '$2a$14$bNiOj4cfl0Fincr2hY3oPOJrYx.zsSmZqgE.R.TuRdZrQo1wGehWO', 'user3@example.com', 2000);

INSERT INTO account (username, password, email, balance)
VALUES ('user4', '$2a$14$CSjB9QTIchD2ALSZXI.3OuZ7LdILd0/e76JaBYUWZdyXbxfn9Df7.', 'user4@example.com', 1500);

INSERT INTO account (username, password, email, balance)
VALUES ('user5', '$2a$14$.XEsC4q4Q4B5LNqLNODXEeUUqPLo5BTkhT4ed/EvSGXymOwGqCN9W', 'user5@example.com', 800);


/* insert Group values for Testing */
INSERT INTO groups (admin_id, name, description, split_for)
VALUES (1, 'Group 1', 'This is the first group', 4000);

INSERT INTO groups (admin_id, name, description, split_for)
VALUES (2, 'Group 2', 'This is the second group', 500);

INSERT INTO groups (admin_id, name, description, split_for)
VALUES (3, 'Group 3', 'This is the third group', 800);


/* insert account_group  values for Testing */
INSERT INTO account_Group (account_id, group_id)
VALUES (1, 1);

INSERT INTO account_Group (account_id, group_id)
VALUES (2, 2);

INSERT INTO account_Group (account_id, group_id)
VALUES (3, 3);

INSERT INTO account_Group (account_id, group_id)
VALUES (4, 1);

INSERT INTO account_Group (account_id, group_id)
VALUES (5, 2);

/* insert transaction  values for Testing */
INSERT INTO transaction (Account_Group_id, amount)
VALUES (1, 100);

INSERT INTO transaction (Account_Group_id, amount)
VALUES (2, 50);

INSERT INTO transaction (Account_Group_id, amount)
VALUES (3, 200);

INSERT INTO transaction (Account_Group_id, amount)
VALUES (4, 150);

INSERT INTO transaction (Account_Group_id, amount)
VALUES (5, 80);
