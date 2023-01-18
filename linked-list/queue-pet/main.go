package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type AnimalKind int8

const (
	AnimalKindAny AnimalKind = iota
	AnimalKindCat
	AnimalKindDog
)

func (a AnimalKind) String() string {
	switch a {
	case AnimalKindCat:
		return "C"
	case AnimalKindDog:
		return "D"
	default:
		return "A"
	}
}

const sizeOfKind = 3

type Animal struct {
	id         uuid.UUID
	kind       AnimalKind
	nxGeneral  *Animal
	preGeneral *Animal
	nxSame     *Animal
}

func NewAnimal(k AnimalKind) (*Animal, error) {
	if k == AnimalKindAny {
		return nil, fmt.Errorf("wrong animal kind at new animal")
	}
	return &Animal{
		id:   uuid.New(),
		kind: k,
	}, nil
}

func (a *Animal) idString() string {
	return a.id.String()[:4]
}

func (a *Animal) String() string {
	if a.nxSame != nil {
		return fmt.Sprintf("%v:%v -> %v", a.kind, a.idString(), a.nxSame)
	}
	return fmt.Sprintf("%v:%v", a.kind, a.idString())
}

func (a *Animal) StringAsGeneral() string {
	if a.nxGeneral != nil {
		return fmt.Sprintf("%v:%v -> %v", a.kind, a.idString(), a.nxGeneral.StringAsGeneral())
	}
	return fmt.Sprintf("%v:%v", a.kind, a.idString())
}

func (a *Animal) CleanLinks() {
	if a == nil {
		return
	}
	a.nxGeneral = nil
	a.nxSame = nil
	a.preGeneral = nil
}

type AnimalQueue struct {
	top  []*Animal
	last []*Animal
}

func (a *Animal) Link(b *Animal, linkType AnimalKind) error {
	if linkType != AnimalKindAny && linkType != a.kind {
		return fmt.Errorf("wrong animal kind at link")
	}
	if linkType == AnimalKindAny {
		a.nxGeneral = b
		b.preGeneral = a
	} else {
		a.nxSame = b
	}
	return nil
}

func NewQueue() *AnimalQueue {
	return &AnimalQueue{
		top:  make([]*Animal, sizeOfKind),
		last: make([]*Animal, sizeOfKind),
	}
}

func (q *AnimalQueue) Dump() {
	for i := 0; i < sizeOfKind; i++ {
		k := AnimalKind(i)
		fmt.Println("--- for animal:", k)
		if q.top[i] == nil {
			fmt.Println("queue is empty")
			continue
		}
		if k == AnimalKindAny {
			fmt.Println(q.top[i].StringAsGeneral())
		} else {
			fmt.Println(q.top[i])
		}
		fmt.Println("last: ", q.last[i])
	}
}

func (q *AnimalQueue) enqueue(a *Animal, target AnimalKind) error {
	if q.last[target] != nil {
		err := q.last[target].Link(a, target)
		if err != nil {
			return fmt.Errorf("enqueue: %w", err)
		}
	} else {
		if q.top[target] != nil {
			return fmt.Errorf("inconsistent state: last is nil but top isn't nil")
		}
		q.top[target] = a
	}
	q.last[target] = a
	return nil
}

func (q *AnimalQueue) Enqueue(a *Animal) error {
	if a.kind == AnimalKindAny {
		return fmt.Errorf("animal kind isn't specified")
	}
	q.enqueue(a, AnimalKindAny)
	switch a.kind {
	case AnimalKindCat:
		q.enqueue(a, AnimalKindCat)
	case AnimalKindDog:
		q.enqueue(a, AnimalKindDog)
	}
	return nil
}

func (q *AnimalQueue) Dequeue(k AnimalKind) (anm *Animal) {
	if k == AnimalKindAny {
		anm = q.top[k]
		if anm == nil {
			return
		}
		q.top[k] = anm.nxGeneral
		if anm.nxGeneral == nil {
			q.last[k] = nil
		} else {
			anm.nxGeneral.preGeneral = nil
		}
		q.top[anm.kind] = anm.nxSame
		if anm.nxSame == nil {
			q.last[anm.kind] = nil
		}
		anm.CleanLinks()
		return
	}
	anm = q.top[k]
	if anm == nil {
		return
	}
	if anm.kind != k {
		anm.nxSame = nil
		panic(fmt.Sprintf("inconsistent state of queue: request was %v but %v is returned", k, anm))
	}
	if anm.preGeneral != nil {
		anm.preGeneral.nxGeneral = anm.nxGeneral
	}
	if q.top[AnimalKindAny] == anm {
		q.top[AnimalKindAny] = anm.nxGeneral
		if anm.nxGeneral == nil {
			q.last[AnimalKindAny] = nil
		}
	}
	q.top[k] = anm.nxSame
	if anm.nxSame == nil {
		q.last[k] = nil
	}
	anm.CleanLinks()
	return
}

func main() {
	rand.Seed(3)
	qu := NewQueue()
	for i := 0; i < 10; i++ {
		k := rand.Intn(2) + 1
		a, err := NewAnimal(AnimalKind(k))
		fmt.Println("push: ", a)
		if err != nil {
			panic(err)
		}
		qu.Enqueue(a)
	}
	qu.Dump()

	fmt.Println(qu.Dequeue(AnimalKindAny))
	fmt.Println(qu.Dequeue(AnimalKindCat))
	fmt.Println(qu.Dequeue(AnimalKindDog))
	fmt.Println(qu.Dequeue(AnimalKindAny))
	fmt.Println(qu.Dequeue(AnimalKindAny))
	fmt.Println(qu.Dequeue(AnimalKindDog))
	fmt.Println(qu.Dequeue(AnimalKindDog))
	fmt.Println(qu.Dequeue(AnimalKindDog))
	fmt.Println(qu.Dequeue(AnimalKindDog))
	fmt.Println(qu.Dequeue(AnimalKindDog))
	fmt.Println(qu.Dequeue(AnimalKindAny))
	fmt.Println(qu.Dequeue(AnimalKindAny))
	fmt.Println(qu.Dequeue(AnimalKindAny))
	fmt.Println(qu.Dequeue(AnimalKindAny))
	fmt.Println(qu.Dequeue(AnimalKindAny))
}
