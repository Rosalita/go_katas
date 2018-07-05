// package main

// import(
// 	"testing"
// 	"github.com/stretchr/testify/assert"
// )

// func TestReadStory(t *testing.T){
// 	var tests = []struct {
// 		path string
// 		data string
// 		errtext string
// 	}{
// 		{"testdata/test.json", "", ""},
// 	}
// 	for _, test := range tests {
// 		data, err := readStory(test.path)
// 		assert.Equal(t, test.data, data, "unexpected data returned")
// 		assert.Equal(t, test.errtext, err.Error(), "unexpected error returned")

// 	}
// }
