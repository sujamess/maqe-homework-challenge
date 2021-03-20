# [MAQE Bot](https://maqe.github.io/maqe-bot.html) - Challenge

## How to test?
```
go build && ./maqe-bot [COMMAND]
```

## Benchmark
- Benchmarking with
  - Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz Ram 16GB
  - 100 command length `RW15RW1LLLLLW99RRRRRW88LLLRLW55555W555555W444444W1RRW11RLLW19RRW12LW1W55555RW555555W444444W1RLW23LRR`
- Result
```
BenchmarkBotExecuteCommandWithMaqeBot-8   	  165523	      7205 ns/op	    2480 B/op	     188 allocs/op
```