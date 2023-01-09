export class StudentNotEnrolledResponse {
    data: Student[] = [];
}

export interface Student {
    StudentId: number;
    FirstName: string;
    LastName: string;
}