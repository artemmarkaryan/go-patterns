package pool

import (
	"log"
	"math/rand"
	"time"
)

//Resource : ресурс
type Resource struct {
	resId int
}

//NewResource : что-то медленное
// (e.g., TCP connection, SSL symmetric key acquisition, auth authentication are time-consuming)
func NewResource(id int) *Resource {
	time.Sleep(500 * time.Millisecond)
	return &Resource{resId: id}
}

//Do : Simulation resources are time consuming and random consumption is 0~400ms
func (r *Resource) Do(workId int) {
	time.Sleep(time.Duration(rand.Intn(5)) * 100 * time.Millisecond)
	log.Printf("using resource #%d finished work %d finish\n", r.resId, workId)
}

//Pool использует канал Го, чтобы работать многопоточно
type Pool chan *Resource

//New : кладём в пул ресурсы
func New(size int) Pool {
	p := make(Pool, size)
	for i := 0; i < size; i++ {
		go func(resId int) {
			p <- NewResource(resId)
		}(i)
	}
	return p
}

//GetResource получает ресурс
func (p Pool) GetResource() (r *Resource, err error) {
	select {
	case r := <-p:
		return r, nil
	}
}

//GiveBackResource returns resources to the resource pool
func (p Pool) GiveBackResource(r *Resource) {
	p <- r
}
