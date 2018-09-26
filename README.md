# Challenge Workshop Golang

Se tiene un archivo de log de transacciones, con montos muy grandes de manera que sumarlos todos causaría un overflow en el computador.

Teniendo en cuenta esta restriccion, se debe calcular, de los movimientos logueados, lo siguiente:

Para cada tipo de operación:

- El promedio de movimientos

- El usuario con mayor cantidad de movimientos

Bonus:

- [Percentil](https://es.wikipedia.org/wiki/Percentil) 95 para cada tipo de operación

## Asunciones

- Cada monto de transaccion cabe en un float64

- Se conoce el numero total de transacciones a procesar

- No hay limitaciones de memoria

## Solucion

- Para el promedio de movimientos por operacion

Se usa un algoritmo lineal **_O (n)_** iterando una sola vez la lista y guardando el promedio:

```bash
promedio += (monto_actual - promedio_acumulado) / numero_de_elementos_evaluados
```

parcial en un mapa para tener performance de busqueda constante .

- Para el usuario con mayor cantidad de movimientos por operacion

Se usa un algoritmo lineal **_O (n)_** iterando una sola vez la lista y guardando el numero de transacciones por usuario dentro de un mapa para tener performance de busqueda constante.

- Para el calculo de percentil:

El camino mas facil es ordenar los elementos y luego se puede calcular el percentil deseado.
para esto se usa QuickSort que es simple y con un rendimiento de **_O (log(n))_**