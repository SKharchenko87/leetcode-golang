package p2843

import "slices"

var symmetricIntegers = []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 1001, 1010, 1102, 1111, 1120, 1203, 1212, 1221, 1230, 1304, 1313, 1322, 1331, 1340, 1405, 1414, 1423, 1432, 1441, 1450, 1506, 1515, 1524, 1533, 1542, 1551, 1560, 1607, 1616, 1625, 1634, 1643, 1652, 1661, 1670, 1708, 1717, 1726, 1735, 1744, 1753, 1762, 1771, 1780, 1809, 1818, 1827, 1836, 1845, 1854, 1863, 1872, 1881, 1890, 1919, 1928, 1937, 1946, 1955, 1964, 1973, 1982, 1991, 2002, 2011, 2020, 2103, 2112, 2121, 2130, 2204, 2213, 2222, 2231, 2240, 2305, 2314, 2323, 2332, 2341, 2350, 2406, 2415, 2424, 2433, 2442, 2451, 2460, 2507, 2516, 2525, 2534, 2543, 2552, 2561, 2570, 2608, 2617, 2626, 2635, 2644, 2653, 2662, 2671, 2680, 2709, 2718, 2727, 2736, 2745, 2754, 2763, 2772, 2781, 2790, 2819, 2828, 2837, 2846, 2855, 2864, 2873, 2882, 2891, 2929, 2938, 2947, 2956, 2965, 2974, 2983, 2992, 3003, 3012, 3021, 3030, 3104, 3113, 3122, 3131, 3140, 3205, 3214, 3223, 3232, 3241, 3250, 3306, 3315, 3324, 3333, 3342, 3351, 3360, 3407, 3416, 3425, 3434, 3443, 3452, 3461, 3470, 3508, 3517, 3526, 3535, 3544, 3553, 3562, 3571, 3580, 3609, 3618, 3627, 3636, 3645, 3654, 3663, 3672, 3681, 3690, 3719, 3728, 3737, 3746, 3755, 3764, 3773, 3782, 3791, 3829, 3838, 3847, 3856, 3865, 3874, 3883, 3892, 3939, 3948, 3957, 3966, 3975, 3984, 3993, 4004, 4013, 4022, 4031, 4040, 4105, 4114, 4123, 4132, 4141, 4150, 4206, 4215, 4224, 4233, 4242, 4251, 4260, 4307, 4316, 4325, 4334, 4343, 4352, 4361, 4370, 4408, 4417, 4426, 4435, 4444, 4453, 4462, 4471, 4480, 4509, 4518, 4527, 4536, 4545, 4554, 4563, 4572, 4581, 4590, 4619, 4628, 4637, 4646, 4655, 4664, 4673, 4682, 4691, 4729, 4738, 4747, 4756, 4765, 4774, 4783, 4792, 4839, 4848, 4857, 4866, 4875, 4884, 4893, 4949, 4958, 4967, 4976, 4985, 4994, 5005, 5014, 5023, 5032, 5041, 5050, 5106, 5115, 5124, 5133, 5142, 5151, 5160, 5207, 5216, 5225, 5234, 5243, 5252, 5261, 5270, 5308, 5317, 5326, 5335, 5344, 5353, 5362, 5371, 5380, 5409, 5418, 5427, 5436, 5445, 5454, 5463, 5472, 5481, 5490, 5519, 5528, 5537, 5546, 5555, 5564, 5573, 5582, 5591, 5629, 5638, 5647, 5656, 5665, 5674, 5683, 5692, 5739, 5748, 5757, 5766, 5775, 5784, 5793, 5849, 5858, 5867, 5876, 5885, 5894, 5959, 5968, 5977, 5986, 5995, 6006, 6015, 6024, 6033, 6042, 6051, 6060, 6107, 6116, 6125, 6134, 6143, 6152, 6161, 6170, 6208, 6217, 6226, 6235, 6244, 6253, 6262, 6271, 6280, 6309, 6318, 6327, 6336, 6345, 6354, 6363, 6372, 6381, 6390, 6419, 6428, 6437, 6446, 6455, 6464, 6473, 6482, 6491, 6529, 6538, 6547, 6556, 6565, 6574, 6583, 6592, 6639, 6648, 6657, 6666, 6675, 6684, 6693, 6749, 6758, 6767, 6776, 6785, 6794, 6859, 6868, 6877, 6886, 6895, 6969, 6978, 6987, 6996, 7007, 7016, 7025, 7034, 7043, 7052, 7061, 7070, 7108, 7117, 7126, 7135, 7144, 7153, 7162, 7171, 7180, 7209, 7218, 7227, 7236, 7245, 7254, 7263, 7272, 7281, 7290, 7319, 7328, 7337, 7346, 7355, 7364, 7373, 7382, 7391, 7429, 7438, 7447, 7456, 7465, 7474, 7483, 7492, 7539, 7548, 7557, 7566, 7575, 7584, 7593, 7649, 7658, 7667, 7676, 7685, 7694, 7759, 7768, 7777, 7786, 7795, 7869, 7878, 7887, 7896, 7979, 7988, 7997, 8008, 8017, 8026, 8035, 8044, 8053, 8062, 8071, 8080, 8109, 8118, 8127, 8136, 8145, 8154, 8163, 8172, 8181, 8190, 8219, 8228, 8237, 8246, 8255, 8264, 8273, 8282, 8291, 8329, 8338, 8347, 8356, 8365, 8374, 8383, 8392, 8439, 8448, 8457, 8466, 8475, 8484, 8493, 8549, 8558, 8567, 8576, 8585, 8594, 8659, 8668, 8677, 8686, 8695, 8769, 8778, 8787, 8796, 8879, 8888, 8897, 8989, 8998, 9009, 9018, 9027, 9036, 9045, 9054, 9063, 9072, 9081, 9090, 9119, 9128, 9137, 9146, 9155, 9164, 9173, 9182, 9191, 9229, 9238, 9247, 9256, 9265, 9274, 9283, 9292, 9339, 9348, 9357, 9366, 9375, 9384, 9393, 9449, 9458, 9467, 9476, 9485, 9494, 9559, 9568, 9577, 9586, 9595, 9669, 9678, 9687, 9696, 9779, 9788, 9797, 9889, 9898, 9999}

func countSymmetricIntegers(low int, high int) int {
	lowIndex, _ := slices.BinarySearch(SymmetricIntegers, low)
	highIndex, _ := slices.BinarySearch(SymmetricIntegers, high+1)
	return highIndex - lowIndex
}

var SymmetricIntegers = make([]int, 0, 624)

func init() {
	for i := 11; i < 100; i++ {
		if i%10 == i/10%10 {
			SymmetricIntegers = append(SymmetricIntegers, i)
		}
	}
	for i := 1001; i < 10000; i++ {
		if i%10+i/10%10 == i/100%10+i/1000 {
			SymmetricIntegers = append(SymmetricIntegers, i)
		}
	}
}

func countSymmetricIntegers3(low int, high int) int {
	lowIndex, _ := slices.BinarySearch(SymmetricIntegers, low)
	highIndex, _ := slices.BinarySearch(SymmetricIntegers, high+1)
	return highIndex - lowIndex
}

func countSymmetricIntegers2(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		switch i {
		case 100:
			i = 999
		case 10000:
			continue

		}
		d := getDigits(i)
		if len(d) == 2 && d[0] == d[1] || len(d) == 4 && d[0]+d[1] == d[2]+d[3] {
			count++
		}
	}
	return count
}

func getDigits(n int) []byte {
	res := make([]byte, 0, 5)
	for n > 0 {
		res = append(res, byte(n%10))
		n /= 10
	}
	return res
}

func countSymmetricIntegers1(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		switch i {
		case 100:
			i = 999
		case 10000:
			continue

		}
		d0 := i % 10
		d1 := i / 10 % 10
		d2 := i / 100 % 10
		d3 := i / 1000 % 10
		if d0+d1 == d2+d3 || d2 == 0 && d3 == 0 && d0 == d1 {
			count++
		}
	}
	return count
}

func countSymmetricIntegers0(low int, high int) int {
	count := 0
	for i := low; i <= min(high, 99); i++ {
		if i%10 == i/10%10 {
			count++
		}
	}
	for i := max(1000, low); i <= high; i++ {
		if i%10+i/10%10 == i/100%10+i/1000 {
			count++
		}
	}
	return count
}
