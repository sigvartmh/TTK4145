# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread
from threading import Lock

i = 0
lock = Lock()
def thread_func1():
    global i
    for x in xrange(1,1000000):
        lock.acquire()
        i=i+1
        lock.release()

def thread_func2():
    global i
    for x in xrange(1,1000000):
        lock.acquire()
        i=i-1
        lock.release()

def main():
    thread1 = Thread(target = thread_func1, args = (),)
    thread2 = Thread(target = thread_func2, args = (),)
    thread1.start()
    thread2.start()

    thread2.join()
    thread1.join()
    print("This is the value of i: {0}".format(i))


main()
