package dto

import "github.com/golang/protobuf/ptypes/timestamp"

type (
	User struct {
		ID        int64                `json:"id,omitempty"`
		FirstName string               `json:"first_name,omitempty"`
		LastName  string               `json:"last_name,omitempty"`
		Gender    int32                `json:"gender,omitempty"` //FIXME: Define the gender struct
		CreatedAt *timestamp.Timestamp `json:"created_at,omitempty"`
	}
)
