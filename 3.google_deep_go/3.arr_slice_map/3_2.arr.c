#include <stdio.h>

void print(int arr[], int num)
{
	int i;
	for (i = 0; i < num; i++) {
		printf("%d\n",arr[i]);
	}
}

void test(int arr[],int num)
{
	arr[0]=111111111;
	print(arr,num);
}

////证明:C语言的数组做参数是引用传递
//./3_2.arr
//11
//22
//33
//44
//55
//111111111
//22
//33
//44
//55
//111111111
//22
//33
//44
//55
int main(int argc, const char *argv[])
{
	int arr[5]={11,22,33,44,55};

	print(arr,5);
	test(arr,5);
	print(arr,5);

	return 0;
}
