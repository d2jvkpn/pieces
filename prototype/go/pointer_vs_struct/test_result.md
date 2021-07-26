### pointer_vs_struct

#### BenchmarkNewD1_t1
```bash
go test -bench=BenchmarkNewD1_t1 -run=^BenchmarkNewD1_t1$ -benchmem -count 10 -v
```

```text
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkNewD1_t1
BenchmarkNewD1_t1-16    	1000000000	         0.2447 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2485 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2413 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2423 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2376 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2344 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2345 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2429 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2410 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t1-16    	1000000000	         0.2433 ns/op	       0 B/op	       0 allocs/op
PASS
```

#### BenchmarkNewD1_t2
```bash
go test -bench=BenchmarkNewD1_t2 -run=^BenchmarkNewD1_t2$ -benchmem -count 10 -v
```

```text
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkNewD1_t2
BenchmarkNewD1_t2-16    	 7375413	       155.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 7316650	       155.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 6926701	       156.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 6595531	       155.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 7354388	       156.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 7609461	       156.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 7488026	       155.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 7501852	       157.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 7524066	       158.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD1_t2-16    	 7280926	       155.8 ns/op	       0 B/op	       0 allocs/op
PASS
```

#### BenchmarkNewD2_t1
```bash
go test -bench=BenchmarkNewD2_t1 -run=^BenchmarkNewD2_t1$ -benchmem -count 10 -v
```

```text
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkNewD2_t1
BenchmarkNewD2_t1-16    	1000000000	         0.2381 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2435 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2410 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2428 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2431 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2429 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2424 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2480 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2416 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewD2_t1-16    	1000000000	         0.2445 ns/op	       0 B/op	       0 allocs/op
PASS
```

#### BenchmarkNewD2_t2
```bash
go test -bench=BenchmarkNewD2_t2 -run=^BenchmarkNewD2_t2$ -benchmem -count 10 -v
```

```text
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkNewD2_t2
BenchmarkNewD2_t2-16    	 3671695	       313.4 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 3833098	       317.9 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 3779040	       297.5 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 4007970	       302.2 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 4081206	       298.0 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 3732890	       304.8 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 3912708	       304.0 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 3745672	       302.8 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 4135250	       305.3 ns/op	      24 B/op	       3 allocs/op
BenchmarkNewD2_t2-16    	 3815491	       304.1 ns/op	      24 B/op	       3 allocs/op
PASS
```

#### BenchmarkNewD2_t3
```bash
go test -bench=BenchmarkNewD2_t3 -run=^BenchmarkNewD2_t3$ -benchmem -count 10 -v
```

```text
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkNewD2_t3
BenchmarkNewD2_t3-16    	 6194415	       185.1 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 6270892	       188.1 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 5942256	       188.6 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 6794126	       183.1 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 6247269	       186.6 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 6558057	       184.6 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 6321091	       188.0 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 6840493	       189.9 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 6105386	       187.4 ns/op	      16 B/op	       2 allocs/op
BenchmarkNewD2_t3-16    	 6536887	       195.1 ns/op	      16 B/op	       2 allocs/op
PASS
```
