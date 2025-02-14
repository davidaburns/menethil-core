SELECT
    unbandate > UNIX_TIMESTAMP() OR unbandate = bandate AS banned,
    NULL as country
FROM
    ip_banned
WHERE
    ip = ?
