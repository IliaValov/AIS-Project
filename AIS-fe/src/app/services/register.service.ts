import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Observable } from 'rxjs';
import { RegisterRequest } from '../dto/requests/register-request';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {
  private registerUrl: string = 'http://localhost:8080/api/register'

  constructor(private httpClient: HttpClient) {}

  public register(registerRequest: RegisterRequest): Observable<any> {
    console.log("Try to register :melting face");
    return this.httpClient.post(this.registerUrl, registerRequest);
  }
}
