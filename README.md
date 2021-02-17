# redis-pubsub

An example pubsub with redis.

## Requirements
Docker Compose is used to define and run the services.

[Docker installation](https://docs.docker.com/engine/install/)

## Usage

Run:
```
make
```

Clean up:
```
make clean
```

## What does this do?
This project consists of a publisher `pub` and a subscriber `sub`. The publisher
will start 3 workers to process 20 jobs. Each job will attempt to provision a
mock resource (like a compute instance). The resource provision will randomly fail.
In the event of a failure, the publisher will publish a message via a redis
pub/sub channel. The subscriber subscribes to the provision failure alert channel
and logs the ID of the failed job. Finally, the publisher service caches the
successfully provisioned resource IDs in redis for later processing.

Example success snippet:
```
pub_1    | 2021/02/12 04:01:42 worker 1 started job 11
pub_1    | 2021/02/12 04:01:43 worker 1 finished job 11
pub_1    | 2021/02/12 04:01:43 {ID:11 Provisioned:true}
pub_1    | 2021/02/12 04:01:43 add provisioned resource with ID 11 to cache
```

Example failure snippet:
```
pub_1    | 2021/02/12 04:01:44 worker 1 started job 17
pub_1    | 2021/02/12 04:01:45 worker 1 finished job 17
pub_1    | 2021/02/12 04:01:45 {ID:17 Provisioned:false}
pub_1    | 2021/02/12 04:01:45 failure for 17, send alert
sub_1    | 2021/02/12 04:01:45 ==> receiving msg chan_provision_failure payload-17
```

### Shortcomings
- many hardcoded values that could be CLI args or env variables
- lacking unit tests
  - would require refactoring to simplify unit testing
- integration tests
- need to gracefully handle panics

### More considerations
- How might this project be scaled?
  - make the worker and job counts tuneable
  - provision multiple publishers
- How might one approach doing sequential versus parallel tasks?
  - the goroutine in this code performs some of the tasks concurrently
  - in order to do all tasks sequentially--remove the goroutine
