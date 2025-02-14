SELECT
    a.id,
    a.session_key,
    a.last_ip,
    a.locked,
    a.lock_country,
    a.expansion,
    a.mutetime,
    a.locale,
    a.recruiter,
    a.os,
    a.totaltime,
    aa.gmlevel,
    ab.unbandate > UNIX_TIMESTAMP() OR ab.unbandate = ab.bandate,
    r.id
FROM
    account a
LEFT JOIN
    account_access aa ON a.id = aa.id AND aa.RealmID IN (-1, ?)
LEFT JOIN
    account_banned ab ON a.id = ab.id AND ab.active = 1
LEFT JOIN
    account r ON a.id = r.recruiter
WHERE
    a.username = ? AND a.session_key IS NOT NULL
ORDER BY
    aa.RealmID DESC
LIMIT 1
