DELETE FROM account_access
WHERE
    id = ? AND
    (RealmID = ? OR RealmID = -1)
