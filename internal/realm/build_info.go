package realm

type BuildInfo struct {
	Build         uint32  `db:"build"`
	MajorVersion  uint32  `db:"majorVersion"`
	MinorVersion  uint32  `db:"minorVersion"`
	BugfixVersion uint32  `db:"bugfixVersion"`
	HotfixVersion []rune  `db:"hotfixVersion"`
	WindowsHash   []uint8 `db:"winChecksumSeed"`
	MacHash       []uint8 `db:"macChecksumSeed"`
}
