package textsummary

import "testing"

func TestBasicSummarizer_Summary(t *testing.T) {
	type args struct {
		text      string
		noOfLines int
	}
	tests := []struct {
		name string
		bs   *BasicSummarizer
		args args
		want string
	}{
		{
			name: "Wrong noOfLines",
			args: args{text: "this is a sample paragraph", noOfLines: 100},
			want: "this is a sample paragraph.",
		},
		{
			name: "Extra full stops",
			args: args{text: "this is a sample paragraph...", noOfLines: 1},
			want: "this is a sample paragraph.",
		},
		{
			name: "Multiple lines with 1 line summary",
			args: args{text: "this is line 1 of the paragraph. This is line 2 of the paragraph", noOfLines: 1},
			want: "this is line 1 of the paragraph.",
		},
		{
			name: "Multiple lines with multi-line summary",
			args: args{text: "this is line 1 of the paragraph. This is line 2 of the paragraph", noOfLines: 2},
			want: "this is line 1 of the paragraph. This is line 2 of the paragraph.",
		},
		{
			name: "Testing sortSent",
			args: args{text: "So, keep working. Keep striving. Never give up. Fall down seven times, get up eight. Ease is a greater threat to progress than hardship. Ease is a greater threat to progress than hardship. So, keep moving, keep growing, keep learning. See you at work.", noOfLines: 1},
			want: "Ease is a greater threat to progress than hardship.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BasicSummarizer{}
			if got := bs.Summary(tt.args.text, tt.args.noOfLines); got != tt.want {
				t.Errorf("BasicSummarizer.Summary() = %v, want %v", got, tt.want)
			}
		})
	}
}
