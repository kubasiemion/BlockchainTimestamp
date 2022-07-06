package api

import (
	"fmt"
	"time"
)

func (nt *Note) BlockTimeString() string {
	return fmt.Sprint(time.Unix(nt.BlockTime.Int64(), 0).Format("2006-01-02 15:04:05"))
}
