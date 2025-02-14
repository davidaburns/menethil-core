SELECT
    a.username, a.last_ip, aa.gmlevel, a.expansion
FROM
    account a
LEFT JOIN
    account_access aa ON (a.id = aa.id)
WHERE
    a.id = ?
ORDER BY
    a.last_ip
