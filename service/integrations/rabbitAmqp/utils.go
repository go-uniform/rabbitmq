package rabbitAmqp

import (
    "crypto/tls"
    "fmt"
    rabbit "github.com/wagslane/go-rabbitmq"
    "strings"
)

func CompileRabbitAmqpConfig(uri, certFile, keyFile string, disableAmqps bool) (string, rabbit.Config) {
    config := rabbit.Config{}

    if !strings.HasPrefix(uri, "amqp://") && !strings.HasPrefix(uri, "amqps://") {
        uri = fmt.Sprintf("amqp://%s", uri)
    }

    if !disableAmqps {
        uri = strings.Replace(uri, "amqp://", "amqps://", -1)

        cert, err := tls.LoadX509KeyPair(certFile, keyFile)
        if err != nil {
            panic(err)
        }

        config.TLSClientConfig = &tls.Config{
            // since RabbitMQ should always be on a private line with self-signed certs, we just skip host verification
            InsecureSkipVerify: true,

            Certificates: []tls.Certificate{cert},
            MinVersion:   tls.VersionTLS12,
        }
    } else {
        uri = strings.Replace(uri, "amqps://", "amqp://", -1)
    }

    return uri, config
}
