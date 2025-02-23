package packets

type AuthResult int
type AuthCommand int
type LoginResult int
type ExpansionFlag int

const (
	WOW_SUCCESS                              AuthResult = 0x00
	WOW_FAIL_BANNED                          AuthResult = 0x03
	WOW_FAIL_UNKNOWN_ACCOUNT                 AuthResult = 0x04
	WOW_FAIL_INCORRECT_PASSWORD              AuthResult = 0x05
	WOW_FAIL_ALREADY_ONLINE                  AuthResult = 0x06
	WOW_FAIL_NO_TIME                         AuthResult = 0x07
	WOW_FAIL_DB_BUSY                         AuthResult = 0x08
	WOW_FAIL_VERSION_INVALID                 AuthResult = 0x09
	WOW_FAIL_VERSION_UPDATE                  AuthResult = 0x0A
	WOW_FAIL_INVALID_SERVER                  AuthResult = 0x0B
	WOW_FAIL_SUSPENDED                       AuthResult = 0x0C
	WOW_FAIL_FAIL_NOACCESS                   AuthResult = 0x0D
	WOW_SUCCESS_SURVEY                       AuthResult = 0x0E
	WOW_FAIL_PARENTCONTROL                   AuthResult = 0x0F
	WOW_FAIL_LOCKED_ENFORCED                 AuthResult = 0x10
	WOW_FAIL_TRIAL_ENDED                     AuthResult = 0x11
	WOW_FAIL_USE_BATTLENET                   AuthResult = 0x12
	WOW_FAIL_ANTI_INDULGENCE                 AuthResult = 0x13
	WOW_FAIL_EXPIRED                         AuthResult = 0x14
	WOW_FAIL_NO_GAME_ACCOUNT                 AuthResult = 0x15
	WOW_FAIL_CHARGEBACK                      AuthResult = 0x16
	WOW_FAIL_INTERNET_GAME_ROOM_WITHOUT_BNET AuthResult = 0x17
	WOW_FAIL_GAME_ACCOUNT_LOCKED             AuthResult = 0x18
	WOW_FAIL_UNLOCKABLE_LOCK                 AuthResult = 0x19
	WOW_FAIL_CONVERSION_REQUIRED             AuthResult = 0x20
	WOW_FAIL_DISCONNECTED                    AuthResult = 0xFF
)

const (
	LOGIN_OK               LoginResult = 0x00
	LOGIN_FAILED           LoginResult = 0x01
	LOGIN_FAILED2          LoginResult = 0x02
	LOGIN_BANNED           LoginResult = 0x03
	LOGIN_UNKNOWN_ACCOUNT  LoginResult = 0x04
	LOGIN_UNKNOWN_ACCOUNT3 LoginResult = 0x05
	LOGIN_ALREADYONLINE    LoginResult = 0x06
	LOGIN_NOTIME           LoginResult = 0x07
	LOGIN_DBBUSY           LoginResult = 0x08
	LOGIN_BADVERSION       LoginResult = 0x09
	LOGIN_DOWNLOAD_FILE    LoginResult = 0x0A
	LOGIN_FAILED3          LoginResult = 0x0B
	LOGIN_SUSPENDED        LoginResult = 0x0C
	LOGIN_FAILED4          LoginResult = 0x0D
	LOGIN_CONNECTED        LoginResult = 0x0E
	LOGIN_PARENTALCONTROL  LoginResult = 0x0F
	LOGIN_LOCKED_ENFORCED  LoginResult = 0x10
)

const (
	POST_BC_EXP_FLAG  ExpansionFlag = 0x2
	PRE_BC_EXP_FLAG   ExpansionFlag = 0x1
	NO_VALID_EXP_FLAG ExpansionFlag = 0x0
)

const (
	AUTH_LOGON_CHALLENGE     AuthCommand = 0x00
	AUTH_LOGON_PROOF         AuthCommand = 0x01
	AUTH_RECONNECT_CHALLENGE AuthCommand = 0x02
	AUTH_RECONNECT_PROOF     AuthCommand = 0x03
	AUTH_REALM_LIST          AuthCommand = 0x10
	AUTH_XFER_INITIATE       AuthCommand = 0x30
	AUTH_XFER_DATA           AuthCommand = 0x31
	AUTH_XFER_ACCEPT         AuthCommand = 0x32
	AUTH_XFER_RESUME         AuthCommand = 0x33
	AUTH_XFER_CANCEL         AuthCommand = 0x34
)
