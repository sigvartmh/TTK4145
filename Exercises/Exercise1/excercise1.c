#include<pthread.h>
#include<stdio.h>

#define MAX 1000000
int i = 0;

void* thread_func1(){
    for(int x=0; x < MAX; x++){
        i++;
    }
    return NULL;
}

void* thread_func2(){
    for(int x=0; x < MAX; x++){
        i--;
    }
    return NULL;
}

int main(){

    pthread_t thread1;
    pthread_t thread2;
    pthread_create(&thread1, NULL, thread_func1, NULL);
    pthread_create(&thread2, NULL, thread_func2, NULL);
    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);
    printf("Hey this is i: %i",i);

    return 0;
}

