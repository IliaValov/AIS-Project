import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { SubjectResponse } from '../dto/responses/subject-response';
import { TeacherSubjectResponse } from '../dto/responses/teacher-subject-response';

@Injectable({
  providedIn: 'root'
})
export class SubjectService {
  private gradesUrl: string = 'https://localhost:8080/api/subjects/student/{studentId}/grades';
  private teacherSubjectsUrl: string = 'https://localhost:8080/api/subjects/teacher/{teacherId}/courses'

  constructor(private httpClient: HttpClient,
              private cookieService: CookieService) { }

  public getStudentGrades(userId: string): Observable<SubjectResponse> {
    const url = this.gradesUrl.replace('{studentId}', userId);
    return this.httpClient.get<SubjectResponse>(url, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    });
  }

  public getTeacherSubjects(userId: string): Observable<TeacherSubjectResponse> {
    const url = this.teacherSubjectsUrl.replace('{teacherId}', userId);
    return this.httpClient.get<TeacherSubjectResponse>(url, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    })
  }
}
