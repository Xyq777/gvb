package srv_redis

import (
	"context"
	"gvb/internal/global"
	"strconv"
	"time"
)

func SetToken(userID uint, jti string, exp time.Duration) error {
	Prefix := strconv.Itoa(int(userID)) + "_"
	err := global.Redis.Set(context.Background(), Prefix+jti, true, exp).Err()
	if err != nil {
		global.Log.Error(err)
		return err
	}
	return nil
}
func ExistToken(userID uint, jti string) (bool, error) {
	Prefix := strconv.Itoa(int(userID)) + "_"
	exists, err := global.Redis.Exists(context.Background(), Prefix+jti).Result()

	if err != nil {
		global.Log.Error(err)
		return false, err
	}
	return exists > 0, nil
}

func LogoutToken(userID uint, jti string) error {
	Prefix := strconv.Itoa(int(userID)) + "_"
	err := global.Redis.Del(context.Background(), Prefix+jti).Err()
	if err != nil {
		global.Log.Error(err)
		return err
	}
	return nil
}
