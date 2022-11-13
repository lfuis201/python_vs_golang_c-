#include <iostream>
#include <fstream>
#include <math.h>
using namespace std;

void bubbleSort(int arr[], int n)
{
    int i, j;
    for (i = 0; i < n - 1; i++)

        for (j = 0; j < n - i - 1; j++)
            if (arr[j] > arr[j + 1])
                swap(arr[j], arr[j + 1]);
}
 
void creararray(int array[], int n){
   for(int i=0;i<n;i++){
       array[i]=rand()%50000;
   }
}


int main()

{   
    ofstream file;
    file.open("/home/luisfelipe/Proyectos/ev4/tiemposc++.txt");
    ofstream file2;
    file2.open("/home/luisfelipe/Proyectos/ev4/desviacionestarc++.txt");
    int e[]={100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
    9000, 10000, 20000, 30000, 40000, 50000};
    
   for(int i=0;i<15;i++){
        double t[5];
       for(int j=0;j<5;j++){
            int lista[e[i]];
            unsigned t0, t1;
            creararray(lista, e[i]);
    
            t0=clock();
            bubbleSort(lista, e[i]);
            t1=clock();
    
            double time = (double(t1-t0)/CLOCKS_PER_SEC);
            t[j] = time;
       }
       double prom = 0;
       for (int k=0;k<5;k++){
           prom += t[k];
       }
       double promf = prom/5;

       double des = 0;
       for (int d=0;d<5;d++){
           des += pow((t[d]-promf),2);
       }
       des=sqrt(des/4);

        file<<promf<<"\n";
        file2<<des<<"\n";
       cout<<"promedio array de tamaño "<<e[i]<<": "<<promf<<" Desviacion Estandar: "<<des<<endl;

   }
    return 0;
}