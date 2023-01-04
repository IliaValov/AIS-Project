export class SubjectResponse {
    data: SubjectDto[] = [];
}

export interface SubjectDto {
    CourseName: string;
    GradeNumber: number;
}