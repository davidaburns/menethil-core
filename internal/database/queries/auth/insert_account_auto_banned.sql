INSERT INTO account_banned
VALUES
    (?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()+?, 'realmd', 'Failed login autoban', 1)
