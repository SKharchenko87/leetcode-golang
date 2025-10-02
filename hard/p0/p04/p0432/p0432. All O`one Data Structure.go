package p0432

import "container/list"

type AllOne struct {
	entry   map[string]*list.Element
	buckets *list.List
}

type Bucket struct {
	count    int
	elements map[string]bool
}

func Constructor() AllOne {
	return AllOne{
		entry:   make(map[string]*list.Element),
		buckets: list.New(),
	}
}

func (this *AllOne) Inc(key string) {
	if curr, found := this.entry[key]; found {
		currBucket := curr.Value.(*Bucket)

		cnt := currBucket.count + 1
		next := curr.Next()
		if next != nil && next.Value.(*Bucket).count == cnt {
			next.Value.(*Bucket).elements[key] = true
			this.entry[key] = next
		} else {
			bucket := &Bucket{
				count:    cnt,
				elements: map[string]bool{key: true},
			}
			if next == nil {
				newElement := this.buckets.PushBack(bucket)
				this.entry[key] = newElement
			} else {
				newElement := this.buckets.InsertAfter(bucket, curr)
				this.entry[key] = newElement
			}
		}

		// clean old elements
		delete(currBucket.elements, key)
		if len(currBucket.elements) == 0 {
			this.buckets.Remove(curr)
		}
	} else {
		curr := this.buckets.Front()
		if curr != nil && curr.Value.(*Bucket).count == 1 {
			curr.Value.(*Bucket).elements[key] = true
			this.entry[key] = curr
		} else {
			// new element
			bucket := &Bucket{
				count:    1,
				elements: map[string]bool{key: true},
			}
			newElement := this.buckets.PushFront(bucket)
			this.entry[key] = newElement
		}
	}
}

func (this *AllOne) Dec(key string) {
	if curr, found := this.entry[key]; found {
		currBucket := curr.Value.(*Bucket)
		cnt := currBucket.count - 1

		if cnt == 0 {
			delete(this.entry, key)
			delete(currBucket.elements, key)
			if len(currBucket.elements) == 0 {
				this.buckets.Remove(curr)
			}
			return
		}

		prev := curr.Prev()
		if prev != nil && prev.Value.(*Bucket).count == cnt {
			prev.Value.(*Bucket).elements[key] = true
			this.entry[key] = prev
		} else {
			bucket := &Bucket{
				count:    cnt,
				elements: map[string]bool{key: true},
			}

			if prev == nil {
				newElement := this.buckets.PushFront(bucket)
				this.entry[key] = newElement
			} else {
				newElement := this.buckets.InsertAfter(bucket, prev)
				this.entry[key] = newElement
			}
		}

		// clean old elements
		delete(currBucket.elements, key)
		if len(currBucket.elements) == 0 {
			this.buckets.Remove(curr)
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.buckets.Len() == 0 {
		return ""
	}

	curr := this.buckets.Back()
	currBucket := curr.Value.(*Bucket)
	for v, _ := range currBucket.elements {
		return v
	}
	return ""
}

func (this *AllOne) GetMinKey() string {

	if this.buckets.Len() == 0 {
		return ""
	}
	curr := this.buckets.Front()
	currBucket := curr.Value.(*Bucket)
	for v, _ := range currBucket.elements {
		return v
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
