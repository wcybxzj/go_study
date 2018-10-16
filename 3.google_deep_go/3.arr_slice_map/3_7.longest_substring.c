#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define NUM 1000

//判断字符是否是不重复的字符串
//1 is ok
//0 is invaild
int string_is_invalid(char *str, int len)
{
	int i, j;
	int current_index=0;
	if (len==0) {
		return 0;
	}
	for (i = 0; i <len; i++) {
		char current_char = str[i];
		for (j = 0; j < len; j++) {
			if (i == j) {
				continue;
			}
			if (current_char == str[j]) {
				return 0;
			}
		}
	}
	return 1;
}

void test_string_is_invalid()
{
	int ret;
	char *str ="";
	ret = string_is_invalid(str, strlen(str));
	printf("%d\n", ret);//0

	str ="abcabcbb";
	ret = string_is_invalid(str, strlen(str));
	printf("%d\n", ret);//0

	str="abcdef";
	ret = string_is_invalid(str, strlen(str));
	printf("%d\n", ret);//1
}

//获取最长不重复字符串
char *longest_substing(char *origin_str, int len)
{
	int i,ret;
	int final_index=0;
	int tmp_index=0;
	char *final_str = calloc(NUM,sizeof(char));
	char *tmp_str = calloc(NUM,sizeof(char));
	if (final_str ==NULL || tmp_str==NULL) {
		printf("malloc error\n");
		exit(1);
	}
	for (i = 0; i < len; i++) {
		tmp_str[tmp_index] = origin_str[i];
		tmp_index++;
		ret = string_is_invalid(tmp_str, tmp_index);
		if (ret) {//格式ok
			if (tmp_index > final_index) {
				strcpy(final_str, tmp_str);
				final_index=tmp_index;
			}
		}else{//格式错误
			memset(tmp_str, '\0', strlen(tmp_str));
			tmp_index=0;
			i--;
		}
	}
	free(tmp_str);
	if (final_index) {
		return final_str;
	}else{
		return NULL;
	}
}

void test_longest_substing(){
	char *str="abcabcbbabc123bbb";
	char *ret = longest_substing(str,strlen(str));
	printf("ret:%s\n",ret);//bc123

	str="abcabcbbbbb";
	ret = longest_substing(str,strlen(str));
	printf("ret:%s\n",ret);//abc

}


//例:寻找最长的不含有重复字符的字串
//abcabcbb --> abc
//bbbbbbb  --> b
//pwwkew   -->wke
int main(int argc, const char *argv[])
{
	//test_string_is_invalid();
	test_longest_substing();
	return 0;
}
