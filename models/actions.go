package models

import (
    "time"
    "encoding/json"

    "github.com/garyburd/redigo/redis"
    "github.com/earaujoassis/space/memstore"
)

type Action struct {
    UUID string                 `validate:"omitempty,uuid4" json:"uuid"`
    User User                   `validate:"exists" json:"-"`
    UserID uint                 `json:"user_id"`
    Client Client               `validate:"exists" json:"-"`
    ClientID uint               `json:"client_id"`
    Moment int64                `json:"moment"`
    ExpiresIn int64             `json:"expires_in"`
    Ip string                   `validate:"required" json:"ip"`
    UserAgent string            `validate:"required" json:"user_agent"`
    Token string                `validate:"omitempty,alphanum" json:"token"`
    Scopes string               `validate:"required,scope" json:"scopes"`
    CreatedAt time.Time         `json:"created_at"`
}

func (action *Action) Save() error {
    action.UserID = action.User.ID
    action.ClientID = action.Client.ID
    action.UUID = generateUUID()
    action.CreatedAt = time.Now().UTC()
    action.Token = GenerateRandomString(64)
    action.Moment = time.Now().UTC().Unix()
    action.ExpiresIn = shortestExpirationLength
    if err := validateModel("validate", action); err != nil {
        return err
    }
    memstore.Start()
    defer memstore.Close()
    actionJson, _ := json.Marshal(action)
    memstore.Do("HSET", "models.actions", action.UUID, actionJson)
    memstore.Do("HSET", "models.actions.indexes", action.Token, action.UUID)
    memstore.Do("ZADD", "models.actions.rank", action.Moment, action.UUID)
    return nil
}

func (action *Action) Delete() {
    memstore.Start()
    defer memstore.Close()
    if actionExists, _ := redis.Bool(memstore.Do("HEXISTS", "models.actions", action.UUID)); !actionExists {
        return
    }
    memstore.Do("HDEL", "models.actions.indexes", action.Token)
    memstore.Do("HDEL", "models.actions", action.UUID)
    memstore.Do("ZREM", "models.actions.rank", action.UUID)
}

func (action *Action) WithinExpirationWindow() bool {
    now := time.Now().UTC().Unix()
    if action.ExpiresIn == eternalExpirationLength || action.Moment + action.ExpiresIn >= now {
        return true
    }
    return false
}

func RetrieveActionByUUID(uuid string) Action {
    var action Action
    memstore.Start()
    defer memstore.Close()
    if actionExists, _ := redis.Bool(memstore.Do("HEXISTS", "models.actions", uuid)); !actionExists {
        return Action{}
    }
    actionString, _ := redis.String(memstore.Do("HGET", "models.actions", uuid))
    if err := json.Unmarshal([]byte(actionString), &action); err != nil {
        return Action{}
    }
    return action
}

func RetrieveActionByToken(token string) Action {
    memstore.Start()
    defer memstore.Close()
    if indexExists, _ := redis.Bool(memstore.Do("HEXISTS", "models.actions.indexes", token)); !indexExists {
        return Action{}
    }
    actionUUID, _ := redis.String(memstore.Do("HGET", "models.actions.indexes", token))
    return RetrieveActionByUUID(actionUUID)
}
