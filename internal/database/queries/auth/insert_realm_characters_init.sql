INSERT INTO realmcharacters (realmid, acctid, numchars)
SELECT
    realmlist.id,
    account.id,
    0
FROM
    realmlist, account
LEFT JOIN
    realmcharacters ON acctid=account.id
WHERE
    acctid IS NULL
