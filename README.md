github.com/dalu/mail

Usage:

```go
m := NewMail("from@example.com", []string{"to@example.com","too@example.com"}, "my cool subject", "my cool\nmessage\nyeah")
if err := m.Send(); err != nil {
	panic(err)
}
```
