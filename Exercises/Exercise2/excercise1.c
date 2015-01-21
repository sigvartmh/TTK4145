#include<pthread.h>
#include<stdio.h>

#define MAX 1000000
int i = 0;
pthread_mutex_t i_mutex;
pthread_mutexattr_t i_mattr;

void* thread_func1(){
    for(int x=0; x < MAX; x++){
        pthread_mutex_lock(&i_mutex);
        i = i + 1;
        pthread_mutex_unlock(&i_mutex);
    }
    return NULL;
}

void* thread_func2(){
    for(int x=0; x < MAX; x++){
        pthread_mutex_lock(&i_mutex);
        i = i - 1;
        pthread_mutex_unlock(&i_mutex);
    }
    return NULL;
}

int main(){

    pthread_mutexattr_init(&i_mattr);
    pthread_mutexattr_setpshared(&i_mattr, PTHREAD_PROCESS_SHARED);
    //pthread_mutexattr_settype(&i_mattr, PTHREAD_MUTEX_ADAPTIVE_NP);
    pthread_mutex_init(&i_mutex, &i_mattr);
    pthread_t thread1;
    pthread_t thread2;
    pthread_create(&thread1, NULL, thread_func1, NULL);
    pthread_create(&thread2, NULL, thread_func2, NULL);
    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);
    pthread_mutex_destroy(&i_mutex);
    printf("Hey this is i: %i\n",i);

    return 0;
}

