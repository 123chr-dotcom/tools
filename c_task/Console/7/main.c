#include <stdio.h>

int main() {
    int num = 10;
    int *ptr = &num;
    
    printf("Value of num: %d\n", num);
    printf("Address of num: %p\n", &num);
    printf("Value of ptr: %p\n", ptr);
    printf("Value pointed by ptr: %d\n", *ptr);
    
    // Modify value through pointer
    *ptr = 20;
    printf("\nAfter modification through pointer:\n");
    printf("Value of num: %d\n", num);
    printf("Value pointed by ptr: %d\n", *ptr);
    
    return 0;
}
