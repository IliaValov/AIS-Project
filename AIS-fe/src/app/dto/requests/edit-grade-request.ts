export class EditGradeRequest {
    courseId: string
    studentId: string
    grade: string;

    constructor(courseId: string, studentId: string, grade: string) {
        this.courseId = courseId;
        this.studentId = studentId;
        this.grade = grade;
    }
}