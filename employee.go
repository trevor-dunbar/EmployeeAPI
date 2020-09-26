package main

type Employee struct {
	Name string `json:"name"`;
	Email string `json:"email"`;
	Skills string `json:"skills"`;
	Role string `json:"role"`;
}

type Employees []Employee