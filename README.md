# miniwallet Service

### How to run:

This project depedency is managed by go mod, therefore you can put it everywhere other than GO-PATH directory.
steps to run using docker-compose:

    1. cd to project root directory
    
    2. run on terminal : 
        docker-compose up --build -d
    or by using Makefile, run on terminal : 
        make run
        
    3. verify the container is running by run:
        docker ps
        
    4. Execute the call
        
    5. To Stop run on terminal:
        make stop

        

### URL
- `http://localhost/api/v1`

