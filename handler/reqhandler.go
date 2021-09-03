package handler

import (
	"strings"
	"strconv"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"	
	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/deepakbatham0007/sugarbox/database"
	"github.com/deepakbatham0007/sugarbox/model"
)

var router *gin.Engine

func Login(c *gin.Context) {
	user := c.PostForm("user")
	pass := c.PostForm("pass")
	fname,err:=database.CheckUser(user,pass)
	if err != nil {
	   fmt.Println(err)
	}

	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge:   60 * 30})	
	session.Set("user", fname)
	session.Save()

	sessUser:=session.Get("user")
	c.HTML(http.StatusOK, "index.html", gin.H{"sessUser": sessUser,})
}

func Logout(c *gin.Context){
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "",})
}

func AddItemType(c *gin.Context){
	seltype := c.PostForm("type")
	typeName := c.PostForm("typeName")
	directorName := c.PostForm("directorName")
	artist := c.PostForm("artist")
	fmt.Println(seltype,typeName,directorName,artist)

	database.Add(seltype,typeName,directorName,artist)

	session := sessions.Default(c)
	sessUser:=session.Get("user")
	if sessUser == nil {
		sessUser=""
	}	
	
	c.HTML(http.StatusOK, "index.html", gin.H{"sessUser": sessUser,})
}


func Middlew(c *gin.Context){
	findItem := c.PostForm("findItem")
	var t model.SearchName
	
	if err := c.BindJSON(&t); err != nil {
		fmt.Println("------------------111",err);
	}

fmt.Println(findItem,"------------------midle",t);
c.Next()
}

func SearchUserItems(c *gin.Context){
	var t model.SearchUserItems
	if err := c.BindJSON(&t); err != nil {
		fmt.Println("------------------111",err);
		c.JSON(200, gin.H{"message": "Invalid request...",})
		return
	}
	res,err:=database.SearchUserAllItems(t.Userid)
	if err != nil {
		c.JSON(200, gin.H{"message": err,})
		return 
	}
	var res1 []model.SearchUserResults
	json.Unmarshal([]byte(res),&res1)
	c.JSON(200, res1)	
}


func SearchItem(c *gin.Context){
findItem := c.PostForm("FindItem")

	var t model.SearchName
	flag:=false
	if findItem != ""{
		t.FindItem=findItem
	}else{
		flag=true
		if err := c.BindJSON(&t); err != nil {
			fmt.Println("------------------111",err);
		}
	}

//fmt.Println(findItem,"------------------midle",t);
//	fmt.Println("------------------midle---",t);

	session := sessions.Default(c)
	sessUser:=session.Get("user")
	if sessUser == nil {
		sessUser=""
	}

	str := fmt.Sprintf("%v", sessUser)
	res,err:=database.SearchItem(t.FindItem,str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(flag,"------------------midle123---",res,"===");

	if flag {
		if strings.TrimSpace(res)!=""{
			var res1 model.SearchRes
			json.Unmarshal([]byte(res),&res1)
			c.JSON(200, res1)
		}else{
			c.JSON(200, gin.H{"message": "No item found, try again...",})
		}
	}else{
		//renderToHtml(res)
		
		if sessUser == "" {
			var res1 model.SearchRes
			json.Unmarshal([]byte(res),&res1)			
			c.HTML(http.StatusOK, "index.html", gin.H{"result": res1,"sessUser": sessUser,})
		}else{
			var res1 model.SearchUserResults
			json.Unmarshal([]byte(res),&res1)			
			c.HTML(http.StatusOK, "index.html", gin.H{"result": res1,"sessUser": sessUser,})
		}
		
	}

	
}

/*func renderToHtml(res string){
	var res1 SearchRes
	 json.Unmarshal([]byte(res),&res1)
	fmt.Println(res1.Name)
}*/

func RateItem(c *gin.Context){
	name:=c.PostForm("Name")
	var t model.RatingModel
	//flag:=false
	if name != ""{
		t.UserId=c.PostForm("UserId")
		t.Name=c.PostForm("Name")
		t.Rating,_=strconv.Atoi(c.PostForm("Rating"))
		//fmt.Println(t.UserId,t.Name, t.Rating)
		res,err:=database.AddRating(t.UserId,t.Name, t.Rating)
		if err != nil {
			fmt.Println(err)
				c.HTML(http.StatusOK, "index.html", gin.H{"result": err,"sessUser": t.UserId,})
		}else{
			c.HTML(http.StatusOK, "index.html", gin.H{"result": res,"sessUser": t.UserId,})
		}

	}else{
		//flag=true
		if err := c.BindJSON(&t); err != nil {
			
			fmt.Println("------------------111",err);
		}
		fmt.Println(t.UserId,t.Name,"rating......",t)
		if t.Rating >5 || t.Rating < 1{
			c.JSON(200, gin.H{"message": "Please give a valid Rating 1-5 ...",})
			return 	
		}
		res,err:=database.AddRating(t.UserId,t.Name, t.Rating)
		if err != nil {
			fmt.Println(err)
				c.JSON(200, gin.H{"message": err,})
		}else{
			c.JSON(200, gin.H{"message": res,})
		}

		
	}

	

}

func CommentItem(c *gin.Context){

	name := c.PostForm("Name")
	var t model.CommentModel
	//flag:=false
	if name != ""{
		t.UserId=c.PostForm("UserId")
		t.Name=name
		t.Comment=c.PostForm("Comment")
		res,err:=database.AddComment(t.UserId,t.Name, t.Comment)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{"result": err,"sessUser": t.UserId,})
		}else{
			c.HTML(http.StatusOK, "index.html", gin.H{"result": res,"sessUser": t.UserId,})
		}
		
	}else{
		//flag=true
		if err := c.BindJSON(&t); err != nil {
			
			fmt.Println("------------------111",err);
		}
		fmt.Println(t.UserId,t.Name,"Comment......",t.Comment,"===")
		if strings.TrimSpace(t.Comment) == "" {
			c.JSON(200, gin.H{"message": "Please give a valid Comment ...",})	
			return 		
		}
		res,err:=database.AddComment(t.UserId,t.Name, t.Comment)
		if err != nil {
			c.JSON(200, gin.H{"message": err,})
		}else{
			c.JSON(200, gin.H{"message": res,})
		}

		
	}

	

}



