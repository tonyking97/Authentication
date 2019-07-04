#include<stdio.h>

int main(){
    int num=897,a=0;
    while (num){
        if(num & 1) a++;
        num = num >> 1;
    }
    printf("%d",a);
    return 0;
}

#include<stdio.h>
#include<conio.h>

void main(){
    int x;
    printf("Enter the number : ");
    scanf("%d",&x);
    if(x <= 3){
        if(x==1){
            printf("X is 1\n");
        }else{
            printf("X is not 1\n");
        }
        if(x==2){
            printf("X is 2\n");
        }else{
            printf("X is not 2\n");
        }
        if(x==3){
            printf("X is 3\n");
        }else{
            printf("X is not 3\n");
        }
    }else{
        printf("X is greeter than 3");
    }
    getch();
}