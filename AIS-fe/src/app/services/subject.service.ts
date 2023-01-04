import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { SubjectResponse } from '../dto/responses/subject-response';

@Injectable({
  providedIn: 'root'
})
export class SubjectService {
  private gradesUrl: string = 'http://localhost:8080/api/subjects/student/{studentId}/grades';

  constructor(private httpClient: HttpClient,
              private cookieService: CookieService) { }

  public getStudentGrades(userId: string): Observable<SubjectResponse> {
    const url = this.gradesUrl.replace('{studentId}', userId);
    return this.httpClient.get<SubjectResponse>(url, {
      headers: { 'Set-Cookie': this.cookieService.get('user-jwt') }
    });
  }
}
