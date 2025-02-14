SELECT locale, text
FROM motd_localized
WHERE realmid = ? OR realmid = -1
ORDER BY realmid DESC
