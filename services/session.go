package services

import (
    "time"

    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

const (
    eternalExpirationLength    int64 = 0
    largestExpirationLength    int64 = 3600 // 60 min
    defaultExpirationLength    int64 = 1800 // 30 min
    shortestExpirationLength   int64 = 300  //  5 min
)

func expirationLengthForTokenType(tokenType string) int64 {
    switch tokenType {
    case models.AccessToken:
        return largestExpirationLength
    case models.RefreshToken:
        return eternalExpirationLength
    case models.GrantToken:
        return shortestExpirationLength
    case models.ActionToken:
        return shortestExpirationLength
    default:
        return defaultExpirationLength
    }
}

func CreateSession(user models.User, client models.Client, ip, userAgent, scopes, tokenType string) models.Session {
    expirationLength := expirationLengthForTokenType(tokenType)
    var session models.Session = models.Session{
        User: user,
        Client: client,
        Ip: ip,
        UserAgent: userAgent,
        Scopes: scopes,
        TokenType: tokenType,
        ExpiresIn: expirationLength,
    }
    dataStore := datastore.GetDataStoreConnection()
    result := dataStore.Create(&session)
    if count := result.RowsAffected; count > 0 {
        return session
    }
    return models.Session{}
}

func SessionGrantsReadAbility(session models.Session) bool {
    return session.Scopes == models.ReadScope || session.Scopes == models.ReadWriteScope
}

func SessionGrantsWriteAbility(session models.Session) bool {
    return session.Scopes == models.ReadWriteScope
}

func withinExpirationWindow(session models.Session) bool {
    now := time.Now().UTC().Unix()
    if session.ExpiresIn == eternalExpirationLength || session.Moment + session.ExpiresIn >= now {
        return true
    }
    return false
}

func FindSessionByUUID(uuid string) models.Session {
    var session models.Session
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Preload("Client").
        Preload("User").
        Preload("User.Client").
        Preload("User.Language").
        Where("uuid = ? AND invalidated = false", uuid).
        First(&session)
    if session.ID != 0 {
        if !withinExpirationWindow(session) {
            InvalidateSession(session)
            return models.Session{}
        }
    }
    return session
}

func FindSessionByToken(token, tokenType string) models.Session {
    var session models.Session
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Preload("Client").
        Preload("User").
        Preload("User.Client").
        Preload("User.Language").
        Where("token = ? AND token_type = ? AND invalidated = false", token, tokenType).
        First(&session)
    if session.ID != 0 {
        if !withinExpirationWindow(session) {
            InvalidateSession(session)
            return models.Session{}
        }
    }
    return session
}

func InvalidateSession(session models.Session) {
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Model(&session).Select("invalidated").Update("invalidated", true)
}

func ActiveSessionsForClient(clientIID, userIID uint) int64 {
    var count struct{
        Count int64
    }

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Raw("SELECT count(*) AS count " +
            "FROM sessions WHERE token_type IN ('access_token', 'refresh_token') AND invalidated = false AND " +
            "client_id = ? AND user_id = ?;", clientIID, userIID).
        Scan(&count)
    return count.Count
}

func RevokeClientAccess(clientIID, userIID uint) {
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Exec("UPDATE sessions SET invalidated = true, updated_at = now() " +
            "WHERE token_type IN ('access_token', 'refresh_token') AND invalidated = false AND " +
            "client_id = ? AND user_id = ?;", clientIID, userIID)
}
