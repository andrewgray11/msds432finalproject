# msds432finalproject
Files and readme for Chicago BI Data ingestion and presentation. 

Requirement 1: (10 Points) Create a Readme file that has a list of every program/microservice that you implemented. 

ProgreSQL database:  Create ProgreSQL datalake to store the data from Chicago public API

Docker: bouid docker containers to allow the application run at any machines

Go functions for data collection:
  
  Building Permit: Insert building permit rows from Chicago data source API into progreSQL database
  
  CCVI: Insert CCVI rows from Chicago data source API into progreSQL database
  
  Census data:  Insert Census data  rows from Chicago data source API into progreSQL database
  
  Community bound: Insert community bound rows from Chicago data source API into progreSQL database
  
  CovidCaseZIp: Insert CovidCaseZIp rows from Chicago data source API into progreSQL database
  
  covidDaily: Insert covidDaily rows from Chicago data source API into progreSQL database
  
  pubHealthData: Insert pubHealthData rows from Chicago data source API into progreSQL database
  
  taxiTirp:  Insert taxiTrip rows from Chicago data source API into progreSQL database
  
  taxiNetProviderTrip:  Insert taxiNetProviderTrip rows from Chicago data source API into progreSQL database

Python for data cleaning/processing to generate report tables

Tableau for frontend data visualization



Requirement 2: (10 Points) Document in the Readme file the steps that are needed to install and run your project deliverables. 

1. Using the build functions to collect data from Chicago public API
2. Store data into ProgreSQL  data lake
3. Manipulate data/ Combine tables to generate report tables and store in the ProgreSQL data warehouse for BI client request
4. When clients request reports, the server will grab report tables from data warehouse and able to visualize on frontend applications
5. Update datasouces daily at midnight, repeat step 1,2, and 3.
