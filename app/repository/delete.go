package repository

func (r *repo) DeleteRedis(key string) (int64, error) {
	result, err := r.redis.Del(key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}
