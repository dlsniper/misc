package hashtables

import "fmt"

type (
	HashTable interface {
		hash(key int) int
		Insert(key int, value string)
		Find(key int) (int, string)
		Delete(key int)
	}

	kv struct {
		key int
		val string
	}

	Hashtable struct {
		noBuckets int
		table     map[int][]kv
	}
)

func (ht *Hashtable) Hash(key int) int {
	return key % ht.noBuckets
}

func (ht *Hashtable) Insert(key int, value string) {
	code := ht.Hash(key)
	ht.table[code] = append(ht.table[code], kv{key: key, val: value})
}

func (ht *Hashtable) Find(key int) (int, string) {
	code := ht.Hash(key)

	for _, pair := range ht.table[code] {
		if pair.key == key {
			return pair.key, pair.val
		}
	}

	return -1, ""
}

func (ht *Hashtable) Delete(key int) {
	code := ht.Hash(key)

	for k, pair := range ht.table[code] {
		if pair.key == key {
			ht.table[code] = append(ht.table[code][:k], ht.table[code][k+1:]...)
			return
		}
	}
}

func (ht *Hashtable) String() string {
	result := ""

	for key, list := range ht.table {
		for k, v := range list {
			result += fmt.Sprintf("%d: (%d) %d -> %q\n", key, k, v.key, v.val)
		}
	}

	return result
}

func NewHashtable(noBuckets int) *Hashtable {
	hashtable := &Hashtable{
		noBuckets: noBuckets,
		table:     make(map[int][]kv, noBuckets),
	}
	return hashtable
}
