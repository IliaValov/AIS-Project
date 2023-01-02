import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(private httpClient: HttpClient) { }

  public login(username: String, password: String): Observable<any> {
    //TODO: to be discussed details around communication with BE
    return this.httpClient.post("", { username, password });
  }
}
