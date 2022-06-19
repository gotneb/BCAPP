# BCAPP
 A BCAPP's library written in Go.
 ## API
 ### Endpoints avaliable on ECC
 
<details><summary>modp</summary><p>
 
It has the same parameter meaning from Ecc.modp(BigInteger n, BigInteger p1). [See java file.](https://github.com/carlosamcruz/BCAPP/blob/master/app/src/main/java/com/nibblelinx/BCAPP/Ecc.java#L47)
 
#### HTTP Request

```
 GET https://bcapp-go.herokuapp.com/modp/<n>/<p1>
```
 #### URL Parameters

| Parameter  |  Description  |
| --- | --- |
|  n |  Big Integer |
|  p1 |  Big Integer |

</p>

</details>

<details><summary>inverse</summary><p>
 
It has the same parameter meaning from Ecc.inverse(BigInteger r, BigInteger p). [See java file.](https://github.com/carlosamcruz/BCAPP/blob/master/app/src/main/java/com/nibblelinx/BCAPP/Ecc.java#L54)
 
#### HTTP Request

```
 GET https://bcapp-go.herokuapp.com/inverse/<r>/<p>
```
 #### URL Parameters

| Parameter  |  Description  |
| --- | --- |
|  r |  Big Integer |
|  p |  Big Integer |

</p>

</details>

<details><summary>doublep</summary><p>
 
It has the same parameter meaning from Ecc.doublep(x, y). [See java file.](https://github.com/carlosamcruz/BCAPP/blob/master/app/src/main/java/com/nibblelinx/BCAPP/Ecc.java#L81)
 
#### HTTP Request

```
 GET https://bcapp-go.herokuapp.com/doublep/<x>/<y>
```
 #### URL Parameters

| Parameter  |  Description  |
| --- | --- |
|  x |  Big Integer |
|  y |  Big Integer |

</p>

</details>
