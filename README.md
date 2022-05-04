# Multiple executions of validation functions

## Run

As the first step, set up aws credentials in `main.go`, then

```shell
go build .
./verifier-debug-multiple
```

## Execution

Since we create the client single time, in case of a second validation call, the output would still contain old items, so if there are cases both subnets having failures, we would need to flush the client

first:

```shell
Terminating ec2 instance with id i-0e71923d019f0f55d
subnet-030455abded968aa2 -> egressURL error: Unable to reach nosnch.in:443
```

second:

```shell
Terminating ec2 instance with id i-09d0874a9a2ee2622
subnet-03146948aeaeb21de -> egressURL error: Unable to reach nosnch.in:443
subnet-03146948aeaeb21de -> egressURL error: Unable to reach nosnch.in:443
```