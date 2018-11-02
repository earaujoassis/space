package policy

import (
    "github.com/garyburd/redigo/redis"
    "github.com/earaujoassis/space/memstore"
)

// SignInAttemptStatus checks and controls sign-in attempts from a Web browser/User
func SignInAttemptStatus(id string) string {
    memstore.Start()
    defer memstore.Close()

    if blockExists, _ := redis.Bool(memstore.Do("HEXISTS", "sign-in.blocked", id)); blockExists {
        return Blocked
    }
    if attemptExists, _ := redis.Bool(memstore.Do("HEXISTS", "sign-in.attempt", id)); attemptExists {
        reply, _ := redis.Int(memstore.Do("HGET", "sign-in.attempt", id))
        switch {
        case reply > 0 && reply <= attemptsUntilPreblock:
            return Clear
        case reply > attemptsUntilPreblock && reply <= attemptsUntilBlock:
            return Preblocked
        case reply > attemptsUntilBlock:
            return Blocked
        }
    }
    return Clear
}

// SignUpAttemptStatus checks and controls sign-up attempts from a Web browser/User
func SignUpAttemptStatus(id string) string {
    memstore.Start()
    defer memstore.Close()

    if blockExists, _ := redis.Bool(memstore.Do("HEXISTS", "sign-up.blocked", id)); blockExists {
        return Blocked
    }
    if attemptExists, _ := redis.Bool(memstore.Do("HEXISTS", "sign-up.attempt", id)); attemptExists {
        reply, _ := redis.Int(memstore.Do("HGET", "sign-up.attempt", id))
        switch {
        case reply > 0 && reply <= attemptsUntilPreblock:
            return Clear
        case reply > attemptsUntilPreblock && reply <= attemptsUntilBlock:
            return Preblocked
        case reply > attemptsUntilBlock:
            return Blocked
        }
    }
    return Clear
}
