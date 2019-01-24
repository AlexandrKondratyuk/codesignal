package atTheCrossroads

func main() {

}

// func tennisSet(score1 int, score2 int) bool {
//     if score1 == score2 || score1 > 7 || score2 > 7{
//         return false
//     } else if (score1 == 5 && score2 == 7) || (score2 == 5 && score1 == 7) {
// 		return true
// 	} else if (score1 == 6 && score2 < 5) || (score2 == 6 && score1 < 5) {
// 		return true
//     } else if score1>=6 && score2 >= 6 && (score1-score2 == 1 || score2-score1 == 1) {
//         return true
//     }
// 	return false
// }

func tennisSet(score1 int, score2 int) bool {
	var min, max int

	if score1 > score2 {
		min, max = score2, score1
	} else {
		min, max = score1, score2
	}

	return max==6 && min<5 || max==7 && min==6 || max==7 && min==5
}