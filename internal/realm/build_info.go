package realm

import (
	"database/sql"

	"github.com/davidaburns/menethil-core/internal/database"
)

type BuildInfo struct {
	Build         uint32
	MajorVersion  uint32
	MinorVersion  uint32
	BugfixVersion uint32
	HotfixVersion []rune
	WindowsHash   []uint8
	MacHash       []uint8
}

func GetBuildInfoFromDb(build uint32, db *database.DatabaseClient) (*BuildInfo, error) {
	var buildInfo *BuildInfo
	parser := func(row *sql.Row) error {
		err := row.Scan(
			&buildInfo.MajorVersion,
			&buildInfo.MinorVersion,
			&buildInfo.BugfixVersion,
			&buildInfo.HotfixVersion,
			&buildInfo.Build,
			&buildInfo.WindowsHash,
			&buildInfo.MacHash,
		)

		if err != nil {
			return err
		}

		return nil
	}

	if err := db.QueryPreparedSingle(database.AUTH_SELECT_BUILD_INFO, []any{build}, parser); err != nil {
		return nil, err
	}

	return buildInfo, nil
}
