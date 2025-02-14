SELECT mutedate, mutetime, mutereason, mutedby
FROM account_muted
WHERE guid = ?
ORDER BY mutedate ASC
