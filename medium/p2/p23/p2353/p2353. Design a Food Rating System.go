package main

import "fmt"

type FoodRating struct {
	cuisines string
	ratings  int
}

type FoodRatings struct {
	foods         map[string]FoodRating
	ratingCuisine map[string]string
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	l := len(foods)
	m := make(map[string]FoodRating, l)
	r := make(map[string]string, l)
	for i := 0; i < l; i++ {
		m[foods[i]] = FoodRating{cuisines[i], ratings[i]}
		if v, ok := r[cuisines[i]]; ok {
			if m[v].ratings < m[foods[i]].ratings {
				r[cuisines[i]] = foods[i]
			} else if m[v].ratings == m[foods[i]].ratings {
				if foods[i] <= v {
					r[cuisines[i]] = foods[i]
				}
			}
		} else {
			r[cuisines[i]] = foods[i]
		}
	}
	return FoodRatings{m, r}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := this.foods[food].cuisines
	this.foods[food] = FoodRating{cuisine, newRating}
	if this.foods[this.ratingCuisine[cuisine]].ratings < newRating {
		this.ratingCuisine[cuisine] = food
	} else if this.foods[this.ratingCuisine[cuisine]].ratings == newRating {
		if food < this.ratingCuisine[cuisine] {
			this.ratingCuisine[cuisine] = food
		}
	}
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.ratingCuisine[cuisine]
}

func main() {

	x := Constructor([]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}, []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7})
	fmt.Println(x.HighestRated("korean"))
	fmt.Println(x.HighestRated("japanese"))
	x.ChangeRating("sushi", 16)
	fmt.Println(x.HighestRated("japanese"))
	x.ChangeRating("ramen", 16)
	fmt.Println(x.HighestRated("japanese"))

}
