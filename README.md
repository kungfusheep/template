#template

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