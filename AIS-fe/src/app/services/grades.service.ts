import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { StudentResponse } from '../dto/responses/student-response';

@Injectable({
  providedIn: 'root'
})
export class GradesService {
  private teacherCourseUrl: string = 'http://localhost:8080/api/subjects/{subjectId}/students'

  constructor(private httpClient: HttpClient) {}

  public getStudentsperCourse(courseId: number): Observable<StudentResponse> {
    const url = this.teacherCourseUrl.replace('{subjectId}', courseId.toString());
    return this.httpClient.get<StudentResponse>(url);
  }

  public editGrade(studentId: number, courseId: number) {
    
  }
}
