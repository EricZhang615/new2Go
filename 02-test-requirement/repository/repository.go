package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type TopicDao struct {
}

type PostDao struct {
}

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
)

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "Topic.json")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "Post.json")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		postTmpMap[post.ParentId] = append(postTmpMap[post.ParentId], &post)
	}
	postIndexMap = postTmpMap
	return nil
}

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostsById(parent_id int64) []*Post {
	return postIndexMap[parent_id]
}
