SELECT 1 FROM account_banned
WHERE
    id = ? AND active = 1
UNION
SELECT 1 FROM ip_banned
WHERE
    ip = ?
