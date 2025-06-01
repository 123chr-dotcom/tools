#include <stdio.h>

int main() {
    int numbers[] = {2, 4, 6, 8, 10};
    int sum = 0;
    int size = sizeof(numbers) / sizeof(numbers[0]);
    
    printf("Array elements: ");
    for(int i = 0; i < size; i++) {
        printf("%d ", numbers[i]);
        sum += numbers[i];
    }
    
    printf("\nSum of array elements: %d\n", sum);
    return 0;
}
