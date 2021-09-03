# sugarbox

This is golang application. This application have web and api functionality for movie releted information.

import database into mysql with 
DB : this reposetiry have database file of sql for use in this application sugarboxdb.sql.

	user := "root"
	pass := "123456"
	dbName := "sugarboxdb"

download all go file as it is with proper go environment 1.14 and above.


execute go app with below command: 
>go run sugarbox.go



Web :

Find movie, documentaries and shows 
Add movie, documentaries and shows 

there are 4 pre define user as below with password, user registration not developed.

test1/123456
test2/123456
test3/123456
admin/admin

Admin user can add movie, documentaries and shows, also find the movie.
Other user find the movie, documentaries and shows, also update rating, comment.

Guest only search the movie, documentaries and shows.

API: below api can be proccess postmain, curl, rest client.

API to search for a movie with name 
	POST : http://localhost:8080/search
	post data: 
	{
	  "FindItem": "Love"		//This is value is movie, documentaries and shows name
	}

API where a user can see movies that he has rated and his comments 
	POST : http://localhost:8080/searchuseritems
	post data:
	{
	  "Userid": "test1"		//pass user id
	}	

API to rate a movie
	POST : http://localhost:8080/rating

	{
		"UserId":"test3",		//user id
		"Name": "Maachis",		//movie, documentaries and shows name
	  	"Rating":4			//rating want to give
	}

API to add comment on movie
	POST : http://localhost:8080/comment
	{
		"UserId":"test3",		//user id	
		"Name": "Maachis",		//movie, documentaries and shows name
	  	"Comment":"testing......"	// comment to save
	}

