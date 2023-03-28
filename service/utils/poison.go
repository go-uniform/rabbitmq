package utils

import (
    "fmt"
    "strings"
)

const PoisonPrefix = "_poison."

func GetPoisonQueue(queueName string) string {
    if IsPoisonQueue(queueName) {
        return queueName
    }
    return fmt.Sprintf("%s%s", PoisonPrefix, queueName)
}

func IsPoisonQueue(queueName string) bool {
    return strings.HasPrefix(queueName, PoisonPrefix)
}
