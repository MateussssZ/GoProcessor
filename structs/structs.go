package structs

import "sync"

type CoresLoad struct {
	sync.Mutex
	Cores []int
}
