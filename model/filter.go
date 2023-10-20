

package model

import (
  "time"
)

type Filter struct {
  ImportId   string                 `json:"importId"`
  BeginTime  time.Time              `json:"beginTime"`
  EndTime    time.Time              `json:"endTime"`
  SrcIp      string                 `json:"srcIp"`
  SrcPort    int                    `json:"srcPort"`
  DstIp      string                 `json:"dstIp"`
  DstPort    int                    `json:"dstPort"`
  Parameters map[string]interface{} `json:"parameters"`
}

func NewFilter() *Filter {
  return &Filter{
    Parameters: make(map[string]interface{}),
  }
}
