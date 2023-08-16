package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/freecode23/go-react-chatapp/pkg/message"
	"github.com/freecode23/go-react-chatapp/pkg/storage"
	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	ctx         context.Context
	redisClient *redis.Client
}

func NewRedisStore() *RedisStore {
	return &RedisStore{
		ctx: context.Background(),
		redisClient: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}
}

/*
*
Save message to redis store
method is always implemented on the concrete type
*
*/
func (rs *RedisStore) SaveMessageToStore(message message.Message) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Use LPUSH to add the message to the front of the list in Redis
	// with the key chatHistory
	// if a list with key chatHistory does not exist yet, it will create it
	err = rs.redisClient.LPush(rs.ctx, "chatHistory", messageJSON).Err()
	if err != nil {
		return err
	}

	// Limit the history size to a certain number of messages
	rs.redisClient.LTrim(rs.ctx, "chatHistory", 0, 9) // Keep the latest 10 messages

	return nil
}

func (rs *RedisStore) GetLast30Messages() ([]message.Message, error) {
	// Fetch the last 10 messages using LRange.
	// Note: As we are storing the latest messages at the start of the list using LPUSH,
	// the last 10 messages would be the first 10 messages in the list.
	redisMessages, err := rs.redisClient.LRange(rs.ctx, "chatHistory", 0, 9).Result()
	if err != nil {
		return nil, err
	}

	// 2. init messages slice
	messages := make([]message.Message, 0, len(redisMessages))

	// 3. for each mesg in redis result
	for _, redisString := range redisMessages {

		// create socketMsg type
		var msg message.Message

		// store the redisString json to msg of struct message
		err := json.Unmarshal([]byte(redisString), &msg)

		if err != nil {
			return nil, err
		}

		//
		messages = append(messages, msg)
	}

	// 4. return list of message struct

	return messages, nil
}

func (rs *RedisStore) UploadMessagesToS3() {
	const threshold = 10

	// 0. get current redis chatHistory length
	length, err := rs.redisClient.LLen(rs.ctx, "chatHistory").Result()
	if err != nil {
		panic(err)
	}

	// 1. if its already full, save to s3
	if length >= threshold {

		// Retrieve 10 messages
		chatHistory10 := rs.redisClient.LRange(rs.ctx, "chatHistory", 0, threshold-1)

		// Push to S3
		storage.SaveChatHistory(chatHistory10)

		// If the upload is successful, remove all messages from Redis
		rs.redisClient.Del(rs.ctx, "chatHistory")

		length, err := rs.redisClient.LLen(rs.ctx, "chatHistory").Result()
		fmt.Println("redis:after delete length redis=", length)
		if err != nil {
			panic(err)
		}
	}
}
