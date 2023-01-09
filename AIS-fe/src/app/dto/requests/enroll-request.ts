export class EnrollRequest {
    studentId: string
    teacherId: string;
    courseId: string

    constructor(courseId: string, studentId: string, teacherId: string) {
        this.courseId = courseId;
        this.studentId = studentId;
        this.teacherId = teacherId;
    }
}