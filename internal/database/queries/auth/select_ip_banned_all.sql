SELECT
    ip,
    bandate,
    unbandate,
    bannedby,
    banreason
FROM
    ip_banned
WHERE
    (bandate = unbandate OR unbandate > UNIX_TIMESTAMP())
ORDER BY
    unbandate
