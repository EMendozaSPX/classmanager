# Technical documentation

This folder contains a list of username and passwords
use the schema icon on the graphql playground app to find
a list of available operations

visit the following link in order to see syntax for querying the server at https://www.thoughtsofjuan.me/graphql
https://graphql.org/learn/queries/

Visit the link and download the application or use online variant https://github.com/prisma/graphql-playground
Enter https://www.thoughtsofjuan.me/graphql in the http directive

first login using the mutation 
`
mutation {
    login(username: username, password: password) {
        token
    }
}
`

and paste the results into the header section as 
`
{
    "authorization": "Bearer {Token}"
}
'

you can now perform queries and mutations allowed for the given user.