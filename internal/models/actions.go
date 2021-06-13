package models

import (
    "time"
    "encoding/json"

    "github.com/garyburd/redigo/redis"

    "github.com/earaujoassis/space/internal/memstore"
)

const (
    // NotSpecialAction action description, used for ephemeral actions with no special meaning
    NotSpecialAction           string = "not_special"
    // UpdateUserAction action description, user for ephemeral actions updating user data
    UpdateUserAction           string = "update_user"
)

// Action is a model/struct used to represent ephemeral actions/sessions in the application
type Action struct {
    UUID string                 `validate:"omitempty,uuid4" json:"uuid"`
    User User                   `validate:"exists" json:"-"`
    UserID uint                 `json:"user_id"`
    Client Client               `validate:"exists" json:"-"`
    ClientID uint               `json:"client_id"`
    Moment int64                `json:"moment"`
    ExpiresIn int64             `json:"expires_in"`
    IP string                   `validate:"required" json:"ip"`
    UserAgent string            `validate:"required" json:"user_agent"`
    Token string                `validate:"omitempty,alphanum" json:"token"`
    Scopes string               `validate:"required,scope" json:"scopes"`
    Description string          `validate:"required,action" json:"description"`
    CreatedAt time.Time         `json:"created_at"`
}

func validAction(top interface{}, current interface{}, field interface{}, param string) bool {
    description := field.(string)
    if description != NotSpecialAction && description != UpdateUserAction {
        return false
    }
    return true
}

// Save saves an Action entry in a memory store (Redis)
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
    actionJSON, _ := json.Marshal(action)
    memstore.Do("HSET", "models.actions", action.UUID, actionJSON)
    memstore.Do("HSET", "models.actions.indexes", action.Token, action.UUID)
    memstore.Do("ZADD", "models.actions.rank", action.Moment, action.UUID)
    return nil
}

// Delete deletes an Action entry in a memory store (Redis)
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

// WithinExpirationWindow checks if an Action entry is still valid (time-based)
func (action *Action) WithinExpirationWindow() bool {
    now := time.Now().UTC().Unix()
    return action.ExpiresIn == eternalExpirationLength || action.Moment + action.ExpiresIn >= now
}

// CanUpdateUser checks if an Action description is valid for user update actions
func (action *Action) CanUpdateUser() bool {
    return action.Description == UpdateUserAction
}

// RetrieveActionByUUID obtains an Action entry from its UUID
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

// RetrieveActionByToken obtains an Action entry from its token-string
func RetrieveActionByToken(token string) Action {
    memstore.Start()
    defer memstore.Close()
    if indexExists, _ := redis.Bool(memstore.Do("HEXISTS", "models.actions.indexes", token)); !indexExists {
        return Action{}
    }
    actionUUID, _ := redis.String(memstore.Do("HGET", "models.actions.indexes", token))
    return RetrieveActionByUUID(actionUUID)
}
