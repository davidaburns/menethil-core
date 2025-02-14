SELECT id, locale, text
FROM autobroadcast_locale
WHERE realmid = ? OR realmid = -1
