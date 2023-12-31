package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/freecode23/go-react-chatapp/pkg/message"
	"github.com/freecode23/go-react-chatapp/pkg/storage"
	redis "github.com/go-redis/redis/v8"
	rejson "github.com/nitishm/go-rejson/v4"
)

type RedisStore struct {
	ctx              context.Context
	redisClient      *redis.Client
	rejsonHandler    *rejson.Handler
	rediSearchClient *redisearch.Client
}

func NewRedisStore() *RedisStore {
	// 1. init redisClient
	rdclient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 2. init json handler for redisClient
	rejsonHandler := rejson.NewReJSONHandler()
	rejsonHandler.SetGoRedisClient(rdclient)

	rediSearchClient := createChatHistoryIndex()

	return &RedisStore{
		ctx:              context.Background(),
		redisClient:      rdclient,
		rediSearchClient: rediSearchClient,
		rejsonHandler:    rejsonHandler,
	}
}

/*
*
1. redisJson add
Save message to redis store as JSON
method is always implemented on the concrete type
*
*/
func (rs *RedisStore) SaveMessageToStore(msgStruct message.Message) error {
	// 1. conver struct to json
	msgJSONbytes, err := json.Marshal(msgStruct)
	if err != nil {
		return err
	}

	// 2. create key for the json object
	key := msgStruct.ID

	// 3. Insert the message into Redis using RedisJSON
	_, err = rs.rejsonHandler.JSONSet(key, ".", msgJSONbytes)
	if err != nil {
		fmt.Println("redis: cannot insert key and message to json")
		return err
	}

	// 4. create a doc of this msgStruct for redisearch
	timestampEpoch := msgStruct.Timestamp.Unix()
	docMessage := redisearch.NewDocument("id"+key, 1).
		Set("room", msgStruct.RoomName).
		Set("user", msgStruct.UserName).
		Set("body", msgStruct.Body).
		Set("timestamp", timestampEpoch)

	// 5. Index the document:
	// meaning it will be stored in such a way that you can quickly search for it based on its content.
	// Index the documents. The API accepts multiple documents at a time
	if err := rs.rediSearchClient.IndexOptions(redisearch.DefaultIndexingOptions, docMessage); err != nil {
		fmt.Println("redis: cannot index this documnet", docMessage)
		log.Fatal(err)
	}

	return nil
}

/*
*
2. redisJson get
*
*/
func (rs *RedisStore) GetLastMessagesStruct(roomName string) ([]message.Message, error) {

	// 1. Create a query to retrieve all messages based on the room name.
	q := redisearch.NewQuery(fmt.Sprintf("@room:'%s'", roomName)).
		SetSortBy("timestamp", false) // Order by timestamp descending.

	// 2. Execute the query
	docs, _, _ := rs.rediSearchClient.Search(q)

	// 3. Convert the search results to message.Message objects.
	messagesOut := make([]message.Message, 0, len(docs))
	for _, doc := range docs {

		// - convert to time format
		timestamp := convertStringSecToTimeStamp(doc.Properties["timestamp"].(string))

		// - init message struct
		msg := message.Message{
			ID:        doc.Id[strings.Index(doc.Id, "d")+1:], // get the numeric part of the id
			RoomName:  doc.Properties["room"].(string),
			UserName:  doc.Properties["user"].(string),
			Body:      doc.Properties["body"].(string),
			Timestamp: timestamp,
		}

		// - add to list
		messagesOut = append(messagesOut, msg)
	}

	return messagesOut, nil
}

/*
*
3. redisJson delete
*
*/
func (rs *RedisStore) UploadMessagesToS3(roomName string) {
	const threshold = 10

	// 0. get current redis chatHistory length
	length, err := rs.countMessagesInRoom(roomName)
	if err != nil {
		log.Printf("redis: Failed to get chatHistory length: %v", err)
	}

	// 1. if its already full, save to s3
	if length >= threshold {

		// 1. Retrieve 10 messages - using
		chatHistory10Struct, _ := rs.GetLastMessagesStruct(roomName)

		// 2. Convert chatHistory to string
		var chatHistory10Str []string
		for _, msgStruct := range chatHistory10Struct {
			jsonByte, err := json.Marshal(msgStruct)
			if err != nil {
				log.Printf("redis: failed to marshal chatHistory10Str to JSON")
			}
			chatHistory10Str = append(chatHistory10Str, string(jsonByte))
		}

		if err != nil {
			log.Printf("redis:Failed to get 10 chatHistory: %v", err)
		}
		// 3. Push to S3
		storage.SaveChatHistory(chatHistory10Str)

		// 4. If the upload is successful, remove all messages from Redis
		err := rs.DeleteMessagesFromRoom(roomName)
		if err != nil {
			log.Printf("redis: Failed to delete chatHistory: %v", err)
		}
	}
}

/*
*
Helper 0. create redis index of chat history so we can query fast
*
*/
func createChatHistoryIndex() *redisearch.Client {
	// 1. Check if "chatHistoryIndex" already exists
	rediSearchClient := redisearch.NewClient("localhost:6379", "chatHistoryIndex")

	_, err := rediSearchClient.Info()
	// 2. Index already exists. So, we return the client.
	if err == nil {
		return rediSearchClient

		// 3. If there's an error other than "Unknown Index name", it's unexpected. Handle appropriately.
	} else if !strings.Contains(err.Error(), "Unknown index name") {
		log.Fatal("redis: error checking index:", err)
	}

	// 4. If we reached here, the index doesn't exist. So, we proceed to create it.
	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("room")).
		AddField(redisearch.NewTextField("user")).
		AddField(redisearch.NewTextFieldOptions("body", redisearch.TextFieldOptions{Weight: 5.0})).
		AddField(redisearch.NewNumericFieldOptions("timestamp", redisearch.NumericFieldOptions{Sortable: true}))

	// 5. Create the index with the given schema
	if err := rediSearchClient.CreateIndex(sc); err != nil {
		log.Fatal("redis: createIndex error:", err)
	}

	// 6. Return the RediSearch client
	return rediSearchClient
}

/*
*
Helper 1. count total number of messages in a chat room
*
*/
func (rs *RedisStore) countMessagesInRoom(roomName string) (int, error) {

	agg := redisearch.NewAggregateQuery().
		Load([]string{"room"}). // Note: Do not use "@" here, as it's used for referencing, not loading
		Filter(fmt.Sprintf("@room=='%s'", roomName))

	_, total, err := rs.rediSearchClient.Aggregate(agg)
	if err != nil {
		fmt.Println("redis: error cannot get total:", err)
		return 0, err
	}

	return total, nil
}

/*
*
Helper 2: Convert string in second to unix time stamp
*
*/
func convertStringSecToTimeStamp(secondsStr string) time.Time {
	secondsInt, _ := strconv.ParseInt(secondsStr, 10, 64)

	timestamp := time.Unix(secondsInt, 0)
	return timestamp
}

/*
*
Helper 3: delete all messages history from a chatroom
*
*/

func (rs *RedisStore) DeleteMessagesFromRoom(roomName string) error {

	// 1. Create a query to fetch all message IDs associated with the room
	q := redisearch.NewQuery(fmt.Sprintf("@room:'%s'", roomName))

	// 2. Execute the query
	docs, _, err := rs.rediSearchClient.Search(q)
	if err != nil {
		return fmt.Errorf("failed to retrieve messages for room: %v", err)
	}

	// 3. Loop through the retrieved docs and delete each one
	for _, doc := range docs {
		// Delete JSON data from Redis
		// fmt.Printf("\nredis:  DeleteMessagesFromRoom %v:\n", doc)
		_, err := rs.rejsonHandler.JSONDel(doc.Id[2:], ".")
		if err != nil {

			return fmt.Errorf("failed to delete message JSON for id %s: %v", doc.Id, err)
		}
		// Delete the document from RediSearch index
		err = rs.rediSearchClient.DeleteDocument(doc.Id)
		if err != nil {

			return fmt.Errorf("failed to delete message index for id %s: %v", doc.Id, err)
		}
	}
	return nil
}
