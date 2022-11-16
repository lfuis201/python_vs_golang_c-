#include <iostream>
#include <fstream>
#include <string>
#include <math.h>
#include <vector>

using namespace std;

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
    string filename("/home/luisfelipe/Proyectos/ev4/datos1.txt");
    
    vector<vector<int>> datos = generarlistas(filename);

    cout<<datos[0].size()<<endl;
    cout<<datos[0][99]<<endl;
}