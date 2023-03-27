package rabbitmq

import (
	"errors"
	"strings"
)

func (r *rabbitmq) Delete(queueName string) (int, error) {
	if strings.TrimSpace(queueName) == "" {
		return 0, errors.New("can't delete a queue without a queueName")
	}

	purgedMessageCount, err := r.Channel.QueueDelete(queueName, false, false, true)
	if err != nil {
		return 0, err
	}

	return purgedMessageCount, nil
}
