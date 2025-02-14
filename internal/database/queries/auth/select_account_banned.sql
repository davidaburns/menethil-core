SELECT
    bandate,
    unbandate
FROM
    account_banned
WHERE
    id = ? AND active = 1
