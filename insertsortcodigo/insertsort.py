import numpy as np
from timeit import default_timer
import math

def insertionSort(arr):
 
    # Traverse through 1 to len(arr)
    for i in range(1, len(arr)):
 
        key = arr[i]
 
        # Move elements of arr[0..i-1], that are
        # greater than key, to one position ahead
        # of their current position
        j = i-1
        while j >= 0 and key < arr[j] :
                arr[j + 1] = arr[j]
                j -= 1
        arr[j + 1] = key

def crearlistas(x):
    listaprincipal = []
    listadatos = []

    with open(x) as archivo:
        for linea in archivo:
            listaprincipal.append(int(linea))
        
    for i in e:
        temp = listaprincipal[0:i]
        listadatos.append(temp)

    return listadatos


f = open("tiemposejecucioninsert/tiempospythoninsertsort.txt", 'w')
f1 = open("desviacionestandarinsert/desviacionstandarpythoninsertsort.txt", 'w')
e = [100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
9000, 10000, 20000, 30000, 40000, 50000]

archivo = "datos"

for i in range(0,4):
    t=[]
    
    listprin = []
    for x in range(0,15):
        archivoindex = archivo + str(x+1) + ".txt"
        listas = crearlistas(archivoindex)
        listprin.append(listas)
    
    for j in range(0,5):
        a = listprin[j][i]

        inicio = default_timer()
        insertionSort(a)

        fin = default_timer()
        delta = fin - inicio
        t.append(delta)
        prom = 0
    for k in range(0,5):
        prom += t[k]
    promf = prom/5

    des = 0
    for d in range(0,5):
        des += pow(t[d]-promf, 2)

    des=math.sqrt(des/4);

    print("tiempo promedio de array de tamaño ", e[i]," ",promf," Desviacion Estandar: ", des)
    f.write(str(promf)+"\n")
    
    f1.write(str(des)+"\n")
f.close()
f1.close()