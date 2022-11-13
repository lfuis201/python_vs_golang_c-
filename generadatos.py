import numpy as np

def creardatos():
    for i in range(0,5):
        
        f = open("datos"+str(i+1)+".txt", 'w')

        lista = np.random.randint(50000, size = 50000)

        for i in lista:
            f.write(str(i)+"\n")

creardatos()