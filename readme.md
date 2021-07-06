Conexion entre Farmacia 19 de Abril y Nuevos Laureles Mediante el 
ServiceQuery Nativo en PHP.

Se usara GO para establecer una conexion entre estas 

Puntos fuertes para hacer este proyecto:
- Las dos no cuentan con una IP Publica fija ni con un Internet
Estable
- Si se usara Tuneles ngrok puede tener problemas con el limit.
- En caso de usar Ngrok hay que estar al pendiente de la conexion.
- Tener la mejor estabilidad que se pueda.

Solucion: Una solucion casi perfecta seria tener un Bridge que funcione como mediador de estas.
Es decir supongamos que el sitio A enviara su inventario SQ complete Query a un servidor y el 
Sitio B tambien lo hara, el servidor hara un match de estas y lo indicara con una flag de cambio,
tanto el servidor A o el B haran una consulta periodicamente cada vez que noten un cambio de peso
de la base de datos, y con el hara una consulta al servidor mediador que pregunta si tiene una 
actualizacion pendiente, en el caso de tenerla osea tener habilitado el flag, este se actualiza-
ra con la lista de precios del servidor.

Si todo fuera mas facil(Tener IP publica fija): El que tenga Internet mas estable manda y alli 
esta el servidor oficial, cuando se cambie un precio en la 19 se cambia en nuevos laureles y 
listo sin nada mas, claro todo esto enviando una lista a el servidor de nuevos laureles para que
el cambie los precios como dice 19 de abril.

La primera solucion es la mas economica, pero cara a largo plazo y la segunda es cara al inicio y 
economica al final.

Es recomendable que estas empresas tengan una buena conexion de Interconexion entre estas. 

La aplicacion va a estar hecha en Golang con conexion a ServiceQuery Native(PHP).
Puede ser hecha en Javascript si la trama del proyecto indican el cambio.

Edit: No se usara ServiceQuery, Necesitamos estabilidad, aunque si esta al critero se puede usar el SQ.