import { Component, OnInit } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CookieService } from 'ngx-cookie-service';
import { SubjectService } from 'src/app/services/subject.service';
import { TeacherSubjectResponse } from 'src/app/dto/responses/teacher-subject-response';
import { Router } from '@angular/router';

@Component({
  selector: 'app-teachers-home',
  templateUrl: './teachers-home.component.html',
  styleUrls: ['./teachers-home.component.css']
})
export class TeachersHomeComponent implements OnInit {
  subjects: TeacherSubjectResponse = new TeacherSubjectResponse();

  constructor(private teacherSubjectService: SubjectService,
    private cookieService: CookieService,
    private router: Router) {}

  ngOnInit(): void {
    this.getTeacherCourses();
  }

  getTeacherCourses() {
    const jwtService = new JwtHelperService();
    const teacherId: string = jwtService.decodeToken(this.cookieService.get("user-jwt"))['user_id'];
    this.teacherSubjectService.getTeacherSubjects(teacherId).subscribe({
      next: (response: TeacherSubjectResponse) => {
        this.subjects = response
      },
      error: error => {
        console.log(error);
      }
    });
  }

  routeToStudentGrades(id: string) {
    this.router.navigate([`/home/teacher/course/${id}`]);
  }

  deleteCookie() {
    this.cookieService.delete('user-jwt');
  }
}
