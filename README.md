# template

testing a template engine based on ideas behind jingo vs JetHTML


## want

```html
<html>
    <body>
        <h1>Bob</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
        
            <li>blue</li>
            <li>green</li>
            <li>mauve</li>
        </ul>
    </body>
</html>
```


## got

```html
<html>
    <body>
        <h1>Bob</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
        { { range .FavoriteColors }} <------broken....
            <li>{ { . }}</li>{ { end }}
        </ul>
    </body>
</html>

```

## current numbers

```text
BenchmarkTemplate-16                    38205490                28.1 ns/op             0 B/op          0 allocs/op
BenchmarkTemplateWithPool-16            29728286                39.9 ns/op             0 B/op          0 allocs/op
BenchmarkJetHTMLTemplate-16              3925189               302 ns/op              32 B/op          1 allocs/op
```
