### Backend for login / signup

-- Manas Jain Kuniya, SPO ps-2 i have used golang for backend, react for frontend for testing.

### API Endpoints

1. Signup
   Endpoint: /signup
   Method: POST
   Request Body:
   json
   {
   "userId": "exampleUser",
   "password": "examplePassword"
   }
   Response:
   200 OK if registration is successful.
   400 Bad Request if the request body is invalid.
   200 OK if the user already exists.
2. Login
   Endpoint: /login
   Method: POST
   Request Body:
   json
   {
   "userId": "exampleUser",
   "password": "examplePassword"
   }
   Response:
   200 OK if login is successful.
   400 Bad Request if the request body is invalid.
   200 OK if the user is not found.
   401 Unauthorized if the password is incorrect.

### functions

## Helpers

1. Connect() - Connects to MongoDB and returns the users collection.
2. LoadEnv() - Loads environment variables from the .env file.
3. Password Hashing, Hash(password string) (string, error)
4. verifying hash, VerifyHash(password, hash string) bool

## Handlers

1. Signup Handler
   Signup(c *gin.Context, users *mongo.Collection) - Handles user registration.
2. Login Handler
   Login(c *gin.Context, users *mongo.Collection) - Handles user login.

## Models

1. User Model
   User - Defines the user model with UserId and Password fields.

### Database

used mongo db for storing the hashed userid and password
Ensure MongoDB is running and accessible with the provided MONGO_URI, while running the program
