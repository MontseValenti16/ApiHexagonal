package repositories

import (
    "arquitecturahex/src/products/Students/domain/entities"
    "database/sql"
)

type StudentRepositoryImpl struct{
    db *sql.DB
}

func NewStudentRepositoryImpl(db *sql.DB) *StudentRepositoryImpl {
    return &StudentRepositoryImpl{db : db}
}

func (repo *StudentRepositoryImpl) Save(student entities.Student) error {
    query := "INSERT INTO students (name, email, age, grade) VALUES (?, ?, ?, ?)"
    _, err := repo.db.Exec(query, student.Name, student.Email, student.Age, student.Grade)
    return err
}

func (repo *StudentRepositoryImpl) GetAllStudents() ([]entities.Student, error) {
    query := "SELECT id, name, email, age, grade FROM students"
    rows, err := repo.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []entities.Student
    for rows.Next() {
        var student entities.Student
        if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Age, &student.Grade); err != nil {
            return nil, err
        }
        students = append(students, student)
    }
    return students, nil
}

func (repo *StudentRepositoryImpl) Update(id int, student entities.Student) error {
    query := "UPDATE students SET name = ?, email = ?, age = ?, grade = ? WHERE id = ?"
    _, err := repo.db.Exec(query, student.Name, student.Email, student.Age, student.Grade, id)
    return err
}

func (repo *StudentRepositoryImpl) Delete(id int) error {
    query := "DELETE FROM students WHERE id = ?"
    _, err := repo.db.Exec(query, id)
    return err
}

func (repo *StudentRepositoryImpl) ExistsById(id int) (bool, error) {
    query := "SELECT COUNT(*) FROM students WHERE id = ?"
    var count int
    err := repo.db.QueryRow(query, id).Scan(&count)
    return count > 0, err
}

func (repo *StudentRepositoryImpl) ExistsByEmail(email string) (bool, error) {
    query := "SELECT COUNT(*) FROM students WHERE email = ?"
    var count int
    err := repo.db.QueryRow(query, email).Scan(&count)
    return count > 0, err
}

func (repo *StudentRepositoryImpl) GetById(id int) (*entities.Student, error) {
    query := "SELECT id, name, email, age, grade FROM students WHERE id = ?"
    row := repo.db.QueryRow(query, id)

    var student entities.Student
    if err := row.Scan(&student.ID, &student.Name, &student.Email, &student.Age, &student.Grade); err != nil {
        return nil, err
    }
    return &student, nil}