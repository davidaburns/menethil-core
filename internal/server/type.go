package server

type ServerType int

const (
	ServerAuth ServerType = iota
	ServerWorld
)

func (st *ServerType) String() string {
	switch *st {
	case ServerAuth:
		return "Authentication"
	case ServerWorld:
		return "World"
	default:
		return "Unkown"
	}
}

func (st *ServerType) StringShort() string {
	switch *st {
	case ServerAuth:
		return "auth"
	case ServerWorld:
		return "world"
	default:
		return "unknown"
	}
}
