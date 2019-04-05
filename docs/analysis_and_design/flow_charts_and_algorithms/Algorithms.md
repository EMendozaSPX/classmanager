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

### Access Persistent Constants
```
BEGIN AccesssPersistentConstants(constant_request)
    IF environment.json NOT exists
        Create environment.json
        CreateEnvironmentConstants()
    END IF
    
    let env_constants = get constants from environment.json
    return env_constants[constant_request]
END AccessPersistentConstants
```

### Create Environment Constants
```
BEGIN CreateEnvironmentConstants
    let SECRET_KEY = generate random string
    store SECRET_KEY in environment.json
END CreateEnvironmentConstants
```

## Authentication and Authorization

* note `bcrypt` is a golang package that contains bcrypt salt and hashing features
* note global_authenticated is a global variable

### Login Form
```
BEGIN LoginForm
    let global_login_state = false
    submit_button.event('click', Login)
END LoginForm
```

### Login Callback
```
BEGIN Login
    let account_type = student | teacher | admin
    let username, password = get_username_state('username')
    let password = get_password_state('password')
    
    POST username, password, account_type TO server
    
    IF response bool == true
        print(response msg)
        let user_access_token = response token
        REDIRECT TO /<usertype:string>/dashboard
    
    ELSE
        print(response msg)
        
    END IF
END Login
```

### Login Server Side
```
BEGIN Login
    let username, password, account_type = LISTEN FOR username, password
    
    IF username NOT IN table user_type
        RESPOND (false, msg: User not found, empty token)
        END Login
    END IF
    
    let password_hash = account_type[username_cypher]
    
    IF CompareHashAndPassword(password_hash, password)
        let user_email = get username from table user_type
        json_web_token = jwt(sectret_key, [ account_type, username, user_email )
        RESPOND (true, msg: Login Success, json_web_token)
        END Login
    END IF
    
    RESPOND (false, msg: Login failed, empty token)
END Login
```

### Access protected url
```
BEGIN RequestProtectedRoute
    POST json_web_token, requested url TO server
    IF RESPONSE == true
        access granted
    ELSE
        redirect to login
    END IF
END RequestProtectedRoute
```

### Access protected url server side
```
BEGIN RequestProtectedRoute
    let json_web_token, requested url = LISTEN FOR json_web_token, requested url
    
    IF VerifyToken(json_web_token)
        RESPOND true
    ELSE
        RESPOND false
    END IF
END RequestProtectedRoute
```

### Verify user with json web token
```
BEGIN VerifyToken(json_web_token, permitted_usertype)
    let secret_key, usertype, username = decode json_web_token
    IF secret_key == SECRET_KEY AND usertype == permitted_usertype
        let username = username
        IF username IN database table user_type
            RETURN true
        END IF
    END IF
    RETURN false
END VerifyToken
```

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
    IF NOT check
        return true
    END IF
    return false
END CompareHashAndPassword
```

## Admin

### View List of Users
* Note: view list of users does not display passwords because passwords are not stored in database
```
BEGIN ViewListOfUsers
    GET list_of_users WITH user_type, json_web_token FROM server
    
    let success, users = server response
    IF success
        COUNT EACH user IN list_of_users
            display user
        END COUNT EACH
    ELSE
        display(access denied)
    END IF
    
    IF user.event(click)
        REDIRECT TO /admin/<user_type:string>/<username:string>
    END IF
END ViewListOfUsers
```

### View List of Users server side
```
BEGIN GetUsersList
    let user_type, json_web_token = LISTEN FOR list_of_users get request
    
    IF VerifyToken(json_web_token, admin)
        let users = get users from user_type table in database
    
        RESPOND (true, users)

    ELSE
        RESPOND (false, null)
    END IF
END GetUsersList
```

### Create a user server function
```
BEGIN CreateUser(account_type, username, email, password)
    let password_hash = SaltAndHash(password)
    
    let error = STORE username, email, password_hash in database table account_type
    
    IF lengthof(error) >= 0
        IF error == username not unique
            RETURN (false, username not unique error)
        ELSE IF error == email not unique
            RETURN (false, email not unique)
        ELSE
            RETURN (false, error)
        END IF
    END IF
    RETURN true
END CreateUser
```

### Admin User Information Panel
```
BEGIN AdminUserPanel(admin_user)
    let username_field = display admin_user.username
    let email_field = display admin_user.email
    let password_field = empty string
    
    IF username_field changed
        username_field state = new username_field state
    ELSE IF email_field changed
        email_field state = new email_field state
    END IF
    
    IF submit.event(click)
        POST (json_web_token, user.id, username_field, email_field, password_field) TO server
    END IF
END AdminUserPanel
```

### Update Admin User server side
```
BEGIN UpdateAdminUser
    json_web_token, user_id, username, email, password = LISTEN FOR user.id, username_field, email_field, password_field
    IF VerifyToken(json_web_token) AND user_id IN database table user_type
        IF username NOT == user.username
            user.username = username
        END IF
        
        IF email NOT == user.email
            user.email = email
        END IF
        
        user.password = SaltAndHash(password)
    END IF
```

### Create an Admin User
```
BEGIN CreateAdminUser
    password_field
    IF submit_button.event('click') AND all_fields NOT == null
         POST (json_web_token, username, email, password) TO server
         RECEIVE is_success, error
            
         IF NOT is_success
             print(error)
         ELSE
             REDIRECT TO /admin/admins/
    ELSE
         highlight null fields
    END IF
END CreateAdminUser
```

### Create an Admin User server side
```
BEGIN CreateAdminUser
    let json_web_token, username, email, password = LISTEN FOR json_web_token, username, email, password
    
    IF VerifyToken(json_web_token, admin)
        let success, error = CreateUser(admin, username, email, password)
        
        IF success
            RESPOND success
        ELSE
            RESPOND (success, error)
        END IF
        
    ELSE
        RESPOND false, user verification failed
    
    END IF
END CreateAdminUser
```

### Create a Class
```
BEGIN CreateClass
    let class
```