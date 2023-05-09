package types

import "github.com/zgwit/iot-master/v3/model"

type Sms struct {
	Id        int64      `json:"id"`
	Type      string     `json:"type,omitempty"`
	Cellphone string     `json:"cellphone,omitempty"`
	Content   string     `json:"content,omitempty"`
	Result    string     `json:"result,omitempty"`
	Created   model.Time `json:"created,omitempty" xorm:"created"`
}
