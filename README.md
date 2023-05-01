# go-json-bench

# How to run

```shell
go test -bench=. -benchmem
```

Results

|               |         |       | ns/op | B/op | allocs/op |
| ------------- | ------- | ----- | ----- | ---- | --------- |
| go-simplejson | 501222  | 2395  | 2166  | 39   |
| insane-json   | 4150257 | 289.2 | 0     | 0    |
