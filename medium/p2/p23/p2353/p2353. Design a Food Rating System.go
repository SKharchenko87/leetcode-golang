package main

import "container/heap"

type FoodRatings struct {
	index     map[string]*pq
	foodIndex map[string]*pqItem
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {

	fr := FoodRatings{}
	fr.index = make(map[string]*pq)
	fr.foodIndex = make(map[string]*pqItem)

	for i := 0; i < len(foods); i++ {
		if fr.index[cuisines[i]] == nil {
			fr.index[cuisines[i]] = new(pq)
		}

		el := &pqItem{
			cuisine: cuisines[i],
			food:    foods[i],
			rating:  ratings[i],
		}

		heap.Push(fr.index[cuisines[i]], el)

		fr.foodIndex[foods[i]] = el
	}

	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	if f, ok := this.foodIndex[food]; ok {
		f.rating = newRating
		heap.Fix(this.index[f.cuisine], f.index)
	}
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	if q, ok := this.index[cuisine]; ok {
		el := q.Peek().(*pqItem)
		return el.food
	}

	return ""
}

type pq struct {
	arr []*pqItem
}

func (p *pq) Len() int {
	return len(p.arr)
}

func (p *pq) Less(i, j int) bool {
	if p.arr[i].rating == p.arr[j].rating {
		return p.arr[i].food < p.arr[j].food
	}
	return p.arr[i].rating > p.arr[j].rating
}

func (p *pq) Swap(i, j int) {
	p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
	p.arr[i].index = i
	p.arr[j].index = j
}

func (p *pq) Push(x any) {
	el := x.(*pqItem)
	el.index = len(p.arr)

	p.arr = append(p.arr, el)
}

func (p *pq) Pop() any {
	el := p.arr[len(p.arr)-1]

	p.arr = p.arr[:len(p.arr)-1]

	return el
}

func (p *pq) Peek() any {
	el := p.arr[0]

	return el
}

type pqItem struct {
	cuisine string
	food    string
	rating  int
	index   int
}

type constructorType struct {
	foods    []string
	cuisines []string
	ratings  []int
}

type changeRatingType struct {
	foods     string
	newRating int
}

func run(commands []string, param []any) []any {
	result := []any{}
	var foodRatings FoodRatings
	for i, command := range commands {
		switch command {
		case "FoodRatings":
			foodRatingParams, _ := param[i].(constructorType)
			foodRatings = Constructor(foodRatingParams.foods, foodRatingParams.cuisines, foodRatingParams.ratings)
			result = append(result, nil)
		case "highestRated":
			rating := foodRatings.HighestRated(param[i].(string))
			result = append(result, rating)
		case "changeRating":
			changeRatingParams := param[i].(changeRatingType)
			foodRatings.ChangeRating(changeRatingParams.foods, changeRatingParams.newRating)
			result = append(result, nil)
		}
	}
	return result
}
