package account

type AccountType uint8

const (
	ACCOUNT_TYPE_SEC_PLAYER AccountType = iota
	ACCOUNT_TYPE_SEC_MODERATOR
	ACCOUNT_TYPE_SEC_GAMEMASTER
	ACCOUNT_TYPE_SEC_ADMINISTRATOR
	ACCOUNT_TYPE_SEC_CONSOLE
)
