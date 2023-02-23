# redacter

json value redacter

## Examples

- raw json

```go
raw := []byte(`{
  "name": "name",
  "secrets": [
   { "id": 1, "name": "secret-1" },
   { "id": 2, "name": "secret-2" },
   { "id": 3, "name": "secret-2", "shared": true }
  ]
 }`)

 res, err := redacter.New().
  WithFields("name").
  JSONRaw(raw)
```

- object

```go
p := Person{ ... }

 res, err := redacter.New().
  WithFields("password", "country").
  WithStrict(false).
  WithRemove(true).
  JSON(p)
```
