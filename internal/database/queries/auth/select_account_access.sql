SELECT
    a.id, aa.gmlevel, aa.RealmID
FROM
    account a
LEFT JOIN
    account_access aa ON (a.id = aa.id)
WHERE
    a.username = ?
