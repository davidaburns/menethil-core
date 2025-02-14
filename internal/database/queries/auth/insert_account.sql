INSERT INTO account(username, salt, verifier, expansion, joindate)
VALUES
    (?, ?, ?, ?, NOW())
