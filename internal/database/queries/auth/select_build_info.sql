SELECT
    majorVersion,
    minorVersion,
    bugfixVersion,
    hotfixVersion,
    build,
    winChecksumSeed,
    macChecksumSeed
FROM
    build_info
ORDER BY
    build ASC
