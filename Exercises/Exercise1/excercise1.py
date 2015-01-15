# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread

i = 0

def thread_func1():
    global i
    for x in xrange(1,1000000):
        i=i+1

def thread_func2():
    global i
    for x in xrange(1,1000000):
        i=i-1

def main():
    thread1 = Thread(target = thread_func1, args = (),)
    thread2 = Thread(target = thread_func2, args = (),)
    thread1.start()
    thread2.start()

    thread2.join()
    thread1.join()
    print("This is the value of i: {0}".format(i))


main()
