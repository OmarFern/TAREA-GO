
a=["malayalam"," ","",-1]

    
def validar_2(a):  
    for i in range (len(a)):
        if ( a[i] ==invertir(a[i])):
            resultado= True
        else:
            resultado= False
        return resultado

def invertir(a):
    lista=[]
    for i in range(len(a)):
        lista.append(a[i])
    resultado="".join(lista)   
    return resultado  
print(validar_2(a))      


    