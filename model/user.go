

package model

import (
  "time"
)

type User struct {
  Id             string    `json:"id"`
  CreateTime     time.Time `json:"createTime"`
  UpdateTime     time.Time `json:"updateTime"`
  Email          string    `json:"email"`
  FirstName      string    `json:"firstName"`
  LastName       string    `json:"lastName"`
  MfaStatus      string    `json:"mfaStatus"`
  Note           string    `json:"note"`
  Roles          []string  `json:"roles"`
  Status         string    `json:"status"`
  SearchUsername string    `json:"searchUsername"`
}

func NewUser() *User {
  return &User{
    CreateTime:     time.Now(),
    Email:          "",
    FirstName:      "",
    LastName:       "",
    Note:           "",
    Status:         "",
    SearchUsername: "",
  }
}

func (user *User) String() string {
  return user.Id
}
