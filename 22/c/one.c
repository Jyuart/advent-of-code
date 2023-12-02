#include <stdio.h>
#include <stdlib.h>

int main() {
	FILE *file;
	file = fopen("./input_one", "r");

	if (file == NULL) {
		printf("Error");
		exit(1);
	}

	char text[10];
	fgets(text, 10, file);
	printf("string is: %s\n", text);

	return 0;
}
