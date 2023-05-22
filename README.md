#This server has 6 kind of operations:

#RegisterUser (POST): "/users" 

- It Registers user details in the Database.
- It Takes four inputs : Name, Email, Date Of Birth, and Password. No field can be empty.
- It returns ID and creation time. The ID is further use to access the database.
- It stores the credentials in PostgreSQL.

#ListUsers (GET): "/users"

- It lists all users registered in the database

#AddMovie (POST): "/movies"

- It Registers Movie details in the Database
- It Takes two inputs : UserID, and title. The UserID is the one mentioned in RegisterUser.
- It returns movieID (referred as id here) and creation time. MovieID is used to access other function here. 
- It stores the data in PostgreSQL

#DeleteMovie (DELETE): "/movies/:id"

- It deletes movies based on the movieID we recieved in AddMovie function.
- If no movie is deleted, it shows an error.

#ListMoviesByUser (GET):

- It Lists all the movies of the user, by its user_id.
- (UPDATE) Now, it requires password along with the user id. Otherwise it will simply skip the tables.  
- (UPDATE) You will get to know if the password is wrong, or user doesn't exist though, so don't worry too much.

#(UPDATE) UpdateCredentials (POST):

- It is used to update your credentials.
- It is not automatic, you have to manually update with the id you get from create user function.
- It is absolutely necessery to do this. Otherwise List Movies by user function will not work.

#Planned Features:

- Integrating Password in other functions to ensure the right user is using the software
- Deleting movies separately per user
- Deleting User and their Credentials
- Merging both Create User and Create Credentials function
