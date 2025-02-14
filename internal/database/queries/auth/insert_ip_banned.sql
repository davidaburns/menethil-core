INSERT INTO ip_banned (ip, bandate, unbandate, bannedby, banreason)
VALUES
    (?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()+?, ?, ?)
