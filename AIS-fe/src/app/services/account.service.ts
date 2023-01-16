import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Observable } from 'rxjs';
import { Token } from '@angular/compiler';
import { Constants } from './constants';

@Injectable({
  providedIn: 'root'
})
export class AccountService {
  private loginUrl: string = Constants.BackendUrl + 'login'

  constructor(private httpClient: HttpClient) { }

  public login(username: String, password: String): Observable<any> {
    return this.httpClient.post(this.loginUrl, { username, password });
  }
}
