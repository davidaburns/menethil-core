SELECT unbandate, bandate = unbandate, bannedby, banreason 
FROM account_banned 
WHERE id = ? AND active 
ORDER BY bandate ASC 
LIMIT 1