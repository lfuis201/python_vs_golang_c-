#include <iostream>
#include <fstream>
#include <math.h>
#include <vector>
#include <bits/stdc++.h>

using namespace std;

void insertionSort(vector<int> &vec)
{
    for (auto it = vec.begin(); it != vec.end(); it++)
    {        
        // Searching the upper bound, i.e., first 
        // element greater than *it from beginning
        auto const insertion_point = 
            upper_bound(vec.begin(), it, *it);
          
        // Shifting the unsorted part
        rotate(insertion_point, it, it+1);        
    }
}
  
vector<int> slicing(vector<int> arr, int min, int max){
    auto start = arr.begin()+min;
    auto end = arr.begin()+max+1;

    vector<int> result(max-min+1);

    copy(start, end, result.begin());

    return result;
}

vector<vector<int>> generarlistas(string file)
{
    int number;
    vector<int> nums{};
    vector<vector<int>> listadatos{};
    vector<int> temp{};
    int e[]={100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
    9000, 10000, 20000, 30000, 40000, 50000};

    ifstream input_file(file);

    while (input_file >> number) {
        nums.push_back(number);
    }

    for (int i = 0; i < 15; i++)
    {   
        temp = slicing(nums,0,e[i]-1);
        listadatos.push_back(temp);
    }
    
    return listadatos;

}
 
int main()

{   
    ofstream file;
    file.open("/home/luisfelipe/Proyectos/ev4/tiemposejecucioninsert/tiemposc++insert.txt");
    ofstream file2;
    file2.open("/home/luisfelipe/Proyectos/ev4/desviacionestandarinsert/desviacionestarc++insert.txt");
    int e[]={100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
    9000, 10000, 20000, 30000, 40000, 50000};
    
    vector<vector<int>> listas{};
    string datos = "datos";
    
   for(int i=0;i<15;i++){
        double t[5];
        for(int j=0;j<5;j++){
            
            listas = generarlistas("/home/luisfelipe/Proyectos/ev4/"+datos+to_string(j+1)+".txt");
            unsigned t0, t1;
            
    
            t0=clock();
            insertionSort(listas[i]);
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

   file.close();
   file2.close();
    return 0;
    
}