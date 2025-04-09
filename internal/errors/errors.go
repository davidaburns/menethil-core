package errors

import err "errors"

var (
	// Server Errors
	ErrServerTypeUnkown = err.New("unkown server type")

	// Config Errors
	ErrConfigSectionNotFound = err.New("config section not found")
	ErrConfigKeyNotFound     = err.New("config key not found")

	// Crypto Errors
	ErrCryptoNotInitialized = err.New("crypto has not been initialized")

	// Header Errors
	ErrHeaderSizeTooLarge = err.New("header size is too large")
)
