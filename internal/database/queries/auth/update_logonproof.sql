UPDATE account
SET
    session_key = ?,
    last_ip = ?,
    last_login = NOW(),
    locale = ?,
    failed_logins = 0,
    os = ?
WHERE
    username = ?
