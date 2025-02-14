SELECT
    id,
    name,
    address,
    localAddress,
    localSubnetMask,
    port,
    icon,
    flag,
    timezone,
    allowedSecurityLevel,
    population,
    gamebuild
FROM
    realmlist
WHERE
    flag <> 3
ORDER BY name
