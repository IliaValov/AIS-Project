import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RegisterRequest } from 'src/app/dto/requests/register-request';
import { RegisterResponse } from 'src/app/dto/responses/register-response';
import { RegisterService } from 'src/app/services/register.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  public form: FormGroup = new FormGroup({});
  public loading = false;
  public submitted = false;
  public successfullyRegistered = false;
  public registerMessage: string = "";

  constructor(private formBuilder: FormBuilder, private registerService: RegisterService) { }

  ngOnInit() {
    this.form = this.formBuilder.group({
        firstName: ['', Validators.required],
        lastName: ['', Validators.required],
        password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  // convenience getter for easy access to form fields
  get f() { return this.form.controls; }

  onSubmit() {
    this.submitted = true;

    if (this.form.invalid) {
      return;
    }

    this.loading = true;

    const registerRequest: RegisterRequest = new RegisterRequest(
      this.form.value.firstName,
      this.form.value.lastName,
      this.form.value.password
    )
    this.registerService.register(registerRequest).subscribe({
      next: (response: RegisterResponse)  => {
        this.successfullyRegistered = true;
        this.loading = false;
        this.registerMessage = 'Successfully registered with ' + response.username;
      },
      error: error => {
        this.successfullyRegistered = true;
        this.registerMessage = 'Failed to register'
        this.loading = false;
      }
    })
  }
}