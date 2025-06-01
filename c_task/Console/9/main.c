#include <stdio.h>

int main() {
    FILE *file;
    char text[] = "This is a sample text to be written to file.\n";
    char buffer[100];
    
    // Write to file
    file = fopen("sample.txt", "w");
    if(file == NULL) {
        printf("Error opening file for writing!\n");
        return 1;
    }
    fputs(text, file);
    fclose(file);
    
    // Read from file
    file = fopen("sample.txt", "r");
    if(file == NULL) {
        printf("Error opening file for reading!\n");
        return 1;
    }
    
    printf("File content:\n");
    while(fgets(buffer, sizeof(buffer), file) != NULL) {
        printf("%s", buffer);
    }
    
    fclose(file);
    return 0;
}
