import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { EditGradeRequest } from 'src/app/dto/requests/edit-grade-request';
import { GradesService } from 'src/app/services/grades.service';

@Component({
  selector: 'app-edit-grade',
  templateUrl: './edit-grade.component.html',
  styleUrls: ['./edit-grade.component.css']
})
export class EditGradeComponent implements OnInit {
  public courseId: number = 0;
  public studentId: number = 0;
  public form: FormGroup = new FormGroup({});

  constructor(private gradesService: GradesService,
    private actRoute: ActivatedRoute,
    private formBuilder: FormBuilder) {}

  ngOnInit(): void {
    this.courseId = Number(this.actRoute.snapshot.paramMap.get('courseId'));
    this.studentId = Number(this.actRoute.snapshot.paramMap.get('studentId'));
    this.form = this.formBuilder.group({
      grade: [''],
    });
  }

  public editGrade() {
    console.log(this.courseId);
    console.log(this.studentId);
    console.log(this.form.value.grade);
    this.gradesService.editGrade(this.courseId, this.studentId, this.form.value.grade).subscribe(
      {
        next: response  => {
          console.log(response)
        },
        error: error => {
          console.log(error)
        }
      }
    );
  }
}
