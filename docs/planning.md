
## Failure handling:

### Retry
In this approach, if the callback function fails, the job queue 
will retry  the callback again after a certain amount of time. 
The number of retries and the time between retries are configurable.

### Dead Letter Queue
In this approach, if the callback function fails, the job is added 
to a dead letter queue (DLQ) for further examination. The DLQ is a
separate queue that holds the failed jobs. These jobs can then be
analyzed to understand why they failed and can be resubmitted to 
the job queue after the problem is fixed.

### Discard
In this approach, if the callback function fails, the job is simply
discarded and not retried. This approach is usually used when the 
job can be easily regenerated or is not critical.

### Alert
In this approach, if the callback function fails, an alert is sent 
to the administrator or the developer. This approach is useful when
the job is critical and requires immediate attention.


## Shutdown handling

### Drain
In this approach, the bucket will stop accepting new jobs, but it
will continue to process the remaining jobs in the queue. Once all
the jobs are processed, the application will shut down.

### Flush
In this approach, the bucket will stop accepting new jobs, and it will
attempt to flush any remaining jobs in the queue. The application will
wait for the flush to complete before shutting down.

### Stop
In this approach, the bucket will stop processing new and remaining jobs
immediately, and it will shut down. Any remaining jobs will be lost.

### Persist
In this approach, the bucket will stop accepting new jobs, will flush 
any remaining jobs and will persist the state of the queue. So the next
time the application starts, it continues to process the persisted jobs.
