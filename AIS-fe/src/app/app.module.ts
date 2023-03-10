import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';
import { StudentsHomeComponent } from './components/students-home/students-home.component';
import { TeachersHomeComponent } from './components/teachers-home/teachers-home.component';
import { GradesComponent } from './components/grades/grades.component';
import { EditGradeComponent } from './components/edit-grade/edit-grade.component';
import { EnrollmentComponent } from './components/enrollment/enrollment.component';

@NgModule({
  declarations: [
    AppComponent,
    RegisterComponent,
    LoginComponent,
    StudentsHomeComponent,
    TeachersHomeComponent,
    GradesComponent,
    EditGradeComponent,
    EnrollmentComponent
  ],
  imports: [
    BrowserModule,
    ReactiveFormsModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
