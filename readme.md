API :

We use the restful format in writing the API.
The format that determines what request to send and what response to receive and how to do these processes.
In general, this format determines the structure for us.

Also, to send a request to the API, we must have a test client that uses postman.

we use Echo framework for running server


controller folder : 
Project structure management, Determining what response to request.

model folder :
Whenever we want to transfer data from one place to another, we need to format it and turn it into an object or a structure so that we can manage it in the program.
  Structs define data structures.
  These structures are placed inside the model folder and make the implementation of the project easier.

config folder :
Everything that is related to before the project is executed and the project should be executed against it is placed in the config.
Any settings that are related to the running part of the project are placed here, for example, the information of a file should be read in a part, we put the settings related to this operation in this folder.

database folder :
Any information or operation related to storage and retrieval from the database is placed here.

routing :
Where and to what address are the requests sent and which department should respond to them in what way.

middleware :
Items to be applied to each request or response

usercontroller file :
moving routing to general state for using easier 
assign response to approprate request by using routing

model folder :
anything related to database that we wil use , puted here