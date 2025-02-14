UPDATE account_banned
SET
    active = 0
WHERE
    active = 1 AND unbandate<>bandate AND unbandate<=UNIX_TIMESTAMP()
