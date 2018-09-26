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
