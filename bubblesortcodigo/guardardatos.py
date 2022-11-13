listassize = [100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
9000, 10000, 20000, 30000, 40000, 50000]

def crearlistas(x):
    listaprincipal = []
    listadatos = []

    with open(x) as archivo:
        for linea in archivo:
            listaprincipal.append(int(linea))
        
    for i in listassize:
        temp = listaprincipal[0:i]
        listadatos.append(temp)

    return listadatos

archivo = "datos"
listprin = []

for x in range(0,5):
        archivoindex = archivo + str(x+1) + ".txt"
        listas = crearlistas(archivoindex)
        listprin.append(listas)
        
print(len(listprin[2][14]))