# Fiber-Gorm-Project
A fiber-gorm project. This was my Internet-Engineering Final and I implemented it with GO for practice.

# Description
We have Student and Course entities and GradeReport as a bridge entity .\
We store Student's informations in Student and Course's information in Course. \
In GradeReport, we save student's grades for course : Student X has Grade 19 in Course Y \
I used SQL datbase and you can use sqlite extension in vscode to obeseve it.\
You can use Postman json file to send declared requests:
### Student Requests:
* Create Student
* Get a Student informations
* Get all Students' informations
* Update a Student's informations
* Delete a Student

### Course Requests:
* Create Course
* Get a Course's informations
* Get All Courses' informations
* Update a Course's informations
* Delete a Course

### GradeReport Requests:
* Enter a Grade for Student
* Get a Student's Grades Report
* Update a Student's Grade
* Delete a Student's Grade
