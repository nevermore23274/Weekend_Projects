#include <stdio.h>
int main(void) {
	int numbers[5];
	int i, aux;
	int swapped;

	/* ask the user to input 5 values */
	for(i = 0; i < 5; i++) {
		printf("\nEnter value #%i\n",i + 1);
		scanf("%d",&numbers[i]);
	}

	/* sort them */
	do {
		swapped = 0;
		for(i = 0; i < 4; i++) {
			if(numbers[i] > numbers[i + 1]) {
				swapped = 1;
				aux = numbers[i];
				numbers[i] = numbers[i + 1];
				numbers[i + 1] = aux;
			}
		}
	} while(swapped);

	/* print results */
	printf("\nSorted array: ");
	for(i = 0; i < 5; i++)
		printf("%d ",numbers[i]);

	printf("\n");
	return 0;
}
