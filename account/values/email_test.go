package values

import "testing"

func TestIsValidEmailAddress(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		want bool
	}{
		{
			desc: "Happy path",
			in:   "testing@gmail.com",
			want: true,
		},
		{
			desc: "Missing @",
			in:   "testgmail.com",
			want: false,
		},
		{
			desc: "invalid domain name",
			in:   "test@111+ytryertyerty.yrte",
			want: false,
		},
		{
			desc: "Too short",
			in:   "aa",
			want: false,
		},
		{
			desc: "Too long",
			in:   "aaaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllll12345@aaaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllll12345.long",
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := isValidEmailAddress(tC.in)
			if got != tC.want {
				t.Errorf("Desc: %s\tValue in: %s\t got %v want %v", tC.desc, tC.in, got, tC.want)
			}

		})
	}
}

func Benchmark(b *testing.B) {
	var got bool
	in := "testing@gmail.com"
	for i := 0; i < b.N; i++ {
		got = isValidEmailAddress(in)
	}
	if !got {
		b.Errorf("%s was not valid", in)
	}

}
