package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name  string
		args  args
		wantR string
	}{
		// TODO: Add test cases.
		{
			name:  "タイトルを h1 で表示",
			args:  args{
				b: "t:this is title",
			},
			wantR: "<h1>this is title</h1>\n",
		},
		{
			name:  "表の li タグの変換",
			args:  args{
				b: "- これはリストの一部です",
			},
			wantR: "<LI>これはリストの一部です</LI>\n",
		},
		{
			name:  "表の ul タグの返還",
			args:  args{
				b: "--:s\n" +
					"- item1\n" +
					"- item2\n" +
					"--:e",
			},
			wantR: "<ul>\n" +
				"<LI>item1</LI>\n" +
				"<LI>item2</LI>\n" +
				"</ul>\n",
		},
		{
			name:  "タイトル + 表",
			args:  args{
				b: "t:this is title\n" +
					"- リスト",
			},
			wantR: "<h1>this is title</h1>\n" +
				"<LI>リスト</LI>\n",
		},
		{
			name:  "タイトル + 表 + 文章",
			args:  args{
				b: "t:test\n" +
					"以下は表です\n" +
					"--:s\n" +
					"- item1\n" +
					"--:e",
			},
			wantR: "<h1>test</h1>\n" +
				"以下は表です\n" +
				"<ul>\n" +
				"<LI>item1</LI>\n" +
				"</ul>\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := convert(tt.args.b); gotR != tt.wantR {
				t.Errorf("convert() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
