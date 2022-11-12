# Portfolio GOLANG(go language) payment with https://www.mitrais.com

**Deploy with local Docker**
- docker-compose up

**Documentation**
- localhost:8080/api/v1/users
  body : 
  ```json 
    {
    "name":"Hendaru",
    "occupation":"Programmer",
    "email":"hendaru@gmail.com",
    "password":"123456"
}
  ```

- localhost:8080/api/v1/sessions
  body:
  ```json
    {
    "email":"hendaru@gmail.com",
    "password":"123456"
}
  ```