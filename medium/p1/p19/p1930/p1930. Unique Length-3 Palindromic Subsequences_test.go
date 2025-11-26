package p1930

import "testing"

func Test_countPalindromicSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "aabca"}, 3},
		{"Example 2", args{s: "adc"}, 0},
		{"Example 3", args{s: "bbcbaba"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPalindromicSubsequence(tt.args.s); got != tt.want {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countPalindromicSubsequence(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPalindromicSubsequence("gczgajmriovlmxilagqfggvriouzyvmhscuhasrwuxnqcurpigifenjrtkhbgbqqjtvecycfhsetnrjzcybjjxdtngvpygsvuhqauwtjpmzfhobgkqjouliegbzxwxagzkvcewcbxsylsenifkcrphacgxqdcxxqkklpfyugimyadjqsulawkajwpxfkcabfygwpimghvotehnxehnwachadehcwgkacsipoxqqhkidvnvbqbpzychojclpkplrmepgtjwcjhhswwgrctejekvbexaedkqojawsegwchevpsfdwdaaldwoecgnqmbbzetpmrznflghefgbypjpjgpbvsrnkronoubqjowawlnaulcmqhjlzzgnfdedhqcpujzattptxbyirneenc")
	}
}

func Benchmark_countPalindromicSubsequence0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPalindromicSubsequence0("gczgajmriovlmxilagqfggvriouzyvmhscuhasrwuxnqcurpigifenjrtkhbgbqqjtvecycfhsetnrjzcybjjxdtngvpygsvuhqauwtjpmzfhobgkqjouliegbzxwxagzkvcewcbxsylsenifkcrphacgxqdcxxqkklpfyugimyadjqsulawkajwpxfkcabfygwpimghvotehnxehnwachadehcwgkacsipoxqqhkidvnvbqbpzychojclpkplrmepgtjwcjhhswwgrctejekvbexaedkqojawsegwchevpsfdwdaaldwoecgnqmbbzetpmrznflghefgbypjpjgpbvsrnkronoubqjowawlnaulcmqhjlzzgnfdedhqcpujzattptxbyirneenc")
	}
}

func Benchmark_countPalindromicSubsequence1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPalindromicSubsequence1("gczgajmriovlmxilagqfggvriouzyvmhscuhasrwuxnqcurpigifenjrtkhbgbqqjtvecycfhsetnrjzcybjjxdtngvpygsvuhqauwtjpmzfhobgkqjouliegbzxwxagzkvcewcbxsylsenifkcrphacgxqdcxxqkklpfyugimyadjqsulawkajwpxfkcabfygwpimghvotehnxehnwachadehcwgkacsipoxqqhkidvnvbqbpzychojclpkplrmepgtjwcjhhswwgrctejekvbexaedkqojawsegwchevpsfdwdaaldwoecgnqmbbzetpmrznflghefgbypjpjgpbvsrnkronoubqjowawlnaulcmqhjlzzgnfdedhqcpujzattptxbyirneenc")
	}
}

func Benchmark_countPalindromicSubsequence2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPalindromicSubsequence2("gczgajmriovlmxilagqfggvriouzyvmhscuhasrwuxnqcurpigifenjrtkhbgbqqjtvecycfhsetnrjzcybjjxdtngvpygsvuhqauwtjpmzfhobgkqjouliegbzxwxagzkvcewcbxsylsenifkcrphacgxqdcxxqkklpfyugimyadjqsulawkajwpxfkcabfygwpimghvotehnxehnwachadehcwgkacsipoxqqhkidvnvbqbpzychojclpkplrmepgtjwcjhhswwgrctejekvbexaedkqojawsegwchevpsfdwdaaldwoecgnqmbbzetpmrznflghefgbypjpjgpbvsrnkronoubqjowawlnaulcmqhjlzzgnfdedhqcpujzattptxbyirneenc")
	}
}

func Benchmark_countPalindromicSubsequence3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPalindromicSubsequence3("gczgajmriovlmxilagqfggvriouzyvmhscuhasrwuxnqcurpigifenjrtkhbgbqqjtvecycfhsetnrjzcybjjxdtngvpygsvuhqauwtjpmzfhobgkqjouliegbzxwxagzkvcewcbxsylsenifkcrphacgxqdcxxqkklpfyugimyadjqsulawkajwpxfkcabfygwpimghvotehnxehnwachadehcwgkacsipoxqqhkidvnvbqbpzychojclpkplrmepgtjwcjhhswwgrctejekvbexaedkqojawsegwchevpsfdwdaaldwoecgnqmbbzetpmrznflghefgbypjpjgpbvsrnkronoubqjowawlnaulcmqhjlzzgnfdedhqcpujzattptxbyirneenc")
	}
}

func Benchmark_countPalindromicSubsequence4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPalindromicSubsequence3("gczgajmriovlmxilagqfggvriouzyvmhscuhasrwuxnqcurpigifenjrtkhbgbqqjtvecycfhsetnrjzcybjjxdtngvpygsvuhqauwtjpmzfhobgkqjouliegbzxwxagzkvcewcbxsylsenifkcrphacgxqdcxxqkklpfyugimyadjqsulawkajwpxfkcabfygwpimghvotehnxehnwachadehcwgkacsipoxqqhkidvnvbqbpzychojclpkplrmepgtjwcjhhswwgrctejekvbexaedkqojawsegwchevpsfdwdaaldwoecgnqmbbzetpmrznflghefgbypjpjgpbvsrnkronoubqjowawlnaulcmqhjlzzgnfdedhqcpujzattptxbyirneenc")
	}
}
