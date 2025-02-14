SELECT
    a.id,
    a.username,
    a.locked,
    a.lock_country,
    a.last_ip,
    a.failed_logins,
    ab.unbandate > UNIX_TIMESTAMP() OR ab.unbandate = ab.bandate,
    ab.unbandate = ab.bandate,
    ipb.unbandate > UNIX_TIMESTAMP() OR ipb.unbandate = ipb.bandate,
    ipb.unbandate = ipb.bandate,
    aa.gmlevel,
    a.session_key
FROM
    account a
LEFT JOIN
    account_access aa ON a.id = aa.id
LEFT JOIN
    account_banned ab ON ab.id = a.id AND ab.active = 1
LEFT JOIN
    ip_banned ipb ON ipb.ip = ?
WHERE
    a.username = ? AND
    a.session_key IS NOT NULL
