package main

import (
	"database/sql"
	"os"

	"examenFinal/cmd/server/handler"
	"examenFinal/docs"
	"examenFinal/internal/appointment"
	"examenFinal/internal/dentist"
	"examenFinal/internal/patient"
	appoinmentS "examenFinal/pkg/store/appointmentS"
	"examenFinal/pkg/store/dentistS"
	"examenFinal/pkg/store/patientS"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title Clinic Appointment System API
//@version 1.0
//@description This API Handle dental appointments
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	db, err := sql.Open("mysql", "root:root@/appointments_go")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storaged := dentistS.NewSqlStore(db)

	repod := dentist.NewRepository(storaged)
	serviced := dentist.NewService(repod)
	dentistHandler := handler.NewDentistHandler(serviced)

	storagep := patientS.NewSqlStore(db)

	repop := patient.NewRepository(storagep)
	servicep := patient.NewService(repop)
	patientHandler := handler.NewPatientHandler(servicep)

	storagea := appoinmentS.NewSqlStore(db)

	repoa := appointment.NewRepository(storagea)
	servicea := appointment.NewService(repoa)
	appointmentHandler := handler.NewAppointmentHandler(servicea)

	r := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.GET("", dentistHandler.GetAll())
		dentists.POST("", dentistHandler.Post())
		dentists.DELETE(":id", dentistHandler.Delete())
		dentists.PATCH(":id", dentistHandler.Patch())
		dentists.PUT(":id", dentistHandler.Put())
	}
	patients := r.Group("/patients")
	{
		patients.GET(":id", patientHandler.GetByID())
		patients.GET("", patientHandler.GetAll())
		patients.POST("", patientHandler.Post())
		patients.DELETE(":id", patientHandler.Delete())
		patients.PATCH(":id", patientHandler.Patch())
		patients.PUT(":id", patientHandler.Put())
	}
	appointments := r.Group("/appointments")
	{
		appointments.GET(":id", appointmentHandler.GetByID())
		appointments.GET("dni/:dni", appointmentHandler.GetByDNI())
		appointments.GET("", appointmentHandler.GetAll())
		appointments.POST("", appointmentHandler.Post())
		appointments.POST(":dni/:lic", appointmentHandler.PostDP())
		appointments.DELETE(":id", appointmentHandler.Delete())
		appointments.PATCH(":id", appointmentHandler.Patch())
		appointments.PUT(":id", appointmentHandler.Put())
	}

	r.Run(":8080")
}
