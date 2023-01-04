import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { first } from 'rxjs/operators';
import {CookieService} from 'ngx-cookie-service';
import { JwtHelperService } from '@auth0/angular-jwt';

import { AccountService } from 'src/app/services/account.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form: FormGroup = new FormGroup({});
  loading = false;
  submitted = false;

  constructor(
      private formBuilder: FormBuilder,
      private route: ActivatedRoute,
      private router: Router,
      private accountService: AccountService,
      private cookieService: CookieService,
  ) { }

  ngOnInit() {
      this.form = this.formBuilder.group({
          username: ['', Validators.required],
          password: ['', Validators.required]
      });
  }

  // convenience getter for easy access to form fields
  get f() { return this.form.controls; }

  onSubmit() {
      // stop here if form is invalid
      if (this.form.invalid) {
        return;
      }

      this.accountService.login(this.f['username'].value, this.f['password'].value)
            .pipe(first())
            .subscribe(
                (response: {token:string}) => {
                    this.cookieService.set('user-jwt', response['token']);
                    //localStorage.setItem("jwt", token)
                    const jwtService: JwtHelperService = new JwtHelperService();
                    const role: boolean = jwtService.decodeToken(response['token'])['user_admin_rights'];
                    if (role) {
                        this.router.navigate(['/home/teacher']);
                    } else {
                        this.router.navigate(['/home/student']);
                    }       
                },
                (error: HttpErrorResponse) => {
                    console.log(error);
                }
            );
  }
}
