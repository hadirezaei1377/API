Target :
Full review of the api along with the use of the http protocol

According to what is defined in the http protocol, we will have a client that can be a human, a machine, a web browser, or anything that is sending a request.
We also have a server that should respond appropriately to this request.
This server evaluates the validity of the request and returns a response.
If we want to create a server whose response is in json format, we use restful.
restful is a structure for us that determines how and in what format we send what request and what response we receive. (determiner of formatting)

We use the Echo framework to run the code on the server.

To better understand how to run the server, there is a code snippet called server setup in the folders.

In the controller folder, the overall management of the project is supposed to take place, for example, determining which request and answer to provide is one of the duties of this department.

model folder:
To explain this folder and its function, we can say:
When we want to transfer the data or make a change in them, we have to convert them into objects or structures so that we can manage them.
These structures that specify our data are placed inside this folder

config folder:
All settings of different parts of the project are placed here.

database:
All tasks related to working with the database are placed here

The overall structure of the project in main function:

get config ,Everything related to the implementation of the project, which requires that we have it before the implementation of the project and execute the project relative to it, is in the config folder, we call them in the main function so that it is placed in the memory.

init server ,Related to making a server

start server ,When the server runs here, it executes the tasks we have defined for it.
In the main function, everything between init server and start server are server functions,
like route

routing ,Know how and when the request comes, what response to send.

middleware ,For some special orders


