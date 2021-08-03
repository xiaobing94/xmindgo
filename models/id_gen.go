package models

import "github.com/google/uuid"

type IDComponent struct {
	ID string `json:"id,omitempty"`
}

func (c *IDComponent) GenID() {
	// 忽略这个err，即使是默认Nil，也可以使用
	uid, _ := uuid.NewUUID()
	c.ID = uid.String()
}
