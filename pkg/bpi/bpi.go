package bpi

import (
	"bilibili/pkg/client"
	"bilibili/pkg/endpoints/article"
	"bilibili/pkg/endpoints/audio"
	"bilibili/pkg/endpoints/video"
	"sync"
)

type BpiService struct {
	Client *client.Client

	articleOnce sync.Once
	audioOnce   sync.Once
	videoOnce   sync.Once

	article *article.Article
	audio   *audio.Audio
	video   *video.Video
}

func New(cli *client.Client) *BpiService {
	return &BpiService{
		Client: cli,
	}
}

func (s *BpiService) Article() *article.Article {
	s.articleOnce.Do(func() {
		s.article = article.New(s.Client)
	})
	return s.article
}

func (s *BpiService) Audio() *audio.Audio {
	s.audioOnce.Do(func() {
		s.audio = audio.New(s.Client)
	})
	return s.audio
}

func (s *BpiService) Video() *video.Video {
	s.videoOnce.Do(func() {
		s.video = video.New(s.Client)
	})
	return s.video
}
