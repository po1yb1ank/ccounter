package watcher

type UpdateMessage struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}
