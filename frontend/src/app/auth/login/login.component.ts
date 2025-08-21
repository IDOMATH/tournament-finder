import { Component, inject } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { TournamentService, UserService } from "../../services/user-service";

@Component({
  selector: "app-login",
  imports: [],
  templateUrl: "./login.component.html",
  styleUrl: "./login.component.css",
})
export class LoginComponent {
  private tournamentService = inject(UserService);
  form = new FormGroup({
    email: new FormControl({ validators: [Validators.required] }),
    password: new FormControl({ validators: [Validators.required] }),
  });

  onSubmit() {
    console.log("Submitted");
  }
}
