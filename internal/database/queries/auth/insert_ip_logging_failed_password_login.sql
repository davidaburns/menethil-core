INSERT INTO logs_ip_actions (account_id,character_guid,type,ip,systemnote,unixtime,time)
VALUES
    (?, 0, 1, ?, ?, unix_timestamp(NOW()), NOW())
