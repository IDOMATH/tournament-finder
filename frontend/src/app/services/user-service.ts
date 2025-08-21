import { HttpClient } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";

@Injectable({ providedIn: "root" })
export class UserService {
  private httpClient = inject(HttpClient);
  signUp(email: string, password: string) {
    this.httpClient.post("http://localhost:8080/users/", { email, password });
  }
  login(email: string, password: string) {
    this.httpClient.post("http://localhost:8080/login/", { email, password });
  }
}
