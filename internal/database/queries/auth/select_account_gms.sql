SELECT
    a.username,
    aa.gmlevel
FROM
    account a,
    account_access aa
WHERE
    a.id=aa.id AND
    aa.gmlevel >= ? AND
    (aa.realmid = -1 OR aa.realmid = ?)
