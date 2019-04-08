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
    display input username_field
    display input password_field
    display input confirm_password_field
    
    submit_button.event('click', Login)
END LoginForm
```

### Login Callback
```
BEGIN Login
    let account_type = student | teacher | admin
    let username, password = get_username_state('username_field')
    let password = get_password_state('password_field')
    let confirm_password = get_password_state('confirm_password_field')
    
    IF password == confirm_password
        POST username, password, account_type TO server
    ELSE
        REDIRECT TO /<usertype:string>/login
        print(passwords dont match)
    END IF
    
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
        json_web_token = jwt(AccessPersistantConstants(SECRET_KEY), [ account_type, username, user_email )
        RESPOND (true, msg: Login Success, json_web_token)
        END Login
    END IF
    
    RESPOND (false, msg: Login failed, empty token)
END Login
```

### Logout
```
BEGIN Logout
    local.user_access_token = ''
    Redirect to /
END Logout
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
    IF secret_key == AccessPersistantConstant(SECRET_KEY) AND usertype == permitted_usertype
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

### Set School Bounds
```
BEGIN SetBounds
    display select school_week = weekly | bi-weekly
    
    display select period = begin | recess | lunch | end | 1 | 2 | 3 .... 16
    display input start_time
    display input end_time
    
    IF add.button(click)
        display select period = begin | recess | lunch | end | 1 | 2 | 3 | .... 16
        display input start_time = previous end_time
        display input end_time
    END IF
    
    display input term_one, start_time, end_time
    display input term_two, start_time, end_time
    display input term_three, start_time, end_time
    display input term_four, start_time, end_time
    
    IF submit.event(click)
        POST all fields TO server
    END IF
END SetBounds
```

### Store School Bounds
```
BEGIN StoreBounds
    LISTEN FOR school_week, [[period, start_time, end_time]], [term_one, start_time, end_time]... SAVE IN environment.json
END StoreBounds
```

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

### Create a User
```
BEGIN CreateUser
    display username_field
    display email_field
    display password_field
    
    IF submit_button.event('click') AND all_fields NOT == null
         POST (json_web_token, user_type, username, email, password) TO server
         RECEIVE is_success, error
            
         IF NOT is_success
             print(error)
             
         ELSE
             REDIRECT TO /admin/<user_type:string>/
             
    ELSE
         highlight null fields
         
    END IF
END CreateAdminUser
```

### Create User Server Side
```
BEGIN CreateUser
    let json_web_token, user_type, username, email, password = LISTEN FOR json_web_token, username, email, password
    
    IF VerifyUser(json_web_token, admin)
        let password_hash = SaltAndHash(password)
    
        STORE username, email, password_hash in database table account_type
    
        IF STORE Failed
            IF username not unique error
                RESPOND false, username not unique
            ELSE IF email not unique error
                RESPOND false, email not unique
            ELSE
                RESPOND false, store_error
            END IF
        END IF
        
    ELSE
        RESPOND true
END CreateUser
```

### View User
```
BEGIN ViewUser
    GET user WITH user_type, user_id FROM server
    
    let username_field = display user.username
    let email_field = display user.email
    let password_field = empty string
    let confirm_password_field = empty string
    
    let changed_state = false
    
    IF username_field changed OR email_field changed OR lengthof(password_field) > 0
        changed_state = true
    END IF
    
    IF submit.event(click) AND changed_state
        POST (json_web_token, user.id, user.type, username_field, email_field, password_field) TO server
    END IF
    
    IF delete.event(click)
        POST json_web_token, user.type, user.id TO server
    END IF
END ViewUser
```

### Read User
```
BEGIN ReadUser
    let user_type, user_id = LISTEN FOR user_type, user_id
    
    IF user_id IN database table user_type
        RESPOND user_type[user_id]
    ELSE
        RESPOND user not found error
    END IF
END ReadUser
```

### Update User
```
BEGIN UpdateUser
    let json_web_token, user_id, user_type, username, email, password = LISTEN FOR json_web_token, user_id, username, email, password
    
    IF VerifyToken(json_web_token, admin)
        IF user_id IN database table user_type
            let password_hash = HashAndSalt(password)
            
            STORE user_id, username, email, password_hash IN database table user_type
        END IF
    END IF
END UpdateUser
```

### Delete User
```
BEGIN DeleteUser
    let json_web_token, user_type, user_id = LISTEN FOR json_web_token, user_type, user_id
    
    IF VerifyToken(json_web_token, admin)
        DELETE user_id FROM database table user_type
    END IF
END DeleteUser
```

### View List of Classes
```
BEGIN ViewClassesList
    GET classes_list USING json_web_token FROM server
    
    COUNT EACH class IN classes_list
        display class
    END COUNT EACH
    
END ViewClassesList
```

### View List of Classes Server Side
```
BEGIN ViewClassesList
    LISTEN FOR json_web_token
    
    IF VerifyToken(json_web_token, admin)
        let classes_list = []
        
        COUNT EACH class IN database table Classes
            classes_list.append(class)
        END COUNT EACH
        
        RESPOND classes_list
    END IF
END ViewClassesList
```

### Create Class
```
BEGIN CreateClass
    display input class_id_field
    display select year_group_field
    display select teacher_field
    
    let student_array = []
    COUNT EACH student IN student_array
        display select student_field
    END COUNT EACH
    
    IF add_student.event(click)
        student_array.append_student
    END IF
    
    let period_array = [[select_period, select_day]
    COUNT EACH period IN periods
        display select period[select_period]
        display select period[select_day]
    END COUNT EACH
    
    IF add_period.event(click)
        period_array.append([select_period, select_day])
    END IF
    
    POST json_web_token, class_id_field, year_group_field, teacher_field, student_array, period_array TO server
    IF RESPONSE == error
        display error
    END IF
END CreateClass
```

### Create Class Server Side
```
BEGIN CreateClass
    LISTEN FOR json_web_token, class_id, year_group, teacher, student_array, period_array
    
    IF VerifyToken(json_web_token, admin)
        STORE class_id, year_group, teacher IN database table Classes
        get class_primary_key WITH class_id FROM database table Classes
    
        COUNT EACH student IN student_array
            STORE class_id, student.id IN database table student_class
        END COUNT EACH
    
        COUNT EACH periods IN period_array
            STORE periods[period], periods[day] IN database table timetable
        END COUNT EACH
    
        IF error
            RESPOND error
        END IF
    END IF
END CreateClass
```

### View Class
```
BEGIN ViewClass
    GET Class, students_array, periods_array WITH json_web_token, class_primary_key FROM server
    
    let changed_state = false
    
    display input class_id_field = class.class_id
    display select year_group_field = class.year_group
    display select teacher_field = class.teacher_id
    
    let student_array = []
    COUNT EACH student IN student_array
        display select student_field
    END COUNT EACH
    
    IF add_student.event(click)
        student_array.append_student
        changed_state = true
    END IF
    
    let period_array = [[select_period, select_day]
    COUNT EACH period IN periods
        display select period[select_period]
        display select period[select_day]
    END COUNT EACH
    
    IF add_period.event(click)
        period_array.append([select_period, select_day])
        changed_state = true
    END IF
    
    IF changed_state
        POST json_web_token, Class.id, class_id_field, year_group_field, teacher_field, student_array, period_array TO server
        IF RESPONSE == error
            display error
        END IF
    END IF
    
    IF delete_user.event(click)
        POST User.id TO server
    END IF
END ViewClass
```

### Read Class
```
BEGIN ReadClass
    LISTEN FOR get_request, json_web_token, class_primary_key
    
    IF VerifyToken(json_web_token, admin)
        get class, students_list, periods_list FROM database table Classes, student_class, periods
        RESPOND class, students_list, periods_list
    ELSE
        RESPOND error
    END IF
END ReadClass
```

### Update Class
```
BEGIN UpdateClass
    LISTEN FOR json_web_token, class_id_int, class_id_string, year_group, teacher, student_array, period_array
    
    IF VerifyToken(json_web_token, admin)
        get teacher_id using teacher in dataabase table Teachers
        STORE class_id_string, year_group, teacher_id IN database table Classes USING class_id_int
        
        let student_id_array = []
        COUNT EACH student IN student_array
            get student_id using student from database table Students
            student_id_array.append(student)
        END COUNT EACH
        
        COUNT EACH student_id IN student_id_array
            STORE class_id_int, student_id IN database table student_class
        END COUNT EACH
        
        COUNT EACH period IN period_array
            STORE period, class_id_int IN database table timetable
        END COUNT EACH
        
        IF error
            RESPOND error
        END IF
    
    ELSE
        RESPOND error
    
    END IF
END UpdateClass
```

### Delete Class
```
BEGIN DeleteClass
    LISTEN FOR json_web_token, class_id
    
    IF VerifyToken(json_web_token, admin)
        DELETE class class_id FROM database table Classes
    
        COUNT EACH item IN database table student_class USING class_id
            DELETE item
        END COUNT EACH
    
        COUNT EACH item IN database table timetable USING class_id
            DELETE item
        END IF
    END IF
END DeleteClass
```

## Teachers

### View Students List
```
BEGIN ViewStudentsList
    GET list_of_students FROM server WITH json_web_token
    
    COUNT EACH student IN list_of_students
        display student
    END COUNT EACH
END ViewStudentsList
```

### View Students List Server Side
```
BEGIN ViewStudentsList
    LISTEN FOR json_web_token
    
    IF VerifyToken(json_web_token, teacher)
        get list_of_students from database table Student
    END IF
END ViewStudentsList
```

### View Timetable
```
BEGIN ViewTimetable
    GET timetable FROM server WITH json_web_token, teacher_id, date.now()
    
    let school_day = RESPONSE
    IF school_day
        let periods, classes_periods = RESPONSE
        COUNT EACH period IN periods
            IF period IN classes_periods
                display classes_periods[period[0]]
                display classes_periods[period[1]]
                display classes_periods[period[2]]
            ELSE
                display free period
            END IF
        END COUNT EACH
        
    ELSE
        let holiday = RESPONSE
        display holiday
        
    END IF
END ViewTimetable
```

### View Timetable Server Side
* note: server uses clients time for the sake of internationalisation
```
BEGIN ViewTimetable
    LISTEN FOR json_web_token, teacher_id, date
    
    IF VerifyToken(json_web_token, teacher)
        let term_one = AccessPersistantConstants[term_one]
        let term_two = AccessPersistantConstants[term_two]
        let term_three = AccessPersistantConstants[term_three]
        let term_four = AccessPersistantConstants[term_four]
        let holidays = AccessPersistantConstants[holidays]
        
        let holiday: boolean
        let term: term
        
        MATCH date
            CASE term_one[start] > date AND term_one[end] < date
                holiday = false
                term = term_one
                BREAk
            CASE term_two[start] > date AND term_two[end] < date
                holiday = false
                term = term_two
                BREAK
            CASE term_three[start] > date AND term_three[end] < date
                holiday = false
                term = term_three
                BREAK
            CASE term_four[start] > date AND term_four[end] < date
                holiday = false
                term = term_four
                BREAK
            ELSE
                holiday = true
                term = holidays
        END MATCH
        
        IF NOT holiday AND date IN holidays
            holiday = true
        END IF
        
        IF NOT holiday
            RESPOND NOT holiday
            let week_rotation = AccessPersistantConstant(weeks)
            let day = 0
            COUNT d = term.start_day, d == date, d += 1
                IF week_rotation == biweekly
                    IF day >= 14
                        day = 1
                    END IF
                    day += 1
               
                ELSE IF week_rotation == weekly
                    IF day >= 7
                        day = 1
                    END IF
                    day +=1
               
                END IF
            END COUNT
            get periods from database table timetable with teacher_id
            RESPOND periods, AccessPersistantConstant(periods)
        
        ELSE
            RESPOND holiday
            RESPOND holidays[date]
        
        END IF
    END IF
END ViewTimetable
```

### View Class
```
BEGIN ViewClass
    GET Class, students_list, periods FROM server WITH class_id, json_web_token
    
    display class_name
    display class_year
    display class_teacher
    FOR EACH student IN students_list
        display student
    END FOR EACH
    FOR EACH period IN periods
        display period[period]
        display period[date]
    END FOR EACH
END ViewClass
```

### View Class Server Side


### Create Behaviour Notes


### Read Behaviour Notes


### Update Behaviour Notes


### Delete Behaviour Notes


### Create Task


### Create Task Server Side


### View Task


### Read Task


### Update Task


### Delete Task


### Mark Students Response


### Mark Students Response Server Side
