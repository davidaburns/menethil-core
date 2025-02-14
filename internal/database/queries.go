package database

import "embed"

//go:embed queries/*
var embeddedQueries embed.FS

type QueryStatementPath string

// AUTH QUERIES
const (
	AUTH_DELETE_ACCOUNT_ACCESS_BY_REALM          QueryStatementPath = "queries/auth/delete_account_access_by_realm.sql"
	AUTH_DELETE_ACCOUNT_ACCESS                   QueryStatementPath = "queries/auth/delete_account_access.sql"
	AUTH_DELETE_ACCOUNT_BANNED                   QueryStatementPath = "queries/auth/delete_account_banned.sql"
	AUTH_DELETE_ACCOUNT_MUTED                    QueryStatementPath = "queries/auth/delete_account_muted.sql"
	AUTH_DELETE_ACCOUNT                          QueryStatementPath = "queries/auth/delete_account.sql"
	AUTH_DELETE_EXPIRED_IP_BANS                  QueryStatementPath = "queries/auth/delete_expired_ip_bans.sql"
	AUTH_DELETE_OLD_LOGS                         QueryStatementPath = "queries/auth/delete_old_logs.sql"
	AUTH_DELETE_REALM_CHARACTERS                 QueryStatementPath = "queries/auth/delete_realm_characters.sql"
	AUTH_DELETE_SECRET_DIGEST                    QueryStatementPath = "queries/auth/delete_secret_digest.sql"
	AUTH_INSERT_ACCOUNT_ACCESS                   QueryStatementPath = "queries/auth/insert_account_access.sql"
	AUTH_INSERT_ACCOUNT_AUTO_BANNED              QueryStatementPath = "queries/auth/insert_account_auto_banned.sql"
	AUTH_INSERT_ACCOUNT_BANNED                   QueryStatementPath = "queries/auth/insert_account_banned.sql"
	AUTH_INSERT_ACCOUNT_MUTE_INFO                QueryStatementPath = "queries/auth/insert_account_mute_info.sql"
	AUTH_INSERT_ACCOUNT_MUTE                     QueryStatementPath = "queries/auth/insert_account_mute.sql"
	AUTH_INSERT_ACCOUNT                          QueryStatementPath = "queries/auth/insert_account.sql"
	AUTH_INSERT_IP_AUTO_BANNED                   QueryStatementPath = "queries/auth/insert_ip_auto_banned.sql"
	AUTH_INSERT_IP_BANNED                        QueryStatementPath = "queries/auth/insert_ip_banned.sql"
	AUTH_INSERT_IP_LOGGING_CHARACTER_DELETE      QueryStatementPath = "queries/auth/insert_ip_logging_character_delete.sql"
	AUTH_INSERT_IP_LOGGING_FAILED_LOGIN          QueryStatementPath = "queries/auth/insert_ip_logging_failed_login.sql"
	AUTH_INSERT_IP_LOGGING_FAILED_PASSWORD_LOGIN QueryStatementPath = "queries/auth/insert_ip_logging_failed_password_login.sql"
	AUTH_INSERT_IP_LOGGING_LOGIN_DELETE          QueryStatementPath = "queries/auth/insert_ip_logging_login_delete.sql"
	AUTH_INSERT_LOG                              QueryStatementPath = "queries/auth/insert_log.sql"
	AUTH_INSERT_REALM_CHARACTERS_INIT            QueryStatementPath = "queries/auth/insert_realm_characters_init.sql"
	AUTH_INSERT_SECRET_DIGEST                    QueryStatementPath = "queries/auth/insert_secret_digest.sql"
	AUTH_INSERT_UPTIME                           QueryStatementPath = "queries/auth/insert_uptime.sql"
	AUTH_REPLACE_MOTD_LOCALE                     QueryStatementPath = "queries/auth/replace_motd_locale.sql"
	AUTH_REPLACE_MOTD                            QueryStatementPath = "queries/auth/replace_motd.sql"
	AUTH_REPLACE_REALM_CHARACTERS                QueryStatementPath = "queries/auth/replace_realm_characters.sql"
	AUTH_SELECT_ACCOUNT_ACCESS_GMLEVEL_TEST      QueryStatementPath = "queries/auth/select_account_access_gmlevel_test.sql"
	AUTH_SELECT_ACCOUNT_ACCESS_GMLEVEL           QueryStatementPath = "queries/auth/select_account_access_gmlevel.sql"
	AUTH_SELECT_ACCOUNT_ACCESS                   QueryStatementPath = "queries/auth/select_account_access.sql"
	AUTH_SELECT_ACCOUNT_BANNED_ALL               QueryStatementPath = "queries/auth/select_account_banned_all.sql"
	AUTH_SELECT_ACCOUNT_BANNED_BY_USERNAME       QueryStatementPath = "queries/auth/select_account_banned_by_username.sql"
	AUTH_SELECT_ACCOUNT_BANNED                   QueryStatementPath = "queries/auth/select_account_banned.sql"
	AUTH_SELECT_ACCOUNT_BANS                     QueryStatementPath = "queries/auth/select_account_bans.sql"
	AUTH_SELECT_ACCOUNT_BY_ID                    QueryStatementPath = "queries/auth/select_account_by_id.sql"
	AUTH_SELECT_ACCOUNT_BY_IP                    QueryStatementPath = "queries/auth/select_account_by_ip.sql"
	AUTH_SELECT_ACCOUNT_CHECK_PASSWORD_BY_NAME   QueryStatementPath = "queries/auth/select_account_check_password_by_name.sql"
	AUTH_SELECT_ACCOUNT_CHECK_PASSWORD           QueryStatementPath = "queries/auth/select_account_check_password.sql"
	AUTH_SELECT_ACCOUNT_GMS                      QueryStatementPath = "queries/auth/select_account_gms.sql"
	AUTH_SELECT_ACCOUNT_ID_BY_NAME               QueryStatementPath = "queries/auth/select_account_id_by_name.sql"
	AUTH_SELECT_ACCOUNT_ID_BY_USERNAME           QueryStatementPath = "queries/auth/select_account_id_by_username.sql"
	AUTH_SELECT_ACCOUNT_INFO_BY_NAME             QueryStatementPath = "queries/auth/select_account_info_by_name.sql"
	AUTH_SELECT_ACCOUNT_INFO                     QueryStatementPath = "queries/auth/select_account_info.sql"
	AUTH_SELECT_ACCOUNT_LAST_ATTEMPT_IP          QueryStatementPath = "queries/auth/select_account_last_attempt_ip.sql"
	AUTH_SELECT_ACCOUNT_LAST_IP                  QueryStatementPath = "queries/auth/select_account_last_ip.sql"
	AUTH_SELECT_ACCOUNT_LIST_BY_EMAIL            QueryStatementPath = "queries/auth/select_account_list_by_email.sql"
	AUTH_SELECT_ACCOUNT_LIST_BY_NAME             QueryStatementPath = "queries/auth/select_account_list_by_name.sql"
	AUTH_SELECT_ACCOUNT_PINFO_BANS               QueryStatementPath = "queries/auth/select_account_pinfo_bans.sql"
	AUTH_SELECT_ACCOUNT_PINFO                    QueryStatementPath = "queries/auth/select_account_pinfo.sql"
	AUTH_SELECT_ACCOUNT_RECRUITER                QueryStatementPath = "queries/auth/select_account_recruiter.sql"
	AUTH_SELECT_ACCOUNT_TOTP_SECRET              QueryStatementPath = "queries/auth/select_account_totp_secret.sql"
	AUTH_SELECT_ACCOUNT_USERNAME_BY_ID           QueryStatementPath = "queries/auth/select_account_username_by_id.sql"
	AUTH_SELECT_ACCOUNT_WHOIS                    QueryStatementPath = "queries/auth/select_account_whois.sql"
	AUTH_SELECT_AUTOBROADCAST_LOCALIZED          QueryStatementPath = "queries/auth/select_autobroadcast_localized.sql"
	AUTH_SELECT_AUTOBROADCAST                    QueryStatementPath = "queries/auth/select_autobroadcast.sql"
	AUTH_SELECT_FAILED_LOGINS                    QueryStatementPath = "queries/auth/select_failed_logins.sql"
	AUTH_SELECT_GMLEVEL_BY_REALMID               QueryStatementPath = "queries/auth/select_gmlevel_by_realmid.sql"
	AUTH_SELECT_IP_BANNED_ALL                    QueryStatementPath = "queries/auth/select_ip_banned_all.sql"
	AUTH_SELECT_IP_BANNED_BY_IP                  QueryStatementPath = "queries/auth/select_ip_banned_by_ip.sql"
	AUTH_SELECT_IP_BANNED                        QueryStatementPath = "queries/auth/select_ip_banned.sql"
	AUTH_SELECT_IP_INFO                          QueryStatementPath = "queries/auth/select_ip_info.sql"
	AUTH_SELECT_LOGIN_CHALLENGE                  QueryStatementPath = "queries/auth/select_login_challenge.sql"
	AUTH_SELECT_MOTD_LOCALE                      QueryStatementPath = "queries/auth/select_motd_locale.sql"
	AUTH_SELECT_MOTD                             QueryStatementPath = "queries/auth/select_motd.sql"
	AUTH_SELECT_NUMCHARS_ON_REALM                QueryStatementPath = "queries/auth/select_numchars_on_realm.sql"
	AUTH_SELECT_REALM_CHARACTERS_COUNT           QueryStatementPath = "queries/auth/select_realm_characters_count.sql"
	AUTH_SELECT_REALMLIST_SECURITY_LEVEL         QueryStatementPath = "queries/auth/select_realmlist_security_level.sql"
	AUTH_SELECT_REALMLIST                        QueryStatementPath = "queries/auth/select_realmlist.sql"
	AUTH_SELECT_RECONNECT_CHALLENGE              QueryStatementPath = "queries/auth/select_reconnect_challenge.sql"
	AUTH_SELECT_SECRET_DIGEST                    QueryStatementPath = "queries/auth/select_secret_digest.sql"
	AUTH_SELECT_SUM_REALM_CHARACTERS             QueryStatementPath = "queries/auth/select_sum_realm_characters.sql"
	AUTH_UPDATE_ACCOUNT_LOCK_COUNTRY             QueryStatementPath = "queries/auth/update_account_lock_country.sql"
	AUTH_UPDATE_ACCOUNT_LOCK                     QueryStatementPath = "queries/auth/update_account_lock.sql"
	AUTH_UPDATE_ACCOUNT_NOT_BANNED               QueryStatementPath = "queries/auth/update_account_not_banned.sql"
	AUTH_UPDATE_ACCOUNT_ONLINE                   QueryStatementPath = "queries/auth/update_account_online.sql"
	AUTH_UPDATE_ACCOUNT_TOTP_SECRET              QueryStatementPath = "queries/auth/update_account_totp_secret.sql"
	AUTH_UPDATE_EMAIL                            QueryStatementPath = "queries/auth/update_email.sql"
	AUTH_UPDATE_EXPANSION                        QueryStatementPath = "queries/auth/update_expansion.sql"
	AUTH_UPDATE_EXPIRED_ACCOUNT_BANS             QueryStatementPath = "queries/auth/update_expired_account_bans.sql"
	AUTH_UPDATE_FAILED_LOGINS                    QueryStatementPath = "queries/auth/update_failed_logins.sql"
	AUTH_UPDATE_LAST_ATTEMPT_IP                  QueryStatementPath = "queries/auth/update_last_attempt_ip.sql"
	AUTH_UPDATE_LAST_IP                          QueryStatementPath = "queries/auth/update_last_ip.sql"
	AUTH_UPDATE_LOGON                            QueryStatementPath = "queries/auth/update_logon.sql"
	AUTH_UPDATE_LOGON_PROOF                      QueryStatementPath = "queries/auth/update_logonproof.sql"
	AUTH_UPDATE_MUTE_TIME_LOGIN                  QueryStatementPath = "queries/auth/update_mute_time_login.sql"
	AUTH_UPDATE_MUTE_TIME                        QueryStatementPath = "queries/auth/update_mute_time.sql"
	AUTH_UPDATE_UPTIME_PLAYERS                   QueryStatementPath = "queries/auth/update_uptime_players.sql"
	AUTH_UPDATE_USERNAME                         QueryStatementPath = "queries/auth/delete_account_access_by_realm.sql"
)

var AuthQueries []QueryStatementPath = []QueryStatementPath {
	AUTH_DELETE_ACCOUNT_ACCESS_BY_REALM,
	AUTH_DELETE_ACCOUNT_ACCESS,
	AUTH_DELETE_ACCOUNT_BANNED,
	AUTH_DELETE_ACCOUNT_MUTED,
	AUTH_DELETE_ACCOUNT,
	AUTH_DELETE_EXPIRED_IP_BANS,
	AUTH_DELETE_OLD_LOGS,
	AUTH_DELETE_REALM_CHARACTERS,
	AUTH_DELETE_SECRET_DIGEST,
	AUTH_INSERT_ACCOUNT_ACCESS,
	AUTH_INSERT_ACCOUNT_AUTO_BANNED,
	AUTH_INSERT_ACCOUNT_BANNED,
	AUTH_INSERT_ACCOUNT_MUTE_INFO,
	AUTH_INSERT_ACCOUNT_MUTE,
	AUTH_INSERT_ACCOUNT,
	AUTH_INSERT_IP_AUTO_BANNED,
	AUTH_INSERT_IP_BANNED,
	AUTH_INSERT_IP_LOGGING_CHARACTER_DELETE,
	AUTH_INSERT_IP_LOGGING_FAILED_LOGIN,
	AUTH_INSERT_IP_LOGGING_FAILED_PASSWORD_LOGIN,
	AUTH_INSERT_IP_LOGGING_LOGIN_DELETE,
	AUTH_INSERT_LOG,
	AUTH_INSERT_REALM_CHARACTERS_INIT,
	AUTH_INSERT_SECRET_DIGEST,
	AUTH_INSERT_UPTIME,
	AUTH_REPLACE_MOTD_LOCALE,
	AUTH_REPLACE_MOTD,
	AUTH_REPLACE_REALM_CHARACTERS,
	AUTH_SELECT_ACCOUNT_ACCESS_GMLEVEL_TEST,
	AUTH_SELECT_ACCOUNT_ACCESS_GMLEVEL,
	AUTH_SELECT_ACCOUNT_ACCESS,
	AUTH_SELECT_ACCOUNT_BANNED_ALL,
	AUTH_SELECT_ACCOUNT_BANNED_BY_USERNAME,
	AUTH_SELECT_ACCOUNT_BANNED,
	AUTH_SELECT_ACCOUNT_BANS,
	AUTH_SELECT_ACCOUNT_BY_ID,
	AUTH_SELECT_ACCOUNT_BY_IP,
	AUTH_SELECT_ACCOUNT_CHECK_PASSWORD_BY_NAME,
	AUTH_SELECT_ACCOUNT_CHECK_PASSWORD,
	AUTH_SELECT_ACCOUNT_GMS,
	AUTH_SELECT_ACCOUNT_ID_BY_NAME,
	AUTH_SELECT_ACCOUNT_ID_BY_USERNAME,
	AUTH_SELECT_ACCOUNT_INFO_BY_NAME,
	AUTH_SELECT_ACCOUNT_INFO,
	AUTH_SELECT_ACCOUNT_LAST_ATTEMPT_IP,
	AUTH_SELECT_ACCOUNT_LAST_IP,
	AUTH_SELECT_ACCOUNT_LIST_BY_EMAIL,
	AUTH_SELECT_ACCOUNT_LIST_BY_NAME,
	AUTH_SELECT_ACCOUNT_PINFO_BANS,
	AUTH_SELECT_ACCOUNT_PINFO,
	AUTH_SELECT_ACCOUNT_RECRUITER,
	AUTH_SELECT_ACCOUNT_TOTP_SECRET,
	AUTH_SELECT_ACCOUNT_USERNAME_BY_ID,
	AUTH_SELECT_ACCOUNT_WHOIS,
	AUTH_SELECT_AUTOBROADCAST_LOCALIZED,
	AUTH_SELECT_AUTOBROADCAST,
	AUTH_SELECT_FAILED_LOGINS,
	AUTH_SELECT_GMLEVEL_BY_REALMID,
	AUTH_SELECT_IP_BANNED_ALL,
	AUTH_SELECT_IP_BANNED_BY_IP,
	AUTH_SELECT_IP_BANNED,
	AUTH_SELECT_IP_INFO,
	AUTH_SELECT_LOGIN_CHALLENGE,
	AUTH_SELECT_MOTD_LOCALE,
	AUTH_SELECT_MOTD,
	AUTH_SELECT_NUMCHARS_ON_REALM,
	AUTH_SELECT_REALM_CHARACTERS_COUNT,
	AUTH_SELECT_REALMLIST_SECURITY_LEVEL,
	AUTH_SELECT_REALMLIST,
	AUTH_SELECT_RECONNECT_CHALLENGE,
	AUTH_SELECT_SECRET_DIGEST,
	AUTH_SELECT_SUM_REALM_CHARACTERS,
	AUTH_UPDATE_ACCOUNT_LOCK_COUNTRY,
	AUTH_UPDATE_ACCOUNT_LOCK,
	AUTH_UPDATE_ACCOUNT_NOT_BANNED,
	AUTH_UPDATE_ACCOUNT_ONLINE,
	AUTH_UPDATE_ACCOUNT_TOTP_SECRET,
	AUTH_UPDATE_EMAIL,
	AUTH_UPDATE_EXPANSION,
	AUTH_UPDATE_EXPIRED_ACCOUNT_BANS,
	AUTH_UPDATE_FAILED_LOGINS,
	AUTH_UPDATE_LAST_ATTEMPT_IP,
	AUTH_UPDATE_LAST_IP,
	AUTH_UPDATE_LOGON,
	AUTH_UPDATE_LOGON_PROOF,
	AUTH_UPDATE_MUTE_TIME_LOGIN,
	AUTH_UPDATE_MUTE_TIME,
	AUTH_UPDATE_UPTIME_PLAYERS,
	AUTH_UPDATE_USERNAME,
}
