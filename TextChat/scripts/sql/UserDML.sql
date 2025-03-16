-- The procedures for User

-- InsertUser
INSERT INTO User (username, password) VALUES (?, ?);

-- SelectUser
SELECT username, time_created FROM User WHERE id = ?;
