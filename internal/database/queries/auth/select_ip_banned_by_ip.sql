SELECT
    ip,
    bandate,
    unbandate,
    bannedby,
    banreason
FROM
    ip_banned
WHERE
    (bandate = unbandate OR unbandate > UNIX_TIMESTAMP()) AND
    ip LIKE CONCAT('%%', ?, '%%')
ORDER BY
    unbandate
