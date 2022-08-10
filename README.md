# ReSTful API {Representational State Transfer}

```
CRUD => {Create, Read, Update, Destroy}

HTTP Methods => {GET, POST, PUT, DELETE, PATCH, OPTIONS, ...}
```

## Employee Management Server (JSON API)

```
CRUD         Action               HTTP Method             URI                   Req body                Res body
---------------------------------------------------------------------------------------------------------------------
Read         Index                GET                   /employees                -                     [{...}, ...]
Read         Show                 GET                   /employees/{id}           -                       {...}
Create       Create               POST                  /employees               {...}                   {id: , ...}
Update       Update               PUT                   /employees/{id}          {...}                    {...}
Update       Update               PATCH                 /employees/{id}          {selected attrs}         {...}
Destroy      Destroy              DELETE                /employees/{id}           -                       - / {...}
```
