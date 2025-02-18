package main

import (
	appStudents "arquitecturahex/src/products/Students/application"
	repoStudents "arquitecturahex/src/products/Students/infraestructure/repositories"
    dependenciesStudents "arquitecturahex/src/products/Students/infraestructure/controllers"
    routesStudents "arquitecturahex/src/products/Students/infraestructure/routes"


	appSubject "arquitecturahex/src/products/Subject/application"
    repoSubject "arquitecturahex/src/products/Subject/infraestructure/repositories"
    dependenciesSubject "arquitecturahex/src/products/Subject/infraestructure/controllers"
    routesSubject "arquitecturahex/src/products/Subject/infraestructure/routes"

    "arquitecturahex/src/core"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    core.InitDB()
    defer core.CloseDB()

    router := gin.Default()

    studentRepo := repoStudents.NewStudentRepositoryImpl(core.DB)
    subjectRepo := repoSubject.NewSubjectRepositoryImpl(core.DB)

    createStudentUseCase := appStudents.NewCreateStudentUseCase(studentRepo)
    viewStudentUseCase := appStudents.NewViewStudentUseCase(studentRepo)
    updateStudentUseCase := appStudents.NewUpdateStudentUseCase(studentRepo)
    deleteStudentUseCase := appStudents.NewDeleteStudentUseCase(studentRepo)
    viewStudentIDUseCase := appStudents.NewViewStudentIdUseCase(studentRepo)

    createSubjectUseCase := appSubject.NewCreateSubjectUseCase(subjectRepo)
    viewSubjectUseCase := appSubject.NewViewSubjectUseCase(subjectRepo)
    updateSubjectUseCase := appSubject.NewUpdateSubjectUseCase(subjectRepo)
    deleteSubjectUseCase := appSubject.NewDeleteSubjectUseCase(subjectRepo)

    studentController := dependenciesStudents.NewStudentController(
        createStudentUseCase,
        viewStudentUseCase,
        updateStudentUseCase,
        deleteStudentUseCase,
        viewStudentIDUseCase,
    )

    subjectController := dependenciesSubject.NewSubjectController(
        createSubjectUseCase,
        viewSubjectUseCase,
        updateSubjectUseCase,
        deleteSubjectUseCase,
    )

    routesStudents.RegisterStudentRoutes(router, studentController)
    routesSubject.RegisterSubjectRoutes(router, subjectController)

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }

    router.Run(":" + port)
}