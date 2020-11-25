package matrix

import (
	"reflect"
	"testing"
)

func Test_newSimpleFormattedMessage(t *testing.T) {
	type args struct {
		formattedBody string
	}
	tests := []struct {
		name string
		args args
		want EventFormattedMessage
	}{
		{
			name: "emptyBodyTest",
			args: args{""},
			want: EventFormattedMessage{
				MsgType:       "m.text",
				Body:          "",
				Format:        "org.matrix.custom.html",
				FormattedBody: "",
			},
		},
		{
			name: "nonHtmlBodyTest",
			args: args{"Hello world"},
			want: EventFormattedMessage{
				MsgType:       "m.text",
				Body:          "Hello world",
				Format:        "org.matrix.custom.html",
				FormattedBody: "Hello world",
			},
		},
		{
			name: "basicHtmlBodyTest",
			args: args{"<b>Hello</b> world"},
			want: EventFormattedMessage{
				MsgType:       "m.text",
				Body:          "Hello world",
				Format:        "org.matrix.custom.html",
				FormattedBody: "<b>Hello</b> world",
			},
		},
		{
			name: "htmlParagraphBodyTest",
			args: args{"<p>Hello</p><p>world</p>"},
			want: EventFormattedMessage{
				MsgType:       "m.text",
				Body:          " Hello  world ",
				Format:        "org.matrix.custom.html",
				FormattedBody: "<p>Hello</p><p>world</p>",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSimpleFormattedMessage(tt.args.formattedBody); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSimpleFormattedMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
