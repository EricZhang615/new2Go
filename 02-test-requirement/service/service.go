package service

import (
	"sync"
	"test-sample/repository"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

type QueryPageInfoFlow struct {
	TopicId  int64
	PageInfo *PageInfo
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		f.PageInfo.Topic = repository.NewTopicDaoInstance().QueryTopicById(f.TopicId)
	}()
	go func() {
		defer wg.Done()
		f.PageInfo.PostList = repository.NewPostDaoInstance().QueryPostsById(f.TopicId)
	}()
	wg.Wait()

	return nil
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	return f.PageInfo, nil
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	queryPageInfoFlow := QueryPageInfoFlow{TopicId: topicId, PageInfo: &PageInfo{}}
	return queryPageInfoFlow.Do()
}
