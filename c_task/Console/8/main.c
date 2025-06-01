#include <stdio.h>
#include <string.h>

// Define student structure
struct Student {
    char name[50];
    int age;
    float gpa;
};

int main() {
    // Create and initialize a student
    struct Student student1;
    strcpy(student1.name, "Alice");
    student1.age = 20;
    student1.gpa = 3.7;
    
    // Print student information
    printf("Student Information:\n");
    printf("Name: %s\n", student1.name);
    printf("Age: %d\n", student1.age);
    printf("GPA: %.2f\n", student1.gpa);
    
    return 0;
}
