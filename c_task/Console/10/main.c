#include <stdio.h>

// Recursive function to calculate Fibonacci number
int fibonacci(int n) {
    if(n <= 1) {
        return n;
    }
    return fibonacci(n-1) + fibonacci(n-2);
}

int main() {
    int n;
    printf("Enter a positive integer: ");
    scanf("%d", &n);
    
    printf("Fibonacci(%d) = %d\n", n, fibonacci(n));
    return 0;
}
