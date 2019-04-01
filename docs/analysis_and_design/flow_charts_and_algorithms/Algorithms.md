# Algorithms

## General

* note: `http` is a standard library package provided with go to handle http requests and responses
* note: `/` and `/graphql` are website routes and information about graphql queries can be found in requirements report

### Routes
```
BEGIN StartWebserver
    let website = path_to_website
    http.handle("/", website)
    
    let graphql_handler = graphql_queries
    http.handle("/graphql", graphql_handler)
    
    http.ListenAndServe()
END StartWebserver
```

### Login
```
BEGIN Login
    let account_type = student | teacher | admin
    let username = input('username')
    let password = input('password')
    
    POST username password account_type TO server
    
    return GET response as bool
END Login
```

### Login Server Side
* Uses subprograms found in this document
```
BEGIN Login
    let username, password, account_type = listen for username and password
    
    let account_hash = account_type[username]
    
    http response CompareHashAndPassword(account_hash, password)
END Login
```

## Authentication

* note `bcrypt` is a golang package that contains bcrypt salt and hashing features

### Salt and Hash
```
BEGIN SaltAndHash(password)
    let hash = bcrypt.GenerateFromPassword(password, bcrypt.MinCost) as string
    return hash
END SaltAndHash
```

### Compare Hashes
```
BEGIN CompareHashAndPassword(hash, password)
    let check = bcrypt.CompareHashAndPassword(hash, password) as bool
    IF !check
        return true
    END IF
    return false
END CompareHashAndPassword
```

### 

## Admin

## Create A New User

### 
