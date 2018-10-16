#include <stdio.h>


void test1()
{
	int a=2;
	int *pa = &a;
	*pa=123;
	printf("%d\n",*pa);

	int b=456;
	*pa = b;
	printf("%d\n",*pa);
}

void swap(int *a, int*b)
{
	int tmp;
	tmp =*a;
	*a = *b;
	*b=tmp;
}

void test2()
{
	int a=123,b=456;
	swap(&a, &b);
	printf("a:%d, b:%d\n",a,b);
}


int main(int argc, const char *argv[])
{
	test1();
	printf("================\n");
	test2();
	return 0;
}
