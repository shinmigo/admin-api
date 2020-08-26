package utils

import (
	"fmt"
)

func UserTokenKey(userId uint64) string {
	return fmt.Sprintf("goshop:user:token::%d", userId)
}
