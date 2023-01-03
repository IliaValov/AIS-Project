import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { first } from 'rxjs/operators';
import {CookieService} from 'ngx-cookie-service';

import { AccountService } from 'src/app/services/account.service';
import { JwtDecodeService } from 'src/app/services/jwt-decode.service';

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
      private jwtDecodeService: JwtDecodeService
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
            .subscribe({
                next: token => {
                    // on success
                    this.cookieService.set('user-jwt', token);

                    const role: boolean = this.jwtDecodeService.getDecodedAccessToken(token)['AdminRights'];
                    if (role) {
                        this.router.navigate(['/home/teacher']);
                    } else {
                        this.router.navigate(['/home/student']);
                    }
                },
                error: error => {
                    // on error
                }
            });
  }
}
