import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Observable } from 'rxjs';
import { Token } from '@angular/compiler';

@Injectable({
  providedIn: 'root'
})
export class AccountService {
  private loginUrl: string = 'http://localhost:8080/api/login'

  constructor(private httpClient: HttpClient) { }

  public login(username: String, password: String): Observable<any> {
    console.log("Try to login");
    return this.httpClient.post(this.loginUrl, { username, password });
  }
}
