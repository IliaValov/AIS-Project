import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { StudentResponse } from 'src/app/dto/responses/student-response';
import { GradesService } from 'src/app/services/grades.service';

@Component({
  selector: 'app-grades',
  templateUrl: './grades.component.html',
  styleUrls: ['./grades.component.css']
})
export class GradesComponent implements OnInit {
  public students: StudentResponse = new StudentResponse();
  private courseId: number = 0;

  constructor(private gradesService: GradesService,
    private actRoute: ActivatedRoute,
    private router: Router) {}

  ngOnInit(): void {
    this.courseId = Number(this.actRoute.snapshot.paramMap.get('id'));
    this.gradesService.getStudentsperCourse(this.courseId);
  }

  routeToEditFormGrade(studentId: number) {
    this.router.navigate([`/home/teacher/course/${this.courseId}/student/${studentId}/edit-grade`]);
  }
}
