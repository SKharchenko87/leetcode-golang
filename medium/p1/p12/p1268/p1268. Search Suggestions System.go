package p1268

import (
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	lengthSearchWord := len(searchWord)
	res := make([][]string, lengthSearchWord)
out:
	for _, product := range products {
		for i := 0; i < min(lengthSearchWord, len(product)); i++ {
			if product[i] == searchWord[i] {
				if len(res[i]) == 3 {
					continue
				}
				res[i] = append(res[i], product)
			} else if product[i] < searchWord[i] {
				break
			} else if i == 0 && product[i] < searchWord[i] {
				break out
			} else {
				break
			}
		}
	}
	return res
}

//func suggestedProducts(products []string, searchWord string) [][]string {
//	lengthSearchWord := len(searchWord)
//	res := make([][]string, lengthSearchWord)
//	sort.Strings(products)
//	for i := 0; i < lengthSearchWord; i++ {
//		var newTmp []string
//		for _, product := range products {
//			if product[i] == searchWord[i] {
//				newTmp = append(newTmp, product)
//			} else if product[i] > searchWord[i] {
//				break
//			}
//		}
//		res[i] = newTmp[:min(len(newTmp), 3)]
//		products = newTmp
//		if len(products) == 0 {
//			return res
//		}
//	}
//	return res
//}
