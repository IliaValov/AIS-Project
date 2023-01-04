import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { StudentsHomeComponent } from './components/students-home/students-home.component';
import { AuthGuard } from './guards/auth.guard';
import { TeachersHomeComponent } from './components/teachers-home/teachers-home.component';
import { RoleGuard } from './guards/role.guard';

const routes: Routes = [
  {
    path: 'register',
    component: RegisterComponent
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'home/student',
    component: StudentsHomeComponent,
    canActivate: [AuthGuard]
  },
  {
    path: 'home/teacher',
    component: TeachersHomeComponent,
    canActivate: [AuthGuard, RoleGuard],
    data: {
      expectedRole: 'TEACHER'
    }
  }  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
