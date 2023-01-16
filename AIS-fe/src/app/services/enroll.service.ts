import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { EnrollRequest } from '../dto/requests/enroll-request';
import { StudentNotEnrolledResponse } from '../dto/responses/student-not-enrolled-response';
import { Constants } from './constants';

@Injectable({
  providedIn: 'root'
})
export class EnrollService {
  private studentsNotEnrolledUrl: string = Constants.BackendUrl + 'subjects/{courseId}/students/notenrolled'
  private enrollStudentUrl: string = Constants.BackendUrl + 'subjects/joinsubject'

  constructor(private httpClient: HttpClient, private cookieService: CookieService) { }

  public getStudentsNotEnrolled(courseId: number): Observable<StudentNotEnrolledResponse> {
    const url = this.studentsNotEnrolledUrl.replace('{courseId}', courseId.toString());
    return this.httpClient.get<StudentNotEnrolledResponse>(url, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    });
  }

  public enrollStudent(studentId: number, teacherId: number, courseId: number) {
    const enrollRequest: EnrollRequest = new EnrollRequest(
      studentId.toString(),
      teacherId.toString(),
      courseId.toString()
    );

    return this.httpClient.post(this.enrollStudentUrl, enrollRequest, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    })
  }
}
