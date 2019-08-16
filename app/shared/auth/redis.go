package auth

// import (
// 	"sample/app/infrastructure"
// 	"strconv"
// 	"time"

// 	"github.com/garyburd/redigo/redis"
// )

// //HALFHOUR if token has lived more than this const it will be updated
// const HALFHOUR = 30

// //CheckAccessTokenFromRedis check and return token if a valid token existed on redis or call GCloud API to get a New token
// func CheckAccessTokenFromRedis(r redis.Conn) (string, error) {
// 	tokenKey := infrastructure.GetConfigString("gcp_jwt_redis_token_key")
// 	ttlKey := infrastructure.GetConfigString("gcp_jwt_redis_ttl_key")
// 	var (
// 		tm         time.Time
// 		willUpdate bool
// 		value      string
// 	)

// 	ttl, err := redis.String(r.Do("GET", ttlKey))
// 	if err == nil {
// 		tm, err = StringToTime(ttl)
// 		if err != nil {
// 			willUpdate = true
// 			_, delErr := r.Do("DEL", ttlKey)
// 			if delErr != nil {
// 				err = shareErrors.WithTrace(delErr, nil, nil)
// 			}
// 		}
// 		diffTime := time.Since(tm).Minutes()
// 		if diffTime > HALFHOUR {
// 			willUpdate = true
// 		}
// 	}

// 	//If no error getting token and TTL is valid, not null and not over HALFHOUR
// 	value, err = redis.String(r.Do("GET", tokenKey))
// 	if err == nil && !willUpdate && ttl != "" {
// 		return value, nil
// 	}

// 	jwt, err := GenerateGoogleJwtToken()
// 	if err != nil {
// 		return "", shareErrors.WithTrace(err, nil, nil)
// 	}
// 	at, err := GetGoogleOAuth2Token(jwt)
// 	if err != nil {
// 		return "", shareErrors.WithTrace(err, nil, nil)
// 	}
// 	if at.Error != "" {
// 		return "", shareErrors.WithTrace(shareErrors.New(at.Error), nil, nil)
// 	}

// 	_, err = r.Do("SET", ttlKey, strconv.FormatInt(time.Now().Unix(), 10))
// 	if err != nil {
// 		return at.AccessToken, shareErrors.WithTrace(err, nil, nil)
// 	}
// 	_, err = r.Do("SET", tokenKey, at.AccessToken)
// 	if err != nil {
// 		return at.AccessToken, shareErrors.WithTrace(err, nil, nil)
// 	}

// 	return at.AccessToken, nil
// }

// //StringToTime convert string to time.Time
// func StringToTime(timeStr string) (time.Time, error) {
// 	var tm time.Time
// 	i, err := strconv.ParseInt(timeStr, 10, 64)
// 	if err != nil {
// 		return tm, shareErrors.WithTrace(err, shareErrors.Fields{
// 			"timeStr": timeStr,
// 		}, nil)
// 	}
// 	tm = time.Unix(i, 0)
// 	return tm, err
// }
