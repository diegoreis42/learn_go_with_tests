package main

func Hello(name string, language string) string {
  switch language {

  case "English":
    return "Hello, " + name
  
  case "Portuguese":
    return "Olá, " + name
  
  default:
    return "Hello, World" 
  }
}


