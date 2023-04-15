package object_test

import (
	"github.com/jombG/object"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObject_Merge(t *testing.T) {
	tests := []struct {
		name string
		o    object.Object
		arg  object.Object
		want object.Object
	}{
		{
			name: "Basic test",
			o: object.Object{
				"foo": object.Object{
					"bar": "1",
					"baz": "2",
				},
				"bar": object.Object{
					"baz": "3",
				},
			},
			arg: object.Object{
				"foo1": "q",
				"foo": object.Object{
					"bar1": "3",
					"bar2": "4",
				},
				"bar": object.Object{
					"baz": 4,
				},
			},
			want: object.Object{
				"foo1": "q",
				"foo": object.Object{
					"bar":  "1",
					"baz":  "2",
					"bar1": "3",
					"bar2": "4",
				},
				"bar": object.Object{
					"baz": 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.o.Merge(tt.arg))
		})
	}
}

func TestObject_Property(t *testing.T) {
	tests := []struct {
		name string
		in   object.Object
		out  object.Object
	}{
		{
			name: "Basic test",
			in: object.CreateObjects(
				object.Property("foo1", "q"),
				object.Property("foo", object.CreateObjects(
					object.Property("bar", "1"),
					object.Property("baz", "2"),
					object.Property("bar1", "3"),
					object.Property("bar2", "4"),
				)),
				object.Property(
					"bar", object.CreateObjects(
						object.Property("baz", 4),
					),
				)),
			out: object.Object{
				"foo1": "q",
				"foo": object.Object{
					"bar":  "1",
					"baz":  "2",
					"bar1": "3",
					"bar2": "4",
				},
				"bar": object.Object{
					"baz": 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in)
		})
	}
}
