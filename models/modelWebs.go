package models

type Webs struct {
	Name         string `dynamodbav:"Name"`
	NameUser     string `dynamodbav:"NameUser"`
	Url          string `dynamodbav:"Url"`
	DateCreation string `dynamodbav:"DateCreation"`
	Description  string `dynamodbav:"Description"`
	Rol          string `dynamodbav:"Rol"`
	RolCognito   string `dynamodbav:"RolCognito"`
	Phone        int    `dynamodbav:"Phone"`
	Email        string `dynamodbav:"Email"`
}
