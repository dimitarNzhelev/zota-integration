package Implementations

import (
	"zota_integration/Interfaces"
)

type HttpStruct struct {
    payload string
    url     string
}

// func (h *HttpStruct) getPayload() []byte {
//     return []byte(h.payload)
//  }
 
//  func (h *HttpStruct) getURL() string {
//      return h.url
//  }


func newHttpStruct(url, payload string) Interfaces.Http {
    h := HttpStruct{url: url, payload: payload}
    return &h
}


