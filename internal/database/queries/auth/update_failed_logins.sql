UPDATE account
SET
    failed_logins = failed_logins + 1
WHERE
    username = ?
