package gotrie

import "testing"

func Test_lcs(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantR1     string
		wantR2     string
	}{
		{
			args:   args{s1: "hello", s2: "world"},
			wantR1: "hello", wantR2: "world",
		},
		{
			args:       args{s1: "helo", s2: "hello"},
			wantResult: "hel", wantR1: "o", wantR2: "lo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotR1, gotR2 := lcs(tt.args.s1, tt.args.s2)
			if gotResult != tt.wantResult {
				t.Errorf("lcs() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotR1 != tt.wantR1 {
				t.Errorf("lcs() gotR1 = %v, want %v", gotR1, tt.wantR1)
			}
			if gotR2 != tt.wantR2 {
				t.Errorf("lcs() gotR2 = %v, want %v", gotR2, tt.wantR2)
			}
		})
	}
}
