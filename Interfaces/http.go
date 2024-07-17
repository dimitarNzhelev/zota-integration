package Interfaces

type Http interface {
    getPayload() []byte
    getURL() string
}

