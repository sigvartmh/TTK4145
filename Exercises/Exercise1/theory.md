1. Thinking about elevators

Software crashing -> redundance? Watchdogtimers, multiple treads, if software crashes lock elevator in place and restart? Robust software/hardware

Communication failing:
Back-up plan ? Serve last order(last recived que order?) and stop? Work independently(Causes a problem when you have multiple elevators and 1 working independently)

User being trolls:

Case 1: Press all the Internal elevator buttons
If elevator got weight limits limit number of floors pressed to the "N-people" weight capacity.

Case 2: External button pressed both up and down?
Really not much you can do to prevent this. Should always serve external pressed floors in the correct que order

Case 3: Pressing stop button?
Require a 3 second push? Still if someone presses stop the elevator should stop.

2. Setup source controll and build tools

Decided we would go for GIT? Kurt if you have any objections to this decision then you should raise your voice now and yell at me ^-^

Unsure about build tools(very languge dependant) Would prefer something like writing in Go using a regualar text editor and go compiler in terminal. Else we could be doing things a bit more a complicated and use the intelliJ  IDEA golang.org plugin
http://plugins.jetbrains.com/plugin/5047

It's available for students for free through jetbrains website just registrer with sutdent account. It does have some nice features but takes up some computer resources so it's all about that pros and cons.

3. Why concurrency?

Why do we use concurrent execution (multithreading/multiprocessing, or the like)?

* To do more at the same time.(Better preformance) on a multicore system/CPU with the possibility to run multiple threads.
* Robustness(If one thread/process dies it can be restarte by another)
* Correctenss(???) -> When components that operate concurrently interact by messaging or by sharing accessed data (in memory or storage), a certain component's consistency may be violated by another component. The general area of concurrency control provides rules, methods, design methodologies, and theories to maintain the consistency of components operating concurrently while interacting, and thus the consistency and correctness of the whole system.
*Remain responsive to input by having a thread taking input while another is working?

It can make programming simpler by reducing the complexity of the program. However programming can become more difficult by introducing other problems such as deadlocks and raceconditions.

Difference between processes and threads are that typicaly threads runs of the same process using a shared memory spaces while processes runs in a seperate memoryspaces.
-> Multiple threads can exist within the same process and share resources such as memory, while different processes do not share these resources. In particular, the threads of a process share its instructions (the code) and its context (the values of its variables at any given moment).



Process are truly concurrent, at least in the presence of suitable hardware support. Exist within their own address space. Can be managed by and OS.

Threads is the smallest sequence of programmed instructions that can be managed independently by a scheduler(Managed by )

Green threads are threads that are scheduled by a virtual machine (VM) instead of natively by the underlying operating system.

Coroutines are computer program components that generalize subroutines for nonpreemptive multitasking, by allowing multiple entry points for suspending and resuming execution at certain locations.

Which one of :...
The pthread_create() function is used to create a new thread same goes for threading.Thread() (Python) and go (Go)

The python Global Interpeter Lock(GIL) influences the way a python Thread behaves in the CPython interpreter to assure that only one thread executes Python bytecode at a time.

Work around:
multiprocessing module.
You would have to share resources variables/data manually


What does func GOMAXPROCS(n int) int change?
	The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code simultaneously.

Excercise 4:

What happens is that i is returned with a different result and not the value 0 as which you would expect. This is because 1 thread finishes before the other.(I think?) -> Racecondition.


