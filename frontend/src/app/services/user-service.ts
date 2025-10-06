import { HttpClient } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { catchError } from "rxjs";

@Injectable({ providedIn: "root" })
export class UserService {
  private httpClient = inject(HttpClient);
  signUp(email: string, password: string) {
    this.httpClient
      .post("http://localhost:8080/users/", { email, password })
      .pipe(
        catchError((error) => {
          console.log("failed to sign user up");
          return throwError(() => new Error("failed to sign user up"));
        })
      );
  }
  login(email: string, password: string) {
    this.httpClient.post("http://localhost:8080/login/", { email, password });
  }
}
