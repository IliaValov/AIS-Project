import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { EditGradeRequest } from '../dto/requests/edit-grade-request';
import { StudentResponse } from '../dto/responses/student-response';

@Injectable({
  providedIn: 'root'
})
export class GradesService {
  private teacherCourseUrl: string = 'http://localhost:8080/api/subjects/{subjectId}/students'
  private studentsWithGradesPerCourse: string = 'http://localhost:8080/api/subjects/teacher/course/{courseId}/students/grades'
  private editGradeUrl: string = 'http://localhost:8080/api/subjects/editgrade'

  constructor(private httpClient: HttpClient, private cookieService: CookieService) {}

  public getStudentsperCourse(courseId: number): Observable<StudentResponse> {
    const url = this.studentsWithGradesPerCourse.replace('{courseId}', courseId.toString());
    return this.httpClient.get<StudentResponse>(url, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    });
  }

  public editGrade(studentId: number, courseId: number, grade: string) {
    const editGrade: EditGradeRequest = new EditGradeRequest(
      courseId.toString(),
      studentId.toString(),
      grade
    );

    return this.httpClient.post(this.editGradeUrl, editGrade, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    })
  }
}
