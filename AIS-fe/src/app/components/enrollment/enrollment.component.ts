import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { StudentNotEnrolledResponse } from 'src/app/dto/responses/student-not-enrolled-response';
import { EnrollService } from 'src/app/services/enroll.service';
import {CookieService} from 'ngx-cookie-service';
import { JwtHelperService } from '@auth0/angular-jwt';

@Component({
  selector: 'app-enrollment',
  templateUrl: './enrollment.component.html',
  styleUrls: ['./enrollment.component.css']
})
export class EnrollmentComponent implements OnInit {
  public students: StudentNotEnrolledResponse = new StudentNotEnrolledResponse();
  private courseId: number = 0;

  constructor(
    private enrollService: EnrollService,
    private actRoute: ActivatedRoute,
    private router: Router,
    private cookieService: CookieService) {}

  ngOnInit(): void {
    this.courseId = Number(this.actRoute.snapshot.paramMap.get('studentId'));
    this.enrollService.getStudentsNotEnrolled(this.courseId).subscribe({
      next: (response: StudentNotEnrolledResponse)  => {
        console.log(response);
        this.students = response;
      },
      error: error => {
        console.log(error);
      }
    });
  }

  makeEnrollRequest(studentId: number) {
    const courseId: number = Number(this.actRoute.snapshot.paramMap.get('studentId'));
    
    const jwtService: JwtHelperService = new JwtHelperService();
    const token: string = this.cookieService.get('user-jwt');
    const teacherId: number = Number(jwtService.decodeToken()['user_id']);
    
    this.enrollService.enrollStudent(studentId, teacherId, courseId).subscribe({
      next: (response: any)  => {
        console.log(response);
      },
      error: error => {
        console.log(error);
      }
    });
  }

  deleteCookie() {
    this.cookieService.delete('user-jwt');
  }
}
