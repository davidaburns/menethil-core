SELECT
    a.username,
    aa.gmlevel,
    a.email,
    a.reg_mail,
    a.last_ip,
    DATE_FORMAT(a.last_login, '%Y-%m-%d %T'),
    a.mutetime,
    a.mutereason,
    a.muteby,
    a.failed_logins,
    a.locked,
    a.OS
FROM
    account a
LEFT JOIN
    account_access aa ON (a.id = aa.id AND (aa.RealmID = ? OR aa.RealmID = -1))
WHERE
    a.id = ?
