## Tech test as Go Developer

# Assumptions and Disclaimers

- lack of time, I have started for a template I have on published on github
- the name of the item in the list is the actual primary key
- assumption that framework do not mean, packages, then assumption fasthttp is allowed
- the authentication system is not provided
- due to lack of time I have not written the dockerfile for it

This code have been produced in between the 7.00 am and the 9.30 on Thu, May 9, 2019 as
its first commit version

# Techs

I have used a previous template I have written in using fasthttp
  to speed up the development and becouse is by far faster than net/http

- Go 1.11
- mgo
- mongoDb
  
# Todo

- Rewrite the comments
- Complete the testing Suite


# Future Improvements
- optimization, keep the connection to the database alive all the time instead of opening and closing for each endpoint
