# EmployeeAPI
A GO API that connects to a MySQL DB to store Employee Data

# Set up
Create a local connection in MySQL (127.0.0.1:3306) with username: root an password:password OR change the connection string in the sql.Open statement in the main method.

## Create Database
CREATE DATABASE employees;
USE employees;

## Create Table
CREATE TABLE employees (
name varchar(255),
email varchar(255),
role varchar(255),
skills varchar(255)
);

## Run the application
From the directory with the files run the following commands:
go build *.go
go run *.go

## Load in some data
POST to http:localhost:8081/employees with Body:
{
    "name": "John Doe",
    "email": "John@Doe.com",
    "role": "Developer",
    "skills": "React, AWS, Java"
}
