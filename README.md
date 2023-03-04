# msds432finalproject
Files and readme for Chicago BI Data ingestion and presentation. 

** Requirement 1: (10 Points) Create a Readme file that has a list of every program/microservice that you implemented. 

PostgreSQL database:  Create PostgreSQL datalake to store the data from Chicago public API

Docker: bouid docker containers to allow the application run at any machines

Go functions for data collection:
  
  Building Permit: Insert building permit rows from Chicago data source API into PostgreSQL database
  
  CCVI: Insert CCVI rows from Chicago data source API into PostgreSQL database
  
  Census data:  Insert Census data  rows from Chicago data source API into PostgreSQL database
  
  Community bound: Insert community bound rows from Chicago data source API into PostgreSQL database
  
  CovidCaseZIp: Insert CovidCaseZIp rows from Chicago data source API into PostgreSQL database
  
  covidDaily: Insert covidDaily rows from Chicago data source API into PostgreSQL database
  
  pubHealthData: Insert pubHealthData rows from Chicago data source API into PostgreSQL database
  
  taxiTirp:  Insert taxiTrip rows from Chicago data source API into PostgreSQL database
  
  taxiNetProviderTrip:  Insert taxiNetProviderTrip rows from Chicago data source API into PostgreSQL database

Python for data cleaning/processing to generate report tables

Tableau/PowerBI for frontend data visualization



** Requirement 2: (10 Points) Document in the Readme file the steps that are needed to install and run your project deliverables. 

Install and download go, ProgreSQL and docker.  Download program from github. 
1. Using the build functions to collect data from Chicago public API
2. Store data into PostgreSQL  data lake
3. Manipulate data/ Combine tables to generate report tables and store in the ProgreSQL data warehouse for BI client request
4. When clients request reports, the server will grab report tables from data warehouse and able to visualize on frontend applications
5. Update datasouces daily at midnight, repeat step 1,2, and 3.
