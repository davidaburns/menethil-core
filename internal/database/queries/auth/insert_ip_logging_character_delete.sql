INSERT INTO logs_ip_actions (account_id,character_guid,type,ip,systemnote,unixtime,time)
VALUES
    (?, ?, ?, ?, ?, unix_timestamp(NOW()), NOW())
