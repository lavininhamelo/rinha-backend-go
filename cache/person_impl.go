package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"rinha-basic/entities"
)

type CacheStorage struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewCacheStorage(rdb *redis.Client) CacheStorage {
	return CacheStorage{
		Client: rdb,
		Ctx:    context.Background(),
	}
}

func (rs *CacheStorage) SaveUsername(username string) (string, error) {
	err := rs.Client.Set(rs.Ctx, username, username, 0).Err()
	if err != nil {
		return "", err
	}
	return username, nil
}

func (rs *CacheStorage) VerifyUsername(username string) (bool, error) {
	val, err := rs.Client.Get(rs.Ctx, username).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return val == username, nil
}

func (rs *CacheStorage) SavePerson(id string, person *entities.Person) (bool, error) {
	personJSON, err := json.Marshal(person)
	if err != nil {
		return false, err
	}
	err = rs.Client.Set(rs.Ctx, id, personJSON, 30*time.Second).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (rs *CacheStorage) RecoveryPerson(id string) (*entities.Person, error) {
	val, err := rs.Client.Get(rs.Ctx, id).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	person := &entities.Person{}
	err = json.Unmarshal([]byte(val), person)
	if err != nil {
		return nil, err
	}
	return person, nil
}
