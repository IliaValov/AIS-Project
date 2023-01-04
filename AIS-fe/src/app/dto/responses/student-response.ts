export class StudentResponse {
    data: StudentWithGrade[] = [];
}

export interface Grade {
    StudentId: number,
    Grade: number
}

export interface Student {
    FirstName: string;
    LastName: string;
}

export interface StudentWithGrade {
    Student: Student,
    Grade: Grade
}