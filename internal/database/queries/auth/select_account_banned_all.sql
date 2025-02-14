SELECT
    account.id,
    username
FROM
    account,
    account_banned
WHERE
    account.id = account_banned.id AND
    active = 1
GROUP BY
    account.id
