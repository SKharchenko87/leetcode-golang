package p0567

func checkInclusion(s1 string, s2 string) bool {
	// If s1 is longer than s2, return false as it's impossible for s2 to contain a permutation of s1
	if len(s1) > len(s2) {
		return false
	}

	// Helper function to check if two frequency maps are equal
	areEqual := func(a, b [26]int) bool {
		for i := 0; i < 26; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	// Frequency array for s1 and for the current window in s2
	var s1Count, s2Count [26]int

	// Fill the s1Count frequency array for the string s1
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	// Check the first window (of length len(s1))
	if areEqual(s1Count, s2Count) {
		return true
	}

	// Slide the window over the rest of s2
	for i := len(s1); i < len(s2); i++ {
		// Add the new character to the window
		s2Count[s2[i]-'a']++
		// Remove the character that is no longer in the window
		s2Count[s2[i-len(s1)]-'a']--

		// Check if the current window matches the frequency of s1
		if areEqual(s1Count, s2Count) {
			return true
		}
	}

	return false
}
