# Design proposals

## Reflection vs generator in resource types marshalling/unmarshalling

First approach was to use reflection to handle multiple resource types in,
but due to Go limitations about generics whole structure needs to be copied
to avoid infinite recursion.

Thus, another option was to create generator tool that will generate
marshal/unmarshal methods for all passed resource types.

Both approaches were implemented and compared its performance using benchmark.
Below are the results which shows, that with generation mechanism application
performs quicker that with reflection.

### Reflection approach

| Name                         | Iterations | Avg run time |
| ---------------------------- | ---------- | ------------ |
| BenchmarkResourceMarshal-8   | 10000      | 123065 ns/op |
| BenchmarkResourceUnmarshal-8 | 67878      | 18211 ns/op  |

### Generators approach

| Name                         | Iterations | Avg run time |
| ---------------------------- | ---------- | ------------ |
| BenchmarkResourceMarshal-8   | 17338      | 67363 ns/op  |
| BenchmarkResourceUnmarshal-8 | 67506      | 17840 ns/op  |

### Decision

We decided to stick with resource type generation due to its better performance.
