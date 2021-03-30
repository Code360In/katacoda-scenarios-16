#### Log to stdout and to file on same time






```
cat << 'EOF' > test.py
import logging

# set up logging to file - see previous section for more details
logging.basicConfig(level=logging.DEBUG,
                    format='%(asctime)s %(name)-12s %(levelname)-8s %(message)s',
                    datefmt='%m-%d %H:%M',
                    filename='/var/log/myapp.log',
                    filemode='w')
# define a Handler which writes INFO messages or higher to the sys.stderr
console = logging.StreamHandler()
console.setLevel(logging.INFO)
# set a format which is simpler for console use
formatter = logging.Formatter('%(name)-12s: %(levelname)-8s %(message)s')
# tell the handler to use this format
console.setFormatter(formatter)
# add the handler to the root logger
logging.getLogger().addHandler(console)

# Now, we can log to the root logger, or any other logger. First the root...
logging.info('Jackdaws love my big sphinx of quartz.')

# Now, define a couple of other loggers which might represent areas in your
# application:

logger1 = logging.getLogger('myapp.area1')
logger2 = logging.getLogger('myapp.area2')

logger1.debug('Quick zephyrs blow, vexing daft Jim.')
logger1.info('How quickly daft jumping zebras vex.')
logger2.warning('Jail zesty vixen who grabbed pay from quack.')
logger2.error('The five boxing wizards jump quickly.')

EOF

```{{execute}}



```
python test.py 
```{{execute}}


```
more /var/log/myapp.log
```{{execute}}


https://kubernetes.io/docs/concepts/cluster-administration/logging/

https://docs.docker.com/config/containers/logging/configure/


# Reliable logging for a disk-intensive application
<pre class="file">

If your application emits a large amount of logs and performs high I/O workloads, consider using the json-file logging driver in non-blocking mode. This configuration maintains the performance of your container’s applications while still giving you reliable structured logging. Writing logs to local storage is quick, so it’s unlikely that the ring buffer will fill up. If your application doesn’t need to support spikes in logging activity (which could fill up the ring buffer), this configuration can capture all of your logs without interrupting your application.


</pre>




# A RAM-intensive application with an external logging service
<pre class="file">


If you can’t create logs locally, you’ll need to use a non-default driver like gcplogs or awslogs to log to a remote destination. For very memory-hungry applications that require the majority of the RAM available to your containers, you should generally use blocking mode. This is because there may not be enough memory available for the ring buffer if the driver cannot deliver logs to its endpoint, for example because of a network problem.

Although using blocking mode will ensure that you capture every log, we generally don’t recommend using a non-default driver in blocking mode because your application’s performance could deteriorate if logging is interrupted. For mission-critical applications that can’t log to the local filesystem, but where performance is a higher priority than logging reliability, provision enough RAM for a reliable ring buffer and use non-blocking mode. This configuration ensures that your application’s performance does not suffer from a problem with logging, while also providing room to capture the majority of log data.


</pre>




# Optimize logging reliability and application performance
<pre class="file">
The Docker logging driver and log delivery mode you choose can have a noticeable effect on the performance of your containerized applications. We recommend using the json-file driver for reliable logging, consistent performance

</pre>




# Delivery modes

<pre class="file">

Regardless of which logging driver you choose, you can configure your container’s logging in either blocking or non-blocking delivery mode. The mode you choose determines how the container prioritizes logging operations relative to its other tasks.

Blocking
A container in blocking mode—Docker’s default mode—will interrupt the application each time it needs to deliver a message to the driver. This guarantees that all messages will be sent to the driver, but it can have the side effect of introducing latency in the performance of your application: if the logging driver is busy, the container delays the application’s other tasks until it has delivered the message.

The potential effect of blocking mode on the performance of your application depends on which logging driver you choose. For example, the json-file driver writes logs very quickly since it writes to the local filesystem, so it’s unlikely to block and cause latency. On the other hand, drivers that need to open a connection to a remote server—such as gcplogs and awslogs—are more likely to block for longer periods and could cause noticeable latency.

Non-blocking
In non-blocking mode, a container first writes its logs to an in-memory ring buffer, where they’re stored until the logging driver is available to process them. Even if the driver is busy, the container can immediately hand off application output to the ring buffer and resume executing the application. This ensures that a high volume of logging activity won’t affect the performance of the application running in the container.

Non-blocking mode does not guarantee that the logging driver will log all the events, in contrast to blocking mode. If the application emits messages faster than the driver can process them, the ring buffer can run out of space. If this happens, buffered logs will be deleted before they can be handed off to the logging driver. You can use the max-buffer-size option to set the amount of RAM used by the ring buffer. The default value for max-buffer-size is 1 MB, but if you have RAM available, increasing the buffer size can increase the reliability of your container’s logging.

Although blocking mode is Docker’s default for new containers, you can set this to non-blocking mode by adding a log-opts item to Docker’s daemon.json file. The example below extends the previous example, keeping fluentd as the default driver and changing the default mode to non-blocking:
</pre>
