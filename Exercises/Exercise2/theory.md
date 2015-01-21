Exercise 2 : Bottlenecks
========================

Mutex and Channel basics
------------------------

 - An atomic operation?
Kurt: An operation that appears to happen instantanously. That means that it can not be devided; it will either happen or not happen. This is used to ensure that other changes can not take place while the operation is carried out, and ot ensure data integrity in e.g. databases.

 - A semaphore?
Kurt: A system for reserving resources and handing them out in an orderly manner. It can be a simple counter saying how many resources are available, and increment or decrement as they are handed out and returned - or it can be a simple binary reservation for a single resource. The latter may act as a mutex.

 - A mutex?
Kurt: Short for "Mutual exclusion". A mutex locks a shared resource to a specific process, so that it can carry out its critical operation.

 - A critical section?
Kurt: An operation in a computer program that does something to a shared resource. During the critical section, other threads must be prevented from making changes to the same resource. This can be handled by letting all other operations wait, or with semaphores that lock only the affected resource.


