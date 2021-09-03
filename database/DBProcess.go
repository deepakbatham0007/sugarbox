// This package have the database related all files
package database

// Import require package for mysql db
import (
	"errors"
	"fmt"
	"encoding/json"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/deepakbatham0007/sugarbox/model"
)

func CheckUser(user, pass string) (string, error) {  //DATABASE Connection
	db := dbConn()
	defer db.Close()
	var fName string
	selDB := db.QueryRow("SELECT ufname FROM tbluser where uid='"+user+"' and upass='"+pass+"'")
	err := selDB.Scan(&fName)
	if err != nil {
	   return "",err
	}
	return fName,nil
}

func Add(seltype,typeName,directorName,artist string){
	db := dbConn()
	
	insert, err := db.Query("INSERT INTO tblmovieall (typename,seltype,directorname,artist) VALUES ( '"+typeName+"','"+seltype+"','"+directorName+"','"+artist+"' )")

	if err != nil {
        	panic(err.Error())
	}
	
	insert.Close()
	defer db.Close()
}

func SearchItem(itemName string,sessUser string)(string,error){
	db := dbConn()
	defer db.Close()
	qry:="select t1.typeName as Name,t1.seltype as Type,t1.directorName as Director,t1.artist as Artist, "
		qry+="case when avg(tr1.rating) IS NULL then 0 else CAST(AVG(tr1.rating) AS DECIMAL(5, 0)) end  as Rating,"

		if sessUser == "" {
			qry+="case when count(tr1.userId) IS NULL then 0 else count(tr1.userId) end  as NoOfRatedUser"
			qry+=" FROM sugarboxdb.tblmovieall  t1"
			qry+=" left join sugarboxdb.tblmoviealltrans tr1 on t1.idx=tr1.typeId"
			qry+=" where t1.typeName='"+itemName+"'"			
		}else{
			qry+="case when tr1.userComment is null then '' else tr1.userComment end as Comment "
			qry+=" FROM sugarboxdb.tblmovieall  t1"
			qry+=" left join sugarboxdb.tblmoviealltrans tr1 on tr1.typeId = t1.idx"
			qry+=" and tr1.userId=(select idx from sugarboxdb.tbluser where ufname='"+sessUser+"')"
			qry+=" where t1.typeName='"+itemName+"'"
		}
		
	//fmt.Println(qry)

	
	if sessUser == "" {
		res :=new(model.SearchRes)
		fmt.Println("user nil..........")
		selDB := db.QueryRow(qry)
		err := selDB.Scan(&res.Name,&res.Type,&res.Director,&res.Artist,&res.Rating,&res.NoOfRatedUser)
		if err != nil {
		fmt.Println(err)
		   return "",err
		}
		//fmt.Println(res)
		content, _ := json.Marshal(res)
		return string(content), nil
	}else{
		fmt.Println("user not nil..........",qry)
		res :=new(model.SearchUserResults)
		selDB := db.QueryRow(qry)
		err := selDB.Scan(&res.Name,&res.Type,&res.Director,&res.Artist,&res.Rating,&res.Comment)
		if err != nil {
		fmt.Println(err)
		   return "",err
		}
		//fmt.Println(res)
		content, _ := json.Marshal(res)
		return string(content), nil
	}
}

func SearchUserAllItems(userid string)(string,error){
	db := dbConn()
	defer db.Close()
	qry:="select t1.typeName as Name,t1.seltype as Type,t1.directorName as Director,t1.artist as Artist, "
		qry+="tr1.rating as Rating ,tr1.userComment as UserComment "
		qry+="FROM sugarboxdb.tblmovieall  t1 "
		qry+="left join sugarboxdb.tblmoviealltrans tr1 on t1.idx=tr1.typeId "
		qry+="where tr1.userId=(select idx from sugarboxdb.tbluser where uid='"+userid+"')"
	fmt.Println(qry)

	
	results, err := db.Query(qry)
	if err != nil {
		return "",err
	}
	result:=[]model.SearchUserResults{}
	for results.Next() {
		var res model.SearchUserResults
		err := results.Scan(&res.Name,&res.Type,&res.Director,&res.Artist,&res.Rating,&res.Comment)
		if err != nil {
			return "",err
		}
		result=append(result,res)
	}

	content, _ := json.Marshal(result)
	return string(content), nil
}
func AddRating(uid,name string, Rating int)(string,error){
	db := dbConn()
	defer db.Close()

	var userid int
	selDB := db.QueryRow("select idx FROM sugarboxdb.tbluser where uid='"+uid+"'")
	err := selDB.Scan(&userid)
	if err != nil {
	   //fmt.Println(err)
	   return "", errors.New("Invalid user.")
	}

	var itemId int
	selDB = db.QueryRow("SELECT idx FROM sugarboxdb.tblmovieall where typename='"+name+"'")
	err = selDB.Scan(&itemId)
	if err != nil {
		//fmt.Println(err)
	   return "",errors.New("Invalid Item for rating.")
	}

	
	fmt.Println(userid,"-----",itemId);

	results, _ := db.Query("SELECT * FROM tblmoviealltrans where typeId="+fmt.Sprintf("%v", itemId)+" and userId="+fmt.Sprintf("%v", userid))
	
    	if results.Next() {
		stmt, err := db.Prepare("UPDATE tblmoviealltrans set  rating=? where typeId=? and userId=?")
	
		if err != nil {
	        	panic(err.Error())
		}
	
		 _, er := stmt.Exec(Rating,itemId,userid)
		    if er != nil {
		        panic(er.Error())
		} 
		
		stmt.Close()
	
		return ("Comment updated successfully for "+name),nil		
	}else{
		stmt, err := db.Prepare("INSERT INTO tblmoviealltrans (typeId,userId,rating) VALUES (?,?,?)")
	
		if err != nil {
	        	panic(err.Error())
		}
	
		 _, er := stmt.Exec(itemId,userid,Rating)
		    if er != nil {
		        panic(er.Error())
		} 
		
		stmt.Close()
	
		return ("Rating updated successfully for "+name),nil
	}
}

func AddComment(uid,name string, comment string)(string,error){
	db := dbConn()
	defer db.Close()

	var userid int
	selDB := db.QueryRow("select idx FROM sugarboxdb.tbluser where uid='"+uid+"'")
	err := selDB.Scan(&userid)
	if err != nil {
	   //fmt.Println(err)
	   return "", errors.New("Invalid user.")
	}

	var itemId int
	selDB = db.QueryRow("SELECT idx FROM sugarboxdb.tblmovieall where typename='"+name+"'")
	err = selDB.Scan(&itemId)
	if err != nil {
		//fmt.Println(err)
	   return "",errors.New("Invalid Item for rating.")
	}

	
	fmt.Println(userid,"-----",itemId);
	results, _ := db.Query("SELECT * FROM tblmoviealltrans where typeId="+fmt.Sprintf("%v", itemId)+" and userId="+fmt.Sprintf("%v", userid))
	
    	if results.Next() {
		stmt, err := db.Prepare("UPDATE tblmoviealltrans set  userComment=? where typeId=? and userId=?")
	
		if err != nil {
	        	panic(err.Error())
		}
	
		 _, er := stmt.Exec(comment,itemId,userid)
		    if er != nil {
		        panic(er.Error())
		} 
		
		stmt.Close()
	
		return ("Comment updated successfully for "+name),nil		
	}else{
		stmt, err := db.Prepare("INSERT INTO tblmoviealltrans (typeId,userId,userComment) VALUES (?,?,?)")
	
		if err != nil {
	        	panic(err.Error())
		}
	
		 _, er := stmt.Exec(itemId,userid,comment)
		    if er != nil {
		        panic(er.Error())
		} 
		
		stmt.Close()
	
		return ("Comment updated successfully for "+name),nil
	}
}


