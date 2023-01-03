import { Component, OnInit } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CookieService } from 'ngx-cookie-service';
import { SubjectService } from 'src/app/services/subject.service';

@Component({
  selector: 'app-students-home',
  templateUrl: './students-home.component.html',
  styleUrls: ['./students-home.component.css']
})
export class StudentsHomeComponent implements OnInit {
  subjects: Subject[] = [];

  constructor(private subjectService: SubjectService,
              private cookieService: CookieService) {}

  ngOnInit(): void {
    this.getStudentCourses();
  }

  getStudentCourses() {
    const jwtService = new JwtHelperService();
    const studentId: string = jwtService.decodeToken(this.cookieService.get("user-jwt"))['user_id'];
    this.subjectService.getStudentGrades(studentId).subscribe({
      next: response => {
        this.subjects = response
      },
      error: error => {
        console.log(error);
      }
    });
  }
}

export class Subject {
  name: string = "";
//  teacher: Teacher = new Teacher();
}