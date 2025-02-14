SELECT
    id,
    failed_logins
FROM
    account
WHERE
    username = ?
