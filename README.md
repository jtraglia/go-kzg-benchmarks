# go-kzg-benchmarks

Run the benchmarks with this command:

```
go test -bench=Benchmark
```

Example output:

```
goos: darwin
goarch: arm64
pkg: go-kzg-benchmarks
Benchmark/BlobToKZGCommitment-8         	      16	  69280844 ns/op
Benchmark/VerifyKZGProof-8              	     640	   1850233 ns/op
Benchmark/ComputeAggregateKZGProof(blobs=1)-8         	       7	 157965893 ns/op
Benchmark/ComputeAggregateKZGProof(blobs=2)-8         	       5	 228077000 ns/op
Benchmark/ComputeAggregateKZGProof(blobs=4)-8         	       3	 368934195 ns/op
Benchmark/ComputeAggregateKZGProof(blobs=8)-8         	       2	 647058438 ns/op
Benchmark/ComputeAggregateKZGProof(blobs=16)-8        	       1	1205604334 ns/op
Benchmark/ComputeAggregateKZGProof(blobs=32)-8        	       1	2332204834 ns/op
Benchmark/ComputeAggregateKZGProof(blobs=64)-8        	       1	4580128916 ns/op
Benchmark/VerifyAggregateKZGProof(blobs=1)-8          	     328	   3572716 ns/op
Benchmark/VerifyAggregateKZGProof(blobs=2)-8          	     258	   4659953 ns/op
Benchmark/VerifyAggregateKZGProof(blobs=4)-8          	     183	   6515187 ns/op
Benchmark/VerifyAggregateKZGProof(blobs=8)-8          	     100	  10091367 ns/op
Benchmark/VerifyAggregateKZGProof(blobs=16)-8         	      69	  17230699 ns/op
Benchmark/VerifyAggregateKZGProof(blobs=32)-8         	      37	  32111082 ns/op
Benchmark/VerifyAggregateKZGProof(blobs=64)-8         	      19	  58874700 ns/op
```