import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Observable } from 'rxjs';
import { RegisterRequest } from '../dto/requests/register-request';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  constructor(private httpClient: HttpClient) {}

  public register(registerRequest: RegisterRequest): Observable<any> {
    //TODO: to be discussed details around communication with BE
    return this.httpClient.post("", registerRequest);
  }
}
