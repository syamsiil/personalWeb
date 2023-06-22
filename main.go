package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"personalWeb/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// nama dari struct nya adalah project
// yang membangun dari object/properties
type Project struct  {
	ID 				int
	ProjectName 	string
	StartDate 		string
	EndDate 		string
	Description 	string
	DistanceTime	string
	PostDate 		time.Time
	Javascript		bool
	ReactJS			bool
	NodeJS			bool
	CSS3			bool
	Image			string
	Author			string
	FormatDate		string
}

// Data - data yang ditampung, yang kemudian data yang diisi harus sesuai dengan tipe data yang telah dibangun pada struct 
var dataProject = [] Project{
	{
		ProjectName:    "Design Web Apps 2023",
		StartDate:		"2023-06-01",
		EndDate: 		"2023-06-06",
		Description: 	"Description",
		DistanceTime: 	"1 month",
		Javascript:     true,
		ReactJS:    	true,
		NodeJS:			true,
		CSS3: 			true,
	},
	{
		ProjectName:    "Mobile Apps 2023",
		StartDate:		"2023-06-01",
		EndDate: 		"2023-06-06",
		Description: 	"Description 2",
		DistanceTime: 	"1 month",
		Javascript:     true,
		ReactJS:    	true,
		NodeJS:			true,
		CSS3: 			true,
	},
} 

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	// e = echo package
	// GET =  run the method
	// "/" = endpoint/routing ("localhost:5000 , ex. "/home")
	// helloWorld = function that will run if the route are opened

	// mengirimkan folder sebagai folder statis/ directory public
	e.Static("/public", "public")

	// Routing
	// GET
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/projects", projects)
	e.GET("/detail-project/:id", detailProject)
	e.GET("/form-add-project", formAddProject)
	e.GET("/testimonials", testimonials)

	// POST
	e.POST("/add-project", addProject)
	e.POST("/project-delete/:id", deleteProject)



	e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloWorld(c echo.Context)error {
	return c.String(http.StatusOK, "Hello World")
}

func home(c echo.Context)error{
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	// nil penampung data yang dikirimkan
	return tmpl.Execute(c.Response(), nil) 
}

func contact(c echo.Context)error{
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	// nil penampung data yang dikirimkan
	return tmpl.Execute(c.Response(), nil) 
}

func testimonials(c echo.Context)error{
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	// nil penampung data yang dikirimkan
	return tmpl.Execute(c.Response(), nil) 
}

func projects(c echo.Context)error{
	data, errtest := connection.Conn.Query(context.Background(), "SELECT id, project_name, description, 'image', post_date FROM tb_project")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.ID, &each.ProjectName, &each.Description, &each.Image, &each.PostDate)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.FormatDate = each.PostDate.Format("2 January 2006")
		each.Author = "Abel Dustin"

		result = append(result, each)
	}

	projects := map[string]interface{}{
		"Projects": result,
	}

	fmt.Println(result)


	var tmpl, err = template.ParseFiles("views/projects.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	if errtest != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":errtest.Error()})
	}

	// nil penampung data yang dikirimkan
	return tmpl.Execute(c.Response(), projects) 
}

func detailProject(c echo.Context)error{
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, description, image, FROM tb_project WHERE id=$1",id).Scan(&ProjectDetail.ID,&ProjectDetail.Description, &ProjectDetail.Image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	ProjectDetail.Author = "Seeus"
	ProjectDetail.FormatDate = ProjectDetail.PostDate.Format("2 January 2006")

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/detail-project.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddProject(c echo.Context)error{
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	// nil penampung data yang dikirimkan
	return tmpl.Execute(c.Response(), nil) 
}

func calculateDuration(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}

func addProject(c echo.Context)error{
	projectName := 	c.FormValue("input-name")
	startDate := 	c.FormValue("input-startdate")
	endDate := 		c.FormValue("input-enddate")
	description := 	c.FormValue("input-description")
	distanceTime :=	calculateDuration(startDate, endDate)
	
	var javascript bool
	if c.FormValue("javascript") == "yes" {
		javascript = true
	}
	var reactJs bool
	if c.FormValue("react-js") == "yes" {
		reactJs = true
	}
	var nodeJs bool
	if c.FormValue("node-js") == "yes" {
		nodeJs = true
	}
	var css3 bool
	if c.FormValue("css3") == "yes" {
		css3 = true
	}
	
	var newProject = Project{
		ProjectName:    projectName,
		StartDate:		startDate,
		EndDate: 		endDate,
		Description: 	description,
		DistanceTime: 	distanceTime,
		Javascript:     javascript,
		ReactJS:    	reactJs,
		NodeJS:			nodeJs,
		CSS3: 			css3,
	} 

	// append berfungsi untuk menambahkan data newProject ke dalam slice dataProject
	// mirip dengan fungsi push() pada javascript
	// param1 = dimana datanya ditampung
	// param2 = data apa yang akan ditampung
	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)
	
	return c.Redirect(http.StatusMovedPermanently, "/projects")
} 

func deleteProject(c echo.Context)error {
	id,_ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]... )

	return c.Redirect(http.StatusMovedPermanently, "/projects")
} 



