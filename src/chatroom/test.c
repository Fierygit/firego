
#include <stdio.h>

union data{
    int i ;
    char c;
    float f;
}a;
int n;


int main()
{

    // a = 5;
    // a = {2,' a', 1.2};
   // n = a;
    char * a = "dfsdf";
     *a = "fdf";
    printf(a);

}