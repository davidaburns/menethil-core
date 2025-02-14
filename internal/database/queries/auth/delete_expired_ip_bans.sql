DELETE
FROM
    ip_banned
WHERE
    unbandate<>bandate AND unbandate<=UNIX_TIMESTAMP()
