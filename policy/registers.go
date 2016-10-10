package policy

import (
    "time"

    "github.com/garyburd/redigo/redis"
    "github.com/earaujoassis/space/memstore"
)

func RegisterSignInAttempt(id string) {
    memstore.Start()
    defer memstore.Close()

    nowMoment := time.Now().UTC().Unix()
    if blockExists, _ := redis.Bool(memstore.Do("HEXISTS", "sign-in.blocked", id)); blockExists {
        blockReply, _ := redis.Int64(memstore.Do("HGET", "sign-in.blocked", id))
        if (nowMoment - blockReply) >= blockPeriodFailedSignIn  {
            memstore.Do("HDEL", "sign-in.blocked", id)
            memstore.Do("HSET", "sign-in.attempt", id, 1)
        }
        return
    }
    if exists, _ := redis.Bool(memstore.Do("HEXISTS", "sign-in.attempt", id)); !exists {
        memstore.Do("HSET", "sign-in.attempt", id, 1)
    } else {
        memstore.Do("HINCRBY", "sign-in.attempt", id, 1)
        reply, _ := redis.Int(memstore.Do("HGET", "sign-in.attempt", id))
        if reply >= attemptsUntilBlock {
            memstore.Do("HSET", "sign-in.blocked", id, nowMoment)
        }
    }
}

func RegisterSuccessfulSignIn(id string) {
    memstore.Start()
    defer memstore.Close()

    memstore.Do("HDEL", "sign-in.attempt", id)
    memstore.Do("HDEL", "sign-in.blocked", id)
}

func RegisterSignUpAttempt(id string) {
    memstore.Start()
    defer memstore.Close()

    nowMoment := time.Now().UTC().Unix()
    if blockExists, _ := redis.Bool(memstore.Do("HEXISTS", "sign-up.blocked", id)); blockExists {
        blockReply, _ := redis.Int64(memstore.Do("HGET", "sign-up.blocked", id))
        if (nowMoment - blockReply) >= blockPeriodFailedSignUp {
            memstore.Do("HDEL", "sign-up.blocked", id)
            memstore.Do("HSET", "sign-up.attempt", id, 1)
        }
        return
    }
    if exists, _ := redis.Bool(memstore.Do("HEXISTS", "sign-up.attempt", id)); !exists {
        memstore.Do("HSET", "sign-up.attempt", id, 1)
    } else {
        memstore.Do("HINCRBY", "sign-up.attempt", id, 1)
        reply, _ := redis.Int(memstore.Do("HGET", "sign-up.attempt", id))
        if reply >= attemptsUntilBlock {
            memstore.Do("HSET", "sign-up.blocked", id, nowMoment)
        }
    }
}

func RegisterSuccessfulSignUp(id string) {
    memstore.Start()
    defer memstore.Close()

    memstore.Do("HDEL", "sign-up.attempt", id)
    memstore.Do("HDEL", "sign-up.blocked", id)
}
