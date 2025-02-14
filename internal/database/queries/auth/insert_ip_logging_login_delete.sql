INSERT INTO logs_ip_actions (account_id,character_guid,type,ip,systemnote,unixtime,time)
VALUES
    (?, ?, ?, (SELECT last_ip FROM account WHERE id = ?), ?, unix_timestamp(NOW()), NOW())
