package template

import (
	"bytes"
	"github.com/CloudyKit/jet"
	"testing"
)

type User struct {
	FirstName      string
	FavoriteColors []string
}

var (
	test = []byte(`
<html>
    <body>
        <h1>Hi {{ .FirstName }}</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
        { { range .FavoriteColors }} <------broken....
            <li>{ { . }}</li>{ { end }}
        </ul>
    </body>
</html>
`)

	testData = &User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
	}
)

func BenchmarkTemplate(b *testing.B) {

	var tpl = New(test, User{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := NewBufferFromPool()
		tpl.Execute(testData, buf)
		buf.ReturnToPool()
	}
}

func ExampleTemplate() {
	println("hi")
	var tpl = New(test, User{})

	buf := NewBufferFromPool()
	tpl.Execute(testData, buf)
	buf.ReturnToPool()

	println(buf.String())

	// Output:
	// wrong
}

var jetSet = jet.NewHTMLSet("./jet")

func BenchmarkJetHTMLTemplate(b *testing.B) {
	var buf bytes.Buffer

	tmpl, _ := jetSet.GetTemplate("simple.jet")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tmpl.Execute(&buf, nil, testData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

func ExampleJetHTMLTemplate() {
	var buf bytes.Buffer

	tmpl, _ := jetSet.GetTemplate("simple.jet")

	tmpl.Execute(&buf, nil, testData)

	println(buf.String())

	// Output:
	// garbbs
}
