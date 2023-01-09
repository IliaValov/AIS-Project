export class StudentResponse {
    data: Student[] = [];
}

export interface Student {
    UserId: number,
    FirstName: string,
    LastName: string,
    Grade: number
}