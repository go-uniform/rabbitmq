package rabbitApi

type Queue struct {
    Name                   string `json:"name"`
    MessagesReady          int64  `json:"messages_ready"`
    MessagesUnacknowledged int64  `json:"messages_unacknowledged"`
}
